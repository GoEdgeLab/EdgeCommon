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

// NodeConfig 边缘节点配置
type NodeConfig struct {
	Id                     int64                         `yaml:"id" json:"id"`
	NodeId                 string                        `yaml:"nodeId" json:"nodeId"`
	Secret                 string                        `yaml:"secret" json:"secret"`
	IsOn                   bool                          `yaml:"isOn" json:"isOn"`
	Servers                []*serverconfigs.ServerConfig `yaml:"servers" json:"servers"`
	Version                int64                         `yaml:"version" json:"version"`
	Name                   string                        `yaml:"name" json:"name"`
	MaxCPU                 int32                         `yaml:"maxCPU" json:"maxCPU"`
	RegionId               int64                         `yaml:"regionId" json:"regionId"`
	MaxCacheDiskCapacity   *shared.SizeCapacity          `yaml:"maxCacheDiskCapacity" json:"maxCacheDiskCapacity"`
	MaxCacheMemoryCapacity *shared.SizeCapacity          `yaml:"maxCacheMemoryCapacity" json:"maxCacheMemoryCapacity"`

	// 全局配置
	GlobalConfig *serverconfigs.GlobalConfig `yaml:"globalConfig" json:"globalConfig"` // 全局配置

	// 集群统一配置
	HTTPFirewallPolicies []*firewallconfigs.HTTPFirewallPolicy   `yaml:"httpFirewallPolicies" json:"httpFirewallPolicies"`
	HTTPCachePolicies    []*serverconfigs.HTTPCachePolicy        `yaml:"httpCachePolicies" json:"httpCachePolicies"`
	TOA                  *TOAConfig                              `yaml:"toa" json:"toa"`
	SystemServices       map[string]maps.Map                     `yaml:"systemServices" json:"systemServices"` // 系统服务配置 type => params
	FirewallActions      []*firewallconfigs.FirewallActionConfig `yaml:"firewallActions" json:"firewallActions"`

	MetricItems []*serverconfigs.MetricItemConfig `yaml:"metricItems" json:"metricItems"`

	paddedId string

	// firewall
	firewallPolicies []*firewallconfigs.HTTPFirewallPolicy

	// metrics
	hasHTTPConnectionMetrics bool
}

// SharedNodeConfig 取得当前节点配置单例
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

// ResetNodeConfig 重置节点配置
func ResetNodeConfig(nodeConfig *NodeConfig) {
	shared.Locker.Lock()
	sharedNodeConfig = nodeConfig
	shared.Locker.Unlock()
}

// Init 初始化
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
	if len(this.HTTPCachePolicies) > 0 {
		for _, policy := range this.HTTPCachePolicies {
			err := policy.Init()
			if err != nil {
				return err
			}
		}
	}

	// firewall policy
	if len(this.HTTPFirewallPolicies) > 0 {
		for _, policy := range this.HTTPFirewallPolicies {
			err := policy.Init()
			if err != nil {
				return err
			}
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
	for _, policy := range this.HTTPFirewallPolicies {
		if policy.IsOn {
			this.firewallPolicies = append(this.firewallPolicies, policy)
		}
	}
	for _, server := range this.Servers {
		if !server.IsOk() || !server.IsOn {
			continue
		}

		// WAF策略
		if server.HTTPFirewallPolicyId > 0 {
			for _, policy := range this.HTTPFirewallPolicies {
				if server.HTTPFirewallPolicyId == policy.Id {
					server.HTTPFirewallPolicy = policy
					break
				}
			}
		}

		// 缓存策略
		if server.HTTPCachePolicyId > 0 {
			for _, policy := range this.HTTPCachePolicies {
				if server.HTTPCachePolicyId == policy.Id {
					server.HTTPCachePolicy = policy
				}
			}
		}

		if server.Web != nil {
			this.lookupWeb(server, server.Web)
		}
	}

	// firewall actions
	for _, action := range this.FirewallActions {
		err := action.Init()
		if err != nil {
			return err
		}
	}

	// metric items
	this.hasHTTPConnectionMetrics = false
	for _, item := range this.MetricItems {
		err := item.Init()
		if err != nil {
			return err
		}
		if item.IsOn && item.HasHTTPConnectionValue() {
			this.hasHTTPConnectionMetrics = true
		}
	}

	return nil
}

// AvailableGroups 根据网络地址和协议分组
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

// FindAllFirewallPolicies 获取所有的防火墙策略
func (this *NodeConfig) FindAllFirewallPolicies() []*firewallconfigs.HTTPFirewallPolicy {
	return this.firewallPolicies
}

// Save 写入到文件
func (this *NodeConfig) Save() error {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	data, err := json.Marshal(this)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(Tea.ConfigFile("node.json"), data, 0777)
}

// PaddedId 获取填充后的ID
func (this *NodeConfig) PaddedId() string {
	return this.paddedId
}

// HasHTTPConnectionMetrics 是否含有HTTP连接数的指标
func (this *NodeConfig) HasHTTPConnectionMetrics() bool {
	return this.hasHTTPConnectionMetrics
}

// 搜索WAF策略
func (this *NodeConfig) lookupWeb(server *serverconfigs.ServerConfig, web *serverconfigs.HTTPWebConfig) {
	if web == nil || !web.IsOn {
		return
	}
	if web.FirewallPolicy != nil && web.FirewallPolicy.IsOn {
		// 复用节点的拦截选项设置
		if web.FirewallPolicy.BlockOptions == nil && server.HTTPFirewallPolicy != nil && server.HTTPFirewallPolicy.BlockOptions != nil {
			web.FirewallPolicy.BlockOptions = server.HTTPFirewallPolicy.BlockOptions
		}
		this.firewallPolicies = append(this.firewallPolicies, web.FirewallPolicy)
	}
	if len(web.Locations) > 0 {
		for _, location := range web.Locations {
			if location.Web != nil && location.Web.IsOn {
				this.lookupWeb(server, location.Web)
			}
		}
	}
}
