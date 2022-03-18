package nodeconfigs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/maps"
	"io/ioutil"
	"reflect"
	"strconv"
)

var sharedNodeConfig *NodeConfig = nil

type ServerError struct {
	Id      int64
	Message string
}

func NewServerError(serverId int64, message string) *ServerError {
	return &ServerError{Id: serverId, Message: message}
}

// NodeConfig 边缘节点配置
type NodeConfig struct {
	Id           int64                         `yaml:"id" json:"id"`
	NodeId       string                        `yaml:"nodeId" json:"nodeId"`
	Secret       string                        `yaml:"secret" json:"secret"`
	IsOn         bool                          `yaml:"isOn" json:"isOn"`
	Servers      []*serverconfigs.ServerConfig `yaml:"servers" json:"servers"`
	SupportCNAME bool                          `yaml:"supportCNAME" json:"supportCNAME"`
	Version      int64                         `yaml:"version" json:"version"`
	Name         string                        `yaml:"name" json:"name"`
	RegionId     int64                         `yaml:"regionId" json:"regionId"`
	OCSPVersion  int64                         `yaml:"ocspVersion" json:"ocspVersion"`

	// 性能
	MaxCPU                 int32                `yaml:"maxCPU" json:"maxCPU"`
	CacheDiskDir           string               `yaml:"cacheDiskDir" json:"cacheDiskDir"`                     // 文件缓存目录
	MaxCacheDiskCapacity   *shared.SizeCapacity `yaml:"maxCacheDiskCapacity" json:"maxCacheDiskCapacity"`     // 文件缓存容量
	MaxCacheMemoryCapacity *shared.SizeCapacity `yaml:"maxCacheMemoryCapacity" json:"maxCacheMemoryCapacity"` // 内容缓存容量
	MaxThreads             int                  `yaml:"maxThreads" json:"maxThreads"`
	TCPMaxConnections      int                  `yaml:"tcpMaxConnections" json:"tcpMaxConnections"`

	// 全局配置
	GlobalConfig  *serverconfigs.GlobalConfig `yaml:"globalConfig" json:"globalConfig"` // 全局配置
	ProductConfig *ProductConfig              `yaml:"productConfig" json:"productConfig"`

	// 集群统一配置
	HTTPFirewallPolicies []*firewallconfigs.HTTPFirewallPolicy   `yaml:"httpFirewallPolicies" json:"httpFirewallPolicies"`
	HTTPCachePolicies    []*serverconfigs.HTTPCachePolicy        `yaml:"httpCachePolicies" json:"httpCachePolicies"`
	TOA                  *TOAConfig                              `yaml:"toa" json:"toa"`
	SystemServices       map[string]maps.Map                     `yaml:"systemServices" json:"systemServices"` // 系统服务配置 type => params
	FirewallActions      []*firewallconfigs.FirewallActionConfig `yaml:"firewallActions" json:"firewallActions"`
	TimeZone             string                                  `yaml:"timeZone" json:"timeZone"`
	AutoOpenPorts        bool                                    `yaml:"autoOpenPorts" json:"autoOpenPorts"`

	// 指标
	MetricItems []*serverconfigs.MetricItemConfig `yaml:"metricItems" json:"metricItems"`

	// 自动白名单
	AllowedIPs []string `yaml:"allowedIPs" json:"allowedIPs"`

	paddedId string

	// firewall
	firewallPolicies []*firewallconfigs.HTTPFirewallPolicy

	// metrics
	hasHTTPConnectionMetrics bool

	// 源站集合
	originMap map[int64]*serverconfigs.OriginConfig

	// 自动白名单
	allowedIPMap map[string]bool

	// syn flood
	synFlood *firewallconfigs.SYNFloodConfig
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

// CloneNodeConfig 复制节点配置
func CloneNodeConfig(nodeConfig *NodeConfig) (*NodeConfig, error) {
	if nodeConfig == nil {
		return nil, errors.New("node config should not be nil")
	}

	var newConfigValue = reflect.Indirect(reflect.ValueOf(&NodeConfig{}))
	var oldValue = reflect.Indirect(reflect.ValueOf(nodeConfig))
	var valueType = oldValue.Type()
	for i := 0; i < valueType.NumField(); i++ {
		var field = valueType.Field(i)
		var fieldName = field.Name
		if !field.IsExported() {
			continue
		}
		if fieldName == "Servers" {
			continue
		}

		newConfigValue.FieldByName(fieldName).Set(oldValue.FieldByName(fieldName))
	}

	var newConfig = newConfigValue.Interface().(NodeConfig)
	newConfig.Servers = append([]*serverconfigs.ServerConfig{}, nodeConfig.Servers...)
	return &newConfig, nil
}

// Init 初始化
func (this *NodeConfig) Init() (err error, serverErrors []*ServerError) {
	this.paddedId = fmt.Sprintf("%08d", this.Id)

	// servers
	for _, server := range this.Servers {
		// 初始化
		errs := server.Init()
		if len(errs) > 0 {
			// 这里不返回错误，而是继续往下，防止单个服务错误而影响其他服务
			for _, serverErr := range errs {
				serverErrors = append(serverErrors, NewServerError(server.Id, "server '"+strconv.FormatInt(server.Id, 10)+"' init failed: "+serverErr.Error()))
			}
		}

		// 检查ACME支持
		if server.IsOn && server.SupportCNAME {
			this.SupportCNAME = true
		}
	}

	// global config
	if this.GlobalConfig != nil {
		err = this.GlobalConfig.Init()
		if err != nil {
			return
		}
	}

	// cache policy
	if len(this.HTTPCachePolicies) > 0 {
		for _, policy := range this.HTTPCachePolicies {
			err = policy.Init()
			if err != nil {
				return
			}
		}
	}

	// firewall policy
	if len(this.HTTPFirewallPolicies) > 0 {
		for _, policy := range this.HTTPFirewallPolicies {
			err = policy.Init()
			if err != nil {
				return
			}
		}
	}

	// TOA
	if this.TOA != nil {
		err = this.TOA.Init()
		if err != nil {
			return
		}
	}

	// 源站
	this.originMap = map[int64]*serverconfigs.OriginConfig{}

	// 查找FirewallPolicy
	this.synFlood = nil
	this.firewallPolicies = []*firewallconfigs.HTTPFirewallPolicy{}
	for _, policy := range this.HTTPFirewallPolicies {
		if policy.IsOn {
			this.firewallPolicies = append(this.firewallPolicies, policy)
			if policy.SYNFlood != nil && policy.SYNFlood.IsOn {
				this.synFlood = policy.SYNFlood
			}
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

		// 源站
		if server.ReverseProxyRef != nil && server.ReverseProxyRef.IsOn && server.ReverseProxy != nil && server.ReverseProxy.IsOn {
			for _, origin := range server.ReverseProxy.PrimaryOrigins {
				if origin.IsOn {
					this.originMap[origin.Id] = origin
				}
			}
			for _, origin := range server.ReverseProxy.BackupOrigins {
				if origin.IsOn {
					this.originMap[origin.Id] = origin
				}
			}
		}

		if server.Web != nil {
			this.lookupWeb(server, server.Web)
		}
	}

	// firewall actions
	for _, action := range this.FirewallActions {
		err = action.Init()
		if err != nil {
			return
		}
	}

	// metric items
	this.hasHTTPConnectionMetrics = false
	for _, item := range this.MetricItems {
		err = item.Init()
		if err != nil {
			return
		}
		if item.IsOn && item.HasHTTPConnectionValue() {
			this.hasHTTPConnectionMetrics = true
		}
	}

	// 自动白名单
	this.allowedIPMap = map[string]bool{}
	for _, allowIP := range this.AllowedIPs {
		this.allowedIPMap[allowIP] = true
	}

	return
}

// AddServer 添加服务
func (this *NodeConfig) AddServer(server *serverconfigs.ServerConfig) {
	if server == nil {
		return
	}

	var found = false
	for index, oldServer := range this.Servers {
		if oldServer.Id == server.Id {
			this.Servers[index] = server
			found = true
			break
		}
	}
	if !found {
		this.Servers = append(this.Servers, server)
	}
}

// RemoveServer 删除服务
func (this *NodeConfig) RemoveServer(serverId int64) {
	for index, oldServer := range this.Servers {
		if oldServer.Id == serverId {
			this.Servers = append(this.Servers[:index], this.Servers[index+1:]...)
			break
		}
	}
}

// AvailableGroups 根据网络地址和协议分组
func (this *NodeConfig) AvailableGroups() []*serverconfigs.ServerAddressGroup {
	groupMapping := map[string]*serverconfigs.ServerAddressGroup{} // protocol://addr => Server Group
	for _, server := range this.Servers {
		if !server.IsOk() || !server.IsOn {
			continue
		}
		for _, addr := range server.FullAddresses() {
			group, ok := groupMapping[addr]
			if ok {
				group.Add(server)
			} else {
				group = serverconfigs.NewServerAddressGroup(addr)
				group.Add(server)
			}
			groupMapping[addr] = group
		}
	}
	result := []*serverconfigs.ServerAddressGroup{}
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

// FindOrigin 读取源站配置
func (this *NodeConfig) FindOrigin(originId int64) *serverconfigs.OriginConfig {
	if this.originMap == nil {
		return nil
	}
	config, ok := this.originMap[originId]
	if ok {
		return config
	}
	return nil
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
			web.FirewallPolicy.Mode = server.HTTPFirewallPolicy.Mode
			web.FirewallPolicy.UseLocalFirewall = server.HTTPFirewallPolicy.UseLocalFirewall
		}
		this.firewallPolicies = append(this.firewallPolicies, web.FirewallPolicy)
	}
	if len(web.Locations) > 0 {
		for _, location := range web.Locations {
			// 源站
			if location.IsOn && location.ReverseProxyRef != nil && location.ReverseProxyRef.IsOn && location.ReverseProxy != nil && location.ReverseProxy.IsOn {
				for _, origin := range location.ReverseProxy.PrimaryOrigins {
					if origin.IsOn {
						this.originMap[origin.Id] = origin
					}
				}
				for _, origin := range location.ReverseProxy.BackupOrigins {
					if origin.IsOn {
						this.originMap[origin.Id] = origin
					}
				}
			}

			// Web
			if location.Web != nil && location.Web.IsOn {
				this.lookupWeb(server, location.Web)
			}
		}
	}
}

// IPIsAutoAllowed 检查是否自动允许某个IP
func (this *NodeConfig) IPIsAutoAllowed(ip string) bool {
	_, ok := this.allowedIPMap[ip]
	return ok
}

// SYNFloodConfig 获取SYN Flood配置
func (this *NodeConfig) SYNFloodConfig() *firewallconfigs.SYNFloodConfig {
	return this.synFlood
}

// UpdateCertOCSP 修改证书OCSP
func (this *NodeConfig) UpdateCertOCSP(certId int64, ocsp []byte) {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	var servers = this.Servers
	for _, server := range servers {
		if server.HTTPS != nil &&
			server.HTTPS.SSLPolicy != nil &&
			server.HTTPS.SSLPolicy.OCSPIsOn &&
			server.HTTPS.SSLPolicy.ContainsCert(certId) {
			server.HTTPS.SSLPolicy.UpdateCertOCSP(certId, ocsp)
		}

		if server.TLS != nil &&
			server.TLS.SSLPolicy != nil &&
			server.TLS.SSLPolicy.OCSPIsOn &&
			server.TLS.SSLPolicy.ContainsCert(certId) {
			server.TLS.SSLPolicy.UpdateCertOCSP(certId, ocsp)
		}
	}
}
