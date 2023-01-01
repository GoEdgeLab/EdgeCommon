// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

func DefaultGlobalServerConfig() *GlobalServerConfig {
	var config = &GlobalServerConfig{}
	config.HTTPAccessLog.EnableRequestHeaders = true
	config.HTTPAccessLog.EnableResponseHeaders = true
	config.HTTPAccessLog.EnableCookies = true
	config.Log.RecordServerError = false
	config.Performance.AutoWriteTimeout = true
	return config
}

// GlobalServerConfig 全局的服务配置
type GlobalServerConfig struct {
	HTTPAll struct {
		MatchDomainStrictly  bool                  `yaml:"matchDomainStrictly" json:"matchDomainStrictly"`   // 是否严格匹配域名
		AllowMismatchDomains []string              `yaml:"allowMismatchDomains" json:"allowMismatchDomains"` // 允许的不匹配的域名
		AllowNodeIP          bool                  `yaml:"allowNodeIP" json:"allowNodeIP"`                   // 允许IP直接访问
		DefaultDomain        string                `yaml:"defaultDomain" json:"defaultDomain"`               // 默认的域名
		DomainMismatchAction *DomainMismatchAction `yaml:"domainMismatchAction" json:"domainMismatchAction"` // 不匹配时采取的动作
	} `yaml:"httpAll" json:"httpAll"`

	HTTPAccessLog struct {
		EnableRequestHeaders     bool `yaml:"enableRequestHeaders" json:"enableRequestHeaders"`         // 记录请求Header
		CommonRequestHeadersOnly bool `yaml:"commonRequestHeadersOnly" json:"commonRequestHeadersOnly"` // 只保留通用Header
		EnableResponseHeaders    bool `yaml:"enableResponseHeaders" json:"enableResponseHeaders"`       // 记录响应Header
		EnableCookies            bool `yaml:"enableCookies" json:"enableCookies"`                       // 记录Cookie
	} `yaml:"httpAccessLog" json:"httpAccessLog"` // 访问日志配置

	Performance struct {
		Debug            bool `yaml:"debug" json:"debug"`                       // Debug模式
		AutoWriteTimeout bool `yaml:"autoWriteTimeout" json:"autoWriteTimeout"` // 是否自动写超时
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
