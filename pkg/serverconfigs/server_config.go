package serverconfigs

import (
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
)

type ServerConfig struct {
	Id               int64               `yaml:"id" json:"id"`                             // ID
	ClusterId        int64               `yaml:"clusterId" json:"clusterId"`               // 集群ID
	Type             string              `yaml:"type" json:"type"`                         // 类型
	IsOn             bool                `yaml:"isOn" json:"isOn"`                         // 是否开启
	Name             string              `yaml:"name" json:"name"`                         // 名称
	Description      string              `yaml:"description" json:"description"`           // 描述
	AliasServerNames []string            `yaml:"aliasServerNames" json:"aliasServerNames"` // 关联的域名，比如CNAME之类的
	ServerNames      []*ServerNameConfig `yaml:"serverNames" json:"serverNames"`           // 域名
	SupportCNAME     bool                `yaml:"supportCNAME" json:"supportCNAME"`         // 是否支持CNAME

	// 前端协议
	HTTP  *HTTPProtocolConfig  `yaml:"http" json:"http"`   // HTTP配置
	HTTPS *HTTPSProtocolConfig `yaml:"https" json:"https"` // HTTPS配置
	TCP   *TCPProtocolConfig   `yaml:"tcp" json:"tcp"`     // TCP配置
	TLS   *TLSProtocolConfig   `yaml:"tls" json:"tls"`     // TLS配置
	Unix  *UnixProtocolConfig  `yaml:"unix" json:"unix"`   // Unix配置
	UDP   *UDPProtocolConfig   `yaml:"udp" json:"udp"`     // UDP配置

	// Web配置
	Web *HTTPWebConfig `yaml:"web" json:"web"`

	// 反向代理配置
	ReverseProxyRef *ReverseProxyRef    `yaml:"reverseProxyRef" json:"reverseProxyRef"`
	ReverseProxy    *ReverseProxyConfig `yaml:"reverseProxy" json:"reverseProxy"`

	// WAF策略
	HTTPFirewallPolicyId int64                               `yaml:"httpFirewallPolicyId" json:"httpFirewallPolicyId"`
	HTTPFirewallPolicy   *firewallconfigs.HTTPFirewallPolicy `yaml:"httpFirewallPolicy" json:"httpFirewallPolicy"` // 通过 HTTPFirewallPolicyId 获取

	// 缓存策略
	HTTPCachePolicyId int64            `yaml:"httpCachePolicyId" json:"httpCachePolicyId"`
	HTTPCachePolicy   *HTTPCachePolicy `yaml:"httpCachePolicy" json:"httpCachePolicy"` // 通过 HTTPCachePolicyId 获取

	// 流量限制
	TrafficLimit       *TrafficLimitConfig `yaml:"trafficLimit" json:"trafficLimit"`
	TrafficLimitStatus *TrafficLimitStatus `yaml:"trafficLimitStatus" json:"trafficLimitStatus"`

	// 套餐
	UserPlan *UserPlanConfig `yaml:"userPlan" json:"userPlan"`

	// 分组
	Group *ServerGroupConfig `yaml:"group" json:"group"`

	// UAM
	UAM *UAMConfig `yaml:"uam" json:"uam"`

	isOk bool

	planId int64
}

// NewServerConfigFromJSON 从JSON中解析Server配置
func NewServerConfigFromJSON(data []byte) (*ServerConfig, error) {
	config := &ServerConfig{}
	err := json.Unmarshal(data, config)
	return config, err
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{}
}

