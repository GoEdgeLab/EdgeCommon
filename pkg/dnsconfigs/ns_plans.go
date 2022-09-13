// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dnsconfigs

type NSPlanConfig struct {
	SupportCountryRoutes           bool  `json:"supportCountryRoutes"`           // 支持全球国家/地区线路
	SupportChinaProvinceRoutes     bool  `json:"supportChinaProvinceRoutes"`     // 支持国内省份线路
	SupportISPRoutes               bool  `json:"supportISPRoutes"`               // 支持ISP运营商线路
	MaxCustomRoutes                int32 `json:"maxCustomRoutes"`                // 自定义的线路数量
	MinTTL                         int32 `json:"minTTL"`                         // 最小TTL
	MaxDomains                     int32 `json:"maxDomains"`                     // 域名数量
	MaxRecordsPerDomain            int32 `json:"maxRecordsPerDomain"`            // 单域名记录数量
	MaxLoadBalanceRecordsPerRecord int32 `json:"maxLoadBalanceRecordsPerRecord"` // 单记录负载均衡条数
	SupportRecordStats             bool  `json:"supportRecordStats"`             // 支持记录统计
	SupportDomainAlias             bool  `json:"supportDomainAlias"`             // 支持域名别名 TODO

	SupportAPI bool `json:"supportAPI"` // 是否支持API操作 TODO
}

func DefaultNSPlanConfig() *NSPlanConfig {
	return &NSPlanConfig{}
}

func (this *NSPlanConfig) Init() error {
	return nil
}
