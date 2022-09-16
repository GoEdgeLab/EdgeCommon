// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

func DefaultGlobalServerConfig() *GlobalServerConfig {
	var config = &GlobalServerConfig{}
	config.Log.RecordServerError = false
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

	Log struct {
		RecordServerError bool `yaml:"recordServerError" json:"recordServerError"` // 记录服务错误到运行日志
	} `yaml:"log" json:"log"`
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
