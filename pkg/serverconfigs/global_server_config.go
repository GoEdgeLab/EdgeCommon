// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

type LnRequestSchedulingMethod = string

const (
	LnRequestSchedulingMethodRandom     LnRequestSchedulingMethod = "random"
	LnRequestSchedulingMethodURLMapping LnRequestSchedulingMethod = "urlMapping"
)

const (
	DefaultTCPPortRangeMin = 10000
	DefaultTCPPortRangeMax = 40000
)

func NewGlobalServerConfig() *GlobalServerConfig {
	var config = &GlobalServerConfig{}

	config.HTTPAll.SupportsLowVersionHTTP = true
	config.HTTPAll.EnableServerAddrVariable = false
	config.HTTPAll.LnRequestSchedulingMethod = LnRequestSchedulingMethodURLMapping

	config.HTTPAccessLog.IsOn = true
	config.HTTPAccessLog.EnableRequestHeaders = true
	config.HTTPAccessLog.EnableResponseHeaders = true
	config.HTTPAccessLog.EnableCookies = true
	config.HTTPAccessLog.EnableServerNotFound = true

	config.Log.RecordServerError = false

	config.Performance.AutoWriteTimeout = false
	config.Performance.AutoReadTimeout = false
	config.Stat.Upload.MaxCities = 32
	config.Stat.Upload.MaxProviders = 32
	config.Stat.Upload.MaxSystems = 64
	config.Stat.Upload.MaxBrowsers = 64

	return config
}

// GlobalServerConfig 全局的服务配置
type GlobalServerConfig struct {
	HTTPAll struct {
		MatchDomainStrictly  bool                  `yaml:"matchDomainStrictly" json:"matchDomainStrictly"`   // 是否严格匹配域名
		AllowMismatchDomains []string              `yaml:"allowMismatchDomains" json:"allowMismatchDomains"` // 允许的不匹配的域名
		AllowNodeIP          bool                  `yaml:"allowNodeIP" json:"allowNodeIP"`                   // 允许IP直接访问
		NodeIPShowPage       bool                  `yaml:"nodeIPShowPage" json:"nodeIPShowPage"`             // 访问IP地址是否显示页面
		NodeIPPageHTML       string                `yaml:"nodeIPPageHTML" json:"nodeIPPageHTML"`             // 访问IP地址页面内容
		DefaultDomain        string                `yaml:"defaultDomain" json:"defaultDomain"`               // 默认的域名
		DomainMismatchAction *DomainMismatchAction `yaml:"domainMismatchAction" json:"domainMismatchAction"` // 不匹配时采取的动作

		SupportsLowVersionHTTP    bool                      `yaml:"supportsLowVersionHTTP" json:"supportsLowVersionHTTP"`       // 是否启用低版本HTTP
		MatchCertFromAllServers   bool                      `yaml:"matchCertFromAllServers" json:"matchCertFromAllServers"`     // 从所有服务中匹配证书（不要轻易开启！）
		ForceLnRequest            bool                      `yaml:"forceLnRequest" json:"forceLnRequest"`                       // 强制从Ln请求内容
		LnRequestSchedulingMethod LnRequestSchedulingMethod `yaml:"lnRequestSchedulingMethod" json:"lnRequestSchedulingMethod"` // Ln请求调度方法
		ServerName                string                    `yaml:"serverName" json:"serverName"`                               // Server名称
		EnableServerAddrVariable  bool                      `yaml:"enableServerAddrVariable" json:"enableServerAddrVariable"`   // 是否支持${serverAddr}变量

		DomainAuditingIsOn   bool   `yaml:"domainAuditingIsOn" json:"domainAuditingIsOn"`     // 域名是否需要审核
		DomainAuditingPrompt string `yaml:"domainAuditingPrompt" json:"domainAuditingPrompt"` // 域名审核的提示
	} `yaml:"httpAll" json:"httpAll"` // HTTP统一配置

	TCPAll struct {
		PortRangeMin int   `yaml:"portRangeMin" json:"portRangeMin"` // 最小端口
		PortRangeMax int   `yaml:"portRangeMax" json:"portRangeMax"` // 最大端口
		DenyPorts    []int `yaml:"denyPorts" json:"denyPorts"`       // 禁止使用的端口
	} `yaml:"tcpAll" json:"tcpAll"`

	HTTPAccessLog struct {
		IsOn                     bool `yaml:"isOn" json:"isOn"`                                         // 是否启用此功能
		EnableRequestHeaders     bool `yaml:"enableRequestHeaders" json:"enableRequestHeaders"`         // 记录请求Header
		CommonRequestHeadersOnly bool `yaml:"commonRequestHeadersOnly" json:"commonRequestHeadersOnly"` // 只保留通用Header
		EnableResponseHeaders    bool `yaml:"enableResponseHeaders" json:"enableResponseHeaders"`       // 记录响应Header
		EnableCookies            bool `yaml:"enableCookies" json:"enableCookies"`                       // 记录Cookie
		EnableServerNotFound     bool `yaml:"enableServerNotFound" json:"enableServerNotFound"`         // 记录服务找不到的日志
	} `yaml:"httpAccessLog" json:"httpAccessLog"` // 访问日志配置

	Stat struct {
		Upload struct {
			MaxCities    int16 `yaml:"maxCities" json:"maxCities"`       // 最大区域数量
			MaxProviders int16 `yaml:"maxProviders" json:"maxProviders"` // 最大运营商数量
			MaxSystems   int16 `yaml:"maxSystems" json:"maxSystems"`     // 最大操作系统数量
			MaxBrowsers  int16 `yaml:"maxBrowsers" json:"maxBrowsers"`   // 最大浏览器数量
		} `yaml:"upload" json:"upload"` // 上传相关设置
	} `yaml:"stat" json:"stat"` // 统计相关配置

	Performance struct {
		Debug            bool `yaml:"debug" json:"debug"`                       // Debug模式
		AutoWriteTimeout bool `yaml:"autoWriteTimeout" json:"autoWriteTimeout"` // 是否自动写超时
		AutoReadTimeout  bool `yaml:"autoReadTimeout" json:"autoReadTimeout"`   // 是否自动读超时
	} `yaml:"performance" json:"performance"` // 性能

	Log struct {
		RecordServerError bool `yaml:"recordServerError" json:"recordServerError"` // 记录服务错误到运行日志
	} `yaml:"log" json:"log"` // 运行日志配置
}

func (this *GlobalServerConfig) Init() error {
	// 未找到域名时的动作
	if this.HTTPAll.DomainMismatchAction != nil {
		err := this.HTTPAll.DomainMismatchAction.Init()
		if err != nil {
			return err
		}
	}

	return nil
}
