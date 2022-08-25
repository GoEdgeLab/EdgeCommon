// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package firewallconfigs

type HTTPFirewallJavascriptCookieAction struct {
	Life             int32 `yaml:"life" json:"life"`                         // 有效期
	MaxFails         int   `yaml:"maxFails" json:"maxFails"`                 // 最大失败次数
	FailBlockTimeout int   `yaml:"failBlockTimeout" json:"failBlockTimeout"` // 失败拦截时间
}
