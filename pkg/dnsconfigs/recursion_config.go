// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type DNSHost struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

// RecursionConfig 递归DNS设置
type RecursionConfig struct {
	IsOn          bool       `json:"isOn"`
	Hosts         []*DNSHost `json:"hosts"`
	UseLocalHosts bool       `json:"useLocalHosts"` // 自动从本机读取DNS
	AllowDomains  []string   `json:"allowDomains"`
	DenyDomains   []string   `json:"denyDomains"`
}
