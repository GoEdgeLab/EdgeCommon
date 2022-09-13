// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dnsconfigs

func DefaultNSUserConfig() *NSUserConfig {
	return &NSUserConfig{
		DefaultClusterId:  0,
		DefaultPlanConfig: DefaultNSUserPlanConfig(),
	}
}

func DefaultNSUserPlanConfig() *NSPlanConfig {
	return &NSPlanConfig{
		SupportCountryRoutes:           true,
		SupportChinaProvinceRoutes:     true,
		SupportISPRoutes:               true,
		MaxCustomRoutes:                0,
		MaxLoadBalanceRecordsPerRecord: 100,
		MinTTL:                         60,
		MaxDomains:                     100,
		MaxRecordsPerDomain:            1000,
		SupportRecordStats:             true,
		SupportDomainAlias:             false,
		SupportAPI:                     false,
	}
}

type NSUserConfig struct {
	DefaultClusterId  int64         `json:"defaultClusterId"`  // 默认部署到的集群
	DefaultPlanConfig *NSPlanConfig `json:"defaultPlanConfig"` // 默认套餐设置
}
