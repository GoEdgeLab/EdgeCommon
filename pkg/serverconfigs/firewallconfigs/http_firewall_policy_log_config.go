// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package firewallconfigs

var DefaultHTTPFirewallPolicyLogConfig = &HTTPFirewallPolicyLogConfig{
	IsOn:        true,
	RequestBody: true,
}

type HTTPFirewallPolicyLogConfig struct {
	IsOn        bool `yaml:"isOn" json:"isOn"`
	RequestBody bool `yaml:"requestBody" json:"requestBody"`
}

func (this *HTTPFirewallPolicyLogConfig) Init() error {
	return nil
}
