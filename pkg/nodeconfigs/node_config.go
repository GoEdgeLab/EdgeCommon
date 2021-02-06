package nodeconfigs

import (
	"encoding/json"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/maps"
	"io/ioutil"
	"strconv"
)

var sharedNodeConfig *NodeConfig = nil

type NodeConfig struct {
	Id       int64                         `yaml:"id" json:"id"`
	NodeId   string                        `yaml:"nodeId" json:"nodeId"`
	IsOn     bool                          `yaml:"isOn" json:"isOn"`
	Servers  []*serverconfigs.ServerConfig `yaml:"servers" json:"servers"`
	Version  int64                         `yaml:"version" json:"version"`
	Name     string                        `yaml:"name" json:"name"`
	MaxCPU   int32                         `yaml:"maxCPU" json:"maxCPU"`
	RegionId int64                         `yaml:"regionId" json:"regionId"`

	// 全局配置
	GlobalConfig *serverconfigs.GlobalConfig `yaml:"globalConfig" json:"globalConfig"` // 全局配置

	// 集群统一配置
	HTTPFirewallPolicy *firewallconfigs.HTTPFirewallPolicy     `yaml:"httpFirewallPolicy" json:"httpFirewallPolicy"`
	HTTPCachePolicy    *serverconfigs.HTTPCachePolicy          `yaml:"httpCachePolicy" json:"httpCachePolicy"`
	TOA                *TOAConfig                              `yaml:"toa" json:"toa"`
	SystemServices     map[string]maps.Map                     `yaml:"systemServices" json:"systemServices"` // 系统服务配置 type => params
	FirewallActions    []*firewallconfigs.FirewallActionConfig `yaml:"firewallActions" json:"firewallActions"`

	paddedId string

	firewallPolicies []*firewallconfigs.HTTPFirewallPolicy
}

// 取得当前节点配置单例
func SharedNodeConfig() (*NodeConfig, error) {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	if sharedNodeConfig != nil {
		return sharedNodeConfig, nil
	}

	data, err := ioutil.ReadFile(Tea.ConfigFile("node.json"))
	if err != nil {
		return &NodeConfig{}, err
	}

	config := &NodeConfig{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return &NodeConfig{}, err
	}

	sharedNodeConfig = config
	return config, nil
}

// 重置节点配置
func ResetNodeConfig(nodeConfig *NodeConfig) {
	shared.Locker.Lock()
	sharedNodeConfig = nodeConfig
	shared.Locker.Unlock()
}

// 初始化
func (this *NodeConfig) Init() error {
	this.paddedId = fmt.Sprintf("%08d", this.Id)

	// servers
	for _, server := range this.Servers {
		err := server.Init()
		if err != nil {
			// 这里不返回错误，而是继续往下，防止单个服务错误而影响其他服务
			logs.Println("[INIT]server '" + strconv.FormatInt(server.Id, 10) + "' init failed: " + err.Error())
		}
	}

	// global config
	if this.GlobalConfig != nil {
		err := this.GlobalConfig.Init()
		if err != nil {
			return err
		}
	}

	// cache policy
	if this.HTTPCachePolicy != nil {
		err := this.HTTPCachePolicy.Init()
		if err != nil {
			return err
		}
	}

	// firewall policy
	if this.HTTPFirewallPolicy != nil {
		err := this.HTTPFirewallPolicy.Init()
		if err != nil {
			return err
		}
	}

	// TOA
	if this.TOA != nil {
		err := this.TOA.Init()
		if err != nil {
			return err
		}
	}

	// 查找FirewallPolicy
	this.firewallPolicies = []*firewallconfigs.HTTPFirewallPolicy{}
	if this.HTTPFirewallPolicy != nil && this.HTTPFirewallPolicy.IsOn {
		this.firewallPolicies = append(this.firewallPolicies, this.HTTPFirewallPolicy)
	}
	for _, server := range this.Servers {
		if !server.IsOk() || !server.IsOn {
			continue
		}
		if server.Web != nil {
			this.lookupWeb(server.Web)
		}
	}

	// firewall actions
	for _, action := range this.FirewallActions {
		err := action.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 根据网络地址和协议分组
func (this *NodeConfig) AvailableGroups() []*serverconfigs.ServerGroup {
	groupMapping := map[string]*serverconfigs.ServerGroup{} // protocol://addr => Server Group
	for _, server := range this.Servers {
		if !server.IsOk() || !server.IsOn {
			continue
		}
		for _, addr := range server.FullAddresses() {
			group, ok := groupMapping[addr]
			if ok {
				group.Add(server)
			} else {
				group = serverconfigs.NewServerGroup(addr)
				group.Add(server)
			}
			groupMapping[addr] = group
		}
	}
	result := []*serverconfigs.ServerGroup{}
	for _, group := range groupMapping {
		result = append(result, group)
	}
	return result
}

// 获取所有的防火墙策略
func (this *NodeConfig) FindAllFirewallPolicies() []*firewallconfigs.HTTPFirewallPolicy {
	return this.firewallPolicies
}

// 写入到文件
func (this *NodeConfig) Save() error {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	data, err := json.Marshal(this)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(Tea.ConfigFile("node.json"), data, 0777)
}

// 获取填充后的ID
func (this *NodeConfig) PaddedId() string {
	return this.paddedId
}

// 搜索WAF策略
func (this *NodeConfig) lookupWeb(web *serverconfigs.HTTPWebConfig) {
	if web == nil || !web.IsOn {
		return
	}
	if web.FirewallPolicy != nil && web.FirewallPolicy.IsOn {
		// 复用节点的拦截选项设置
		if web.FirewallPolicy.BlockOptions == nil && this.HTTPFirewallPolicy != nil && this.HTTPFirewallPolicy.BlockOptions != nil {
			web.FirewallPolicy.BlockOptions = this.HTTPFirewallPolicy.BlockOptions
		}
		this.firewallPolicies = append(this.firewallPolicies, web.FirewallPolicy)
	}
	if len(web.Locations) > 0 {
		for _, location := range web.Locations {
			if location.Web != nil && location.Web.IsOn {
				this.lookupWeb(location.Web)
			}
		}
	}
}
