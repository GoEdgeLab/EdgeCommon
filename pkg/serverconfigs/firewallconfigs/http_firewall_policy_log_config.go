// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package firewallconfigs

var DefaultHTTPFirewallPolicyLogConfig = &HTTPFirewallPolicyLogConfig{
	IsOn:          true,
	RequestBody:   true,
	RegionDenying: false,
}

type HTTPFirewallPolicyLogConfig struct {
	IsPrior       bool `yaml:"isPrior" json:"isPrior"`
	IsOn          bool `yaml:"isOn" json:"isOn"`
	RequestBody   bool `yaml:"requestBody" json:"requestBody"`     // 是否记录RequestBody
	RegionDenying bool `yaml:"regionDenying" json:"regionDenying"` // 是否记录区域封禁日志
}

func (this *HTTPFirewallPolicyLogConfig) Init() error {
	return nil
}
