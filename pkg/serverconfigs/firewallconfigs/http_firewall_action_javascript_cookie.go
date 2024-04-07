// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package firewallconfigs

type HTTPFirewallJavascriptCookieAction struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	Life              int32  `yaml:"life" json:"life"`                         // 有效期
	MaxFails          int    `yaml:"maxFails" json:"maxFails"`                 // 最大失败次数
	FailBlockTimeout  int    `yaml:"failBlockTimeout" json:"failBlockTimeout"` // 失败拦截时间
	Scope             string `yaml:"scope" json:"scope"`
	FailBlockScopeAll bool   `yaml:"failBlockScopeAll" json:"failBlockScopeAll"`
}

func NewHTTPFirewallJavascriptCookieAction() *HTTPFirewallJavascriptCookieAction {
	return &HTTPFirewallJavascriptCookieAction{
		Life:              600,
		MaxFails:          100,
		FailBlockTimeout:  3600,
		Scope:             FirewallScopeServer,
		FailBlockScopeAll: true,
	}
}