func (this *ServerConfig) Init() (results []error) {
	// 分解Group
	if this.Group != nil && this.Group.IsOn {
		// reverse proxy
		if this.IsHTTPFamily() && this.Group.HTTPReverseProxyRef != nil && this.Group.HTTPReverseProxyRef.IsPrior {
			this.ReverseProxyRef = this.Group.HTTPReverseProxyRef
			this.ReverseProxy = this.Group.HTTPReverseProxy
		}
		if this.IsTCPFamily() && this.Group.TCPReverseProxyRef != nil && this.Group.TCPReverseProxyRef.IsPrior {
			this.ReverseProxyRef = this.Group.TCPReverseProxyRef
			this.ReverseProxy = this.Group.TCPReverseProxy
		}
		if this.IsUDPFamily() && this.Group.UDPReverseProxyRef != nil && this.Group.UDPReverseProxyRef.IsPrior {
			this.ReverseProxyRef = this.Group.UDPReverseProxyRef
			this.ReverseProxy = this.Group.UDPReverseProxy
		}

		// web
		if this.Group.Web != nil {
			if this.Web == nil {
				this.Web = this.Group.Web
			} else {
				var groupWeb = this.Group.Web

				// root
				if groupWeb.Root != nil && groupWeb.Root.IsPrior {
					this.Web.Root = groupWeb.Root
				}

				// waf
				if groupWeb.FirewallRef != nil && groupWeb.FirewallRef.IsPrior {
					this.Web.FirewallRef = groupWeb.FirewallRef
					this.Web.FirewallPolicy = groupWeb.FirewallPolicy
				}

				// cache
				if groupWeb.Cache != nil && groupWeb.Cache.IsPrior {
					this.Web.Cache = groupWeb.Cache
				}

				// charset
				if groupWeb.Charset != nil && groupWeb.Charset.IsPrior {
					this.Web.Charset = groupWeb.Charset
				}

				// accessLog
				if groupWeb.AccessLogRef != nil && groupWeb.AccessLogRef.IsPrior {
					this.Web.AccessLogRef = groupWeb.AccessLogRef
				}

				// stat
				if groupWeb.StatRef != nil && groupWeb.StatRef.IsPrior {
					this.Web.StatRef = groupWeb.StatRef
				}

				// compression
				if groupWeb.Compression != nil && groupWeb.Compression.IsPrior {
					this.Web.Compression = groupWeb.Compression
				}

				// headers
				if groupWeb.RequestHeaderPolicyRef != nil && groupWeb.RequestHeaderPolicyRef.IsPrior {
					this.Web.RequestHeaderPolicyRef = groupWeb.RequestHeaderPolicyRef
					this.Web.RequestHeaderPolicy = groupWeb.RequestHeaderPolicy
				}
				if groupWeb.ResponseHeaderPolicyRef != nil && groupWeb.ResponseHeaderPolicyRef.IsPrior {
					this.Web.ResponseHeaderPolicyRef = groupWeb.ResponseHeaderPolicyRef
					this.Web.ResponseHeaderPolicy = groupWeb.ResponseHeaderPolicy
				}

				// websocket
				if groupWeb.WebsocketRef != nil && groupWeb.WebsocketRef.IsPrior {
					this.Web.WebsocketRef = groupWeb.WebsocketRef
					this.Web.Websocket = groupWeb.Websocket
				}

				// webp
				if groupWeb.WebP != nil && groupWeb.WebP.IsPrior {
					this.Web.WebP = groupWeb.WebP
				}

				// remote addr
				if groupWeb.RemoteAddr != nil && groupWeb.RemoteAddr.IsPrior {
					this.Web.RemoteAddr = groupWeb.RemoteAddr
				}

				// pages
				if len(groupWeb.Pages) > 0 || (groupWeb.Shutdown != nil && groupWeb.Shutdown.IsOn) {
					this.Web.Pages = groupWeb.Pages
					this.Web.Shutdown = groupWeb.Shutdown
				}

				// request limit
				if groupWeb.RequestLimit != nil && groupWeb.RequestLimit.IsPrior {
					this.Web.RequestLimit = groupWeb.RequestLimit
				}
			}
		}
	}

	if this.HTTP != nil {
		err := this.HTTP.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	if this.HTTPS != nil {
		err := this.HTTPS.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	if this.TCP != nil {
		err := this.TCP.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	if this.TLS != nil {
		err := this.TLS.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	if this.Unix != nil {
		err := this.Unix.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	if this.UDP != nil {
		err := this.UDP.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	if this.ReverseProxyRef != nil {
		err := this.ReverseProxyRef.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	if this.ReverseProxy != nil {
		err := this.ReverseProxy.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	if this.Web != nil {
		err := this.Web.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	// 套餐
	if this.UserPlan != nil {
		err := this.UserPlan.Init()
		if err != nil {
			results = append(results, err)
		}

		if this.UserPlan.Plan != nil {
			this.planId = this.UserPlan.Plan.Id
		}
	}

	// UAM
	if this.UAM != nil {
		err := this.UAM.Init()
		if err != nil {
			results = append(results, err)
		}
	}

	this.isOk = true

	return nil
}

// IsOk 配置是否正确
func (this *ServerConfig) IsOk() bool {
	return this.isOk
}

func (this *ServerConfig) FullAddresses() []string {
	result := []string{}
	if this.HTTP != nil && this.HTTP.IsOn {
		result = append(result, this.HTTP.FullAddresses()...)
	}
	if this.HTTPS != nil && this.HTTPS.IsOn {
		result = append(result, this.HTTPS.FullAddresses()...)
	}
	if this.TCP != nil && this.TCP.IsOn {
		result = append(result, this.TCP.FullAddresses()...)
	}
	if this.TLS != nil && this.TLS.IsOn {
		result = append(result, this.TLS.FullAddresses()...)
	}
	if this.Unix != nil && this.Unix.IsOn {
		result = append(result, this.Unix.FullAddresses()...)
	}
	if this.UDP != nil && this.UDP.IsOn {
		result = append(result, this.UDP.FullAddresses()...)
	}

	return result
}

func (this *ServerConfig) Listen() []*NetworkAddressConfig {
	result := []*NetworkAddressConfig{}
	if this.HTTP != nil {
		result = append(result, this.HTTP.Listen...)
	}
	if this.HTTPS != nil {
		result = append(result, this.HTTPS.Listen...)
	}
	if this.TCP != nil {
		result = append(result, this.TCP.Listen...)
	}
	if this.TLS != nil {
		result = append(result, this.TLS.Listen...)
	}
	if this.Unix != nil {
		result = append(result, this.Unix.Listen...)
	}
	if this.UDP != nil {
		result = append(result, this.UDP.Listen...)
	}
	return result
}

func (this *ServerConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}

func (this *ServerConfig) IsHTTPFamily() bool {
	return this.HTTP != nil || this.HTTPS != nil
}

func (this *ServerConfig) IsTCPFamily() bool {
	return this.TCP != nil || this.TLS != nil
}

func (this *ServerConfig) IsUnixFamily() bool {
	return this.Unix != nil
}

func (this *ServerConfig) IsUDPFamily() bool {
	return this.UDP != nil
}

// AllStrictNames 所有严格域名
func (this *ServerConfig) AllStrictNames() []string {
	var result = []string{}
	for _, name := range this.AliasServerNames {
		if len(name) > 0 {
			if !configutils.IsFuzzyDomain(name) {
				result = append(result, name)
			}
		}
	}
	for _, serverName := range this.ServerNames {
		var name = serverName.Name
		if len(name) > 0 {
			if !configutils.IsFuzzyDomain(name) {
				result = append(result, name)
			}
		}
		for _, name := range serverName.SubNames {
			if len(name) > 0 {
				if !configutils.IsFuzzyDomain(name) {
					result = append(result, name)
				}
			}
		}
	}
	return result
}

// AllFuzzyNames 所有模糊域名
func (this *ServerConfig) AllFuzzyNames() []string {
	var result = []string{}
	for _, name := range this.AliasServerNames {
		if len(name) > 0 {
			if configutils.IsFuzzyDomain(name) {
				result = append(result, name)
			}
		}
	}
	for _, serverName := range this.ServerNames {
		var name = serverName.Name
		if len(name) > 0 {
			if configutils.IsFuzzyDomain(name) {
				result = append(result, name)
			}
		}
		for _, name := range serverName.SubNames {
			if len(name) > 0 {
				if configutils.IsFuzzyDomain(name) {
					result = append(result, name)
				}
			}
		}
	}
	return result
}

// SSLPolicy SSL信息
func (this *ServerConfig) SSLPolicy() *sslconfigs.SSLPolicy {
	if this.HTTPS != nil {
		return this.HTTPS.SSLPolicy
	}
	if this.TLS != nil {
		return this.TLS.SSLPolicy
	}
	return nil
}

// FindAndCheckReverseProxy 根据条件查找ReverseProxy
func (this *ServerConfig) FindAndCheckReverseProxy(dataType string) (*ReverseProxyConfig, error) {
	switch dataType {
	case "server":
		if this.ReverseProxy == nil {
			return nil, errors.New("reverse proxy not been configured")
		}
		return this.ReverseProxy, nil
	default:
		return nil, errors.New("invalid data type:'" + dataType + "'")
	}
}

// ShouldCheckTrafficLimit 检查是否需要检查流量限制
func (this *ServerConfig) ShouldCheckTrafficLimit() bool {
	return this.TrafficLimit != nil && !this.TrafficLimit.IsEmpty()
}

// PlanId 套餐ID
func (this *ServerConfig) PlanId() int64 {
	return this.planId
}
