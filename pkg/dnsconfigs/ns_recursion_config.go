// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type NSDNSHost struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

// NSRecursionConfig 递归DNS设置
type NSRecursionConfig struct {
	IsOn          bool         `json:"isOn"`
	Hosts         []*NSDNSHost `json:"hosts"`
	UseLocalHosts bool         `json:"useLocalHosts"` // 自动从本机读取DNS
	AllowDomains  []string     `json:"allowDomains"`
	DenyDomains   []string     `json:"denyDomains"`
}

func (this *NSRecursionConfig) Init() error {
	return nil
}
