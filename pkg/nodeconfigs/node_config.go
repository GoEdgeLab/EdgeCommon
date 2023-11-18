package nodeconfigs

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/nodeutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/ddosconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var sharedNodeConfig *NodeConfig = nil
var uamPolicyLocker = &sync.RWMutex{}
var httpCCPolicyLocker = &sync.RWMutex{}
var http3PolicyLocker = &sync.RWMutex{}
var httpPagesPolicyLocker = &sync.RWMutex{}

type ServerError struct {
	Id      int64
	Message string
}

func NewServerError(serverId int64, message string) *ServerError {
	return &ServerError{Id: serverId, Message: message}
}

// NodeConfig 边缘节点配置
type NodeConfig struct {
	Id                   int64                         `yaml:"id" json:"id"`
	Edition              string                        `yaml:"edition" json:"edition"`
	NodeId               string                        `yaml:"nodeId" json:"nodeId"`
	Secret               string                        `yaml:"secret" json:"secret"`
	IsOn                 bool                          `yaml:"isOn" json:"isOn"`
	Servers              []*serverconfigs.ServerConfig `yaml:"servers" json:"servers"`
	SupportCNAME         bool                          `yaml:"supportCNAME" json:"supportCNAME"`
	Version              int64                         `yaml:"version" json:"version"`
	Name                 string                        `yaml:"name" json:"name"`
	GroupId              int64                         `yaml:"groupId" json:"groupId"`
	RegionId             int64                         `yaml:"regionId" json:"regionId"`
	OCSPVersion          int64                         `yaml:"ocspVersion" json:"ocspVersion"`
	DataMap              *shared.DataMap               `yaml:"dataMap" json:"dataMap"`
	UpdatingServerListId int64                         `yaml:"updatingServerListId" json:"updatingServerListId"`

	// 性能
	MaxCPU       int32                                 `yaml:"maxCPU" json:"maxCPU"`
	APINodeAddrs []*serverconfigs.NetworkAddressConfig `yaml:"apiNodeAddrs" json:"apiNodeAddrs"`

	CacheDiskDir         string               `yaml:"cacheDiskDir" json:"cacheDiskDir"`                 // 文件缓存目录
	MaxCacheDiskCapacity *shared.SizeCapacity `yaml:"maxCacheDiskCapacity" json:"maxCacheDiskCapacity"` // 文件缓存容量

	CacheDiskSubDirs []*serverconfigs.CacheDir `yaml:"cacheDiskSubDirs" json:"cacheDiskSubDirs"` // 其余缓存目录

	MaxCacheMemoryCapacity *shared.SizeCapacity          `yaml:"maxCacheMemoryCapacity" json:"maxCacheMemoryCapacity"` // 内容缓存容量
	MaxThreads             int                           `yaml:"maxThreads" json:"maxThreads"`                         // 最大线程数
	DDoSProtection         *ddosconfigs.ProtectionConfig `yaml:"ddosProtection" json:"ddosProtection"`                 // DDoS防护
	EnableIPLists          bool                          `yaml:"enableIPLists" json:"enableIPLists"`                   // 启用IP名单

	// 级别
	Level       int32                         `yaml:"level" json:"level"`
	ParentNodes map[int64][]*ParentNodeConfig `yaml:"parentNodes" json:"parentNodes"` // clusterId => []*ParentNodeConfig

	// 全局配置
	GlobalServerConfig *serverconfigs.GlobalServerConfig `yaml:"globalServerConfig" json:"globalServerConfig"` // 服务全局配置，用来替代 GlobalConfig
	ProductConfig      *ProductConfig                    `yaml:"productConfig" json:"productConfig"`

	// 集群统一配置
	HTTPFirewallPolicies []*firewallconfigs.HTTPFirewallPolicy   `yaml:"httpFirewallPolicies" json:"httpFirewallPolicies"`
	HTTPCachePolicies    []*serverconfigs.HTTPCachePolicy        `yaml:"httpCachePolicies" json:"httpCachePolicies"`
	TOA                  *TOAConfig                              `yaml:"toa" json:"toa"`
	SystemServices       map[string]maps.Map                     `yaml:"systemServices" json:"systemServices"`           // 系统服务配置 type => params
	FirewallActions      []*firewallconfigs.FirewallActionConfig `yaml:"firewallActions" json:"firewallActions"`         // 防火墙动作
	TimeZone             string                                  `yaml:"timeZone" json:"timeZone"`                       // 自动设置时区
	AutoOpenPorts        bool                                    `yaml:"autoOpenPorts" json:"autoOpenPorts"`             // 自动开放所需端口
	Clock                *ClockConfig                            `yaml:"clock" json:"clock"`                             // 时钟配置
	AutoInstallNftables  bool                                    `yaml:"autoInstallNftables" json:"autoInstallNftables"` // 自动安装nftables
	AutoSystemTuning     bool                                    `yaml:"autoSystemTuning" json:"autoSystemTuning"`       // 自动调整系统参数

	// 指标
	MetricItems []*serverconfigs.MetricItemConfig `yaml:"metricItems" json:"metricItems"`

	IPAddresses []string `yaml:"ipAddresses" json:"ipAddresses"` // IP地址
	AllowedIPs  []string `yaml:"allowedIPs" json:"allowedIPs"`   // 自动IP白名单

	// 脚本
	CommonScripts []*serverconfigs.CommonScript `yaml:"commonScripts" json:"commonScripts"`

	WebPImagePolicies     map[int64]*WebPImagePolicy `yaml:"webpImagePolicies" json:"webpImagePolicies"`         // WebP相关配置，clusterId => *WebPImagePolicy
	UAMPolicies           map[int64]*UAMPolicy       `yaml:"uamPolicies" json:"uamPolicies"`                     // UAM相关配置，clusterId => *UAMPolicy
	HTTPCCPolicies        map[int64]*HTTPCCPolicy    `yaml:"httpCCPolicies" json:"httpCCPolicies"`               // CC相关配置， clusterId => *HTTPCCPolicy
	HTTP3Policies         map[int64]*HTTP3Policy     `yaml:"http3Policies" json:"http3Policies"`                 // HTTP3相关配置， clusterId => *HTTP3Policy
	HTTPPagesPolicies     map[int64]*HTTPPagesPolicy `yaml:"httpPagesPolicies" json:"httpPagesPolicies"`         // 自定义页面，clusterId => *HTTPPagesPolicy
	NetworkSecurityPolicy *NetworkSecurityPolicy     `yaml:"networkSecurityPolicy" json:"networkSecurityPolicy"` // 网络安全策略

	// DNS
	DNSResolver *DNSResolverConfig `yaml:"dnsResolver" json:"dnsResolver"`

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

	secretHash string
}

