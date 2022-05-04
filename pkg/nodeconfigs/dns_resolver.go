// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodeconfigs

type DNSResolverType = string

const (
	DNSResolverTypeDefault  = "default"
	DNSResolverTypeGoNative = "goNative"
	DNSResolverTypeCGO      = "cgo"
)

func DefaultDNSResolverConfig() *DNSResolverConfig {
	return &DNSResolverConfig{
		Type: DNSResolverTypeDefault,
	}
}

type DNSResolverConfig struct {
	Type string `yaml:"type" json:"type"` // 使用Go语言内置的DNS解析器
}

func (this *DNSResolverConfig) Init() error {
	return nil
}
