// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dnsconfigs

func DefaultNSUserConfig() *NSUserConfig {
	return &NSUserConfig{
		DefaultClusterId: 0,
	}
}

type NSUserConfig struct {
	DefaultClusterId int64 `json:"defaultClusterId"` // 默认部署到的集群
}