// SharedNodeConfig 取得当前节点配置单例
func SharedNodeConfig() (*NodeConfig, error) {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	if sharedNodeConfig != nil {
		return sharedNodeConfig, nil
	}

	// 从本地缓存读取
	var configFile = Tea.ConfigFile("node.json")
	var readCacheOk = false
	defer func() {
		if !readCacheOk {
			_ = os.Remove(configFile)
		}
	}()

	data, err := os.ReadFile(configFile)
	if err != nil {
		return &NodeConfig{}, err
	}

	encodedNodeInfo, encodedJSONData, found := bytes.Cut(data, []byte("\n"))
	if !found {
		// 删除缓存文件
		return &NodeConfig{}, errors.New("node.json: invalid data format")
	}

	encodedNodeInfoData, err := base64.StdEncoding.DecodeString(string(encodedNodeInfo))
	if err != nil {
		// 删除缓存文件
		return &NodeConfig{}, err
	}

	nodeUniqueId, nodeSecret, found := strings.Cut(string(encodedNodeInfoData), "|")
	if !found {
		// 删除缓存文件
		return &NodeConfig{}, errors.New("node.json: node info: invalid data format")
	}

	jsonData, err := nodeutils.DecryptData(nodeUniqueId, nodeSecret, string(encodedJSONData))
	if err != nil {
		return &NodeConfig{}, err
	}

	var config = &NodeConfig{}
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return &NodeConfig{}, err
	}

	readCacheOk = true
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

	uamPolicyLocker.RLock()
	defer uamPolicyLocker.RUnlock()

	httpCCPolicyLocker.RLock()
	defer httpCCPolicyLocker.RUnlock()

	http3PolicyLocker.RLock()
	defer http3PolicyLocker.RUnlock()

	httpPagesPolicyLocker.RLock()
	defer httpPagesPolicyLocker.RUnlock()

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
func (this *NodeConfig) Init(ctx context.Context) (err error, serverErrors []*ServerError) {
	// 设置Context
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = context.WithValue(ctx, "DataMap", this.DataMap)

	this.secretHash = fmt.Sprintf("%x", sha256.Sum256([]byte(this.NodeId+"@"+this.Secret)))
	this.paddedId = fmt.Sprintf("%08d", this.Id)

	// servers
	for _, server := range this.Servers {
		// 避免在运行时重新初始化
		if server.IsInitialized() {
			continue
		}

		// 初始化
		errs := server.Init(ctx)
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

	// webp image policy
	if this.WebPImagePolicies != nil {
		for _, policy := range this.WebPImagePolicies {
			err = policy.Init()
			if err != nil {
				return
			}
		}
	}

	// uam policy
	uamPolicyLocker.RLock()
	if len(this.UAMPolicies) > 0 {
		for _, policy := range this.UAMPolicies {
			err = policy.Init()
			if err != nil {
				uamPolicyLocker.RUnlock()
				return
			}
		}
	}
	uamPolicyLocker.RUnlock()

	// http cc policy
	httpCCPolicyLocker.RLock()
	if len(this.HTTPCCPolicies) > 0 {
		for _, policy := range this.HTTPCCPolicies {
			err = policy.Init()
			if err != nil {
				httpCCPolicyLocker.RUnlock()
				return
			}
		}
	}
	httpCCPolicyLocker.RUnlock()

	// http3 policy
	http3PolicyLocker.RLock()
	if len(this.HTTP3Policies) > 0 {
		for _, policy := range this.HTTP3Policies {
			err = policy.Init()
			if err != nil {
				http3PolicyLocker.RUnlock()
				return
			}
		}
	}
	http3PolicyLocker.RUnlock()

	// http pages policy
	httpPagesPolicyLocker.RLock()
	if len(this.HTTPPagesPolicies) > 0 {
		for _, policy := range this.HTTPPagesPolicies {
			err = policy.Init()
			if err != nil {
				httpPagesPolicyLocker.RUnlock()
				return
			}
		}
	}
	httpPagesPolicyLocker.RUnlock()

	// dns resolver
	if this.DNSResolver != nil {
		err = this.DNSResolver.Init()
		if err != nil {
			return
		}
	}

	// 全局服务设置
	if this.GlobalServerConfig != nil {
		err = this.GlobalServerConfig.Init()
		if err != nil {
			return
		}
	}

	// api node addrs
	if len(this.APINodeAddrs) > 0 {
		for _, addr := range this.APINodeAddrs {
			err = addr.Init()
			if err != nil {
				return err, nil
			}
		}
	}

	// network security policy
	if this.NetworkSecurityPolicy != nil {
		err = this.NetworkSecurityPolicy.Init()
		if err != nil {
			return err, nil
		}
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
	var groupMapping = map[string]*serverconfigs.ServerAddressGroup{} // protocol://addr => Server Group
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
	var result = []*serverconfigs.ServerAddressGroup{}
	for _, group := range groupMapping {
		result = append(result, group)
	}
	return result
}

// HTTP3Group HTTP/3网站分组
// 这里暂时不区分集群
func (this *NodeConfig) HTTP3Group() *serverconfigs.ServerAddressGroup {
	var group = serverconfigs.NewServerAddressGroup("HTTP3")
	for _, server := range this.Servers {
		if !server.SupportsHTTP3() {
			continue
		}
		group.Add(server)
	}
	return group
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

	var headerData = []byte(base64.StdEncoding.EncodeToString([]byte(this.NodeId+"|"+this.Secret)) + "\n")

	encodedData, err := nodeutils.EncryptData(this.NodeId, this.Secret, data)
	if err != nil {
		return err
	}

	return os.WriteFile(Tea.ConfigFile("node.json"), append(headerData, encodedData...), 0777)
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
		// 复用节点的选项设置
		if server.HTTPFirewallPolicy != nil {
			if (web.FirewallPolicy.BlockOptions == nil || !web.FirewallPolicy.BlockOptions.IsPrior) && server.HTTPFirewallPolicy.BlockOptions != nil {
				web.FirewallPolicy.BlockOptions = server.HTTPFirewallPolicy.BlockOptions
			}
			if (web.FirewallPolicy.CaptchaOptions == nil || !web.FirewallPolicy.CaptchaOptions.IsPrior) && server.HTTPFirewallPolicy.CaptchaOptions != nil {
				web.FirewallPolicy.CaptchaOptions = server.HTTPFirewallPolicy.CaptchaOptions
			}
			if (web.FirewallPolicy.SYNFlood == nil || !web.FirewallPolicy.SYNFlood.IsPrior) && server.HTTPFirewallPolicy.SYNFlood != nil {
				web.FirewallPolicy.SYNFlood = server.HTTPFirewallPolicy.SYNFlood
			}
			if (web.FirewallPolicy.Log == nil || !web.FirewallPolicy.Log.IsPrior) && server.HTTPFirewallPolicy.Log != nil {
				web.FirewallPolicy.Log = server.HTTPFirewallPolicy.Log
			}

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
func (this *NodeConfig) UpdateCertOCSP(certId int64, ocsp []byte, expiresAt int64) {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	var servers = this.Servers
	for _, server := range servers {
		if server.HTTPS != nil &&
			server.HTTPS.SSLPolicy != nil &&
			server.HTTPS.SSLPolicy.OCSPIsOn &&
			server.HTTPS.SSLPolicy.ContainsCert(certId) {
			server.HTTPS.SSLPolicy.UpdateCertOCSP(certId, ocsp, expiresAt)
		}

		if server.TLS != nil &&
			server.TLS.SSLPolicy != nil &&
			server.TLS.SSLPolicy.OCSPIsOn &&
			server.TLS.SSLPolicy.ContainsCert(certId) {
			server.TLS.SSLPolicy.UpdateCertOCSP(certId, ocsp, expiresAt)
		}
	}
}

// FindWebPImagePolicyWithClusterId 使用集群ID查找WebP策略
func (this *NodeConfig) FindWebPImagePolicyWithClusterId(clusterId int64) *WebPImagePolicy {
	if this.WebPImagePolicies == nil {
		return nil
	}
	return this.WebPImagePolicies[clusterId]
}

// FindUAMPolicyWithClusterId 使用集群ID查找UAM策略
func (this *NodeConfig) FindUAMPolicyWithClusterId(clusterId int64) *UAMPolicy {
	uamPolicyLocker.RLock()
	defer uamPolicyLocker.RUnlock()
	if this.UAMPolicies == nil {
		return nil
	}
	return this.UAMPolicies[clusterId]
}

// UpdateUAMPolicies 修改集群UAM策略
func (this *NodeConfig) UpdateUAMPolicies(policies map[int64]*UAMPolicy) {
	uamPolicyLocker.Lock()
	defer uamPolicyLocker.Unlock()
	this.UAMPolicies = policies
}

// FindHTTPCCPolicyWithClusterId 使用集群ID查找CC策略
func (this *NodeConfig) FindHTTPCCPolicyWithClusterId(clusterId int64) *HTTPCCPolicy {
	httpCCPolicyLocker.RLock()
	defer httpCCPolicyLocker.RUnlock()
	if this.HTTPCCPolicies == nil {
		return nil
	}
	return this.HTTPCCPolicies[clusterId]
}

// UpdateHTTPCCPolicies 修改集群CC策略
func (this *NodeConfig) UpdateHTTPCCPolicies(policies map[int64]*HTTPCCPolicy) {
	httpCCPolicyLocker.Lock()
	defer httpCCPolicyLocker.Unlock()
	this.HTTPCCPolicies = policies
}

// FindHTTP3PolicyWithClusterId 使用集群ID查找HTTP/3策略
func (this *NodeConfig) FindHTTP3PolicyWithClusterId(clusterId int64) *HTTP3Policy {
	http3PolicyLocker.RLock()
	defer http3PolicyLocker.RUnlock()
	if this.HTTP3Policies == nil {
		return nil
	}
	return this.HTTP3Policies[clusterId]
}

// FindHTTP3Ports 查询HTTP/3所有端口
func (this *NodeConfig) FindHTTP3Ports() (ports []int) {
	http3PolicyLocker.RLock()
	defer http3PolicyLocker.RUnlock()
	for _, policy := range this.HTTP3Policies {
		if !policy.IsOn {
			continue
		}
		if policy.Port <= 0 {
			policy.Port = DefaultHTTP3Port
		}
		if !lists.ContainsInt(ports, policy.Port) {
			ports = append(ports, policy.Port)
		}
	}
	return
}

// UpdateHTTP3Policies 修改集群HTTP/3策略
func (this *NodeConfig) UpdateHTTP3Policies(policies map[int64]*HTTP3Policy) {
	http3PolicyLocker.Lock()
	defer http3PolicyLocker.Unlock()
	this.HTTP3Policies = policies
}

// UpdateHTTPPagesPolicies 修改集群自定义页面策略
func (this *NodeConfig) UpdateHTTPPagesPolicies(policies map[int64]*HTTPPagesPolicy) {
	httpPagesPolicyLocker.Lock()
	defer httpPagesPolicyLocker.Unlock()
	this.HTTPPagesPolicies = policies
}

// FindHTTPPagesPolicyWithClusterId 使用集群ID查找自定义页面策略
func (this *NodeConfig) FindHTTPPagesPolicyWithClusterId(clusterId int64) *HTTPPagesPolicy {
	httpPagesPolicyLocker.RLock()
	defer httpPagesPolicyLocker.RUnlock()
	if this.HTTPPagesPolicies == nil {
		return nil
	}
	return this.HTTPPagesPolicies[clusterId]
}

// SecretHash 对Id和Secret的Hash计算
func (this *NodeConfig) SecretHash() string {
	return this.secretHash
}

// HasConnTimeoutSettings 检查是否有连接超时设置
func (this *NodeConfig) HasConnTimeoutSettings() bool {
	return this.GlobalServerConfig != nil && (this.GlobalServerConfig.Performance.AutoReadTimeout || this.GlobalServerConfig.Performance.AutoWriteTimeout)
}
