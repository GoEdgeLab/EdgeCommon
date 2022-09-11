// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dnsconfigs

type NSPlanConfig struct {
	SupportCountryRoutes       bool  `json:"supportCountryRoutes"`       // 支持全球国家/地区线路
	SupportChinaProvinceRoutes bool  `json:"supportChinaProvinceRoutes"` // 支持国内省份线路
	SupportISPRoutes           bool  `json:"supportISPRoutes"`           // 支持ISP运营商线路
	CountCustomRoutes          int   `json:"countCustomRoutes"`          // 自定义的线路数量
	CountLoadBalanceRecords    bool  `json:"countLoadBalanceRecords"`    // 负载均衡条数
	MinTTL                     int32 `json:"minTTL"`                     // 最小TTL

	SupportAPI bool `json:"supportAPI"` // 是否支持API操作
}
