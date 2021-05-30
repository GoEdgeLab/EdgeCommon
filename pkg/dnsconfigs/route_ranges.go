// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type RouteRangeType = string

const (
	RouteRangeTypeIP RouteRangeType = "ipRange"
)

// RouteRangeIPRange IP范围配置
type RouteRangeIPRange struct {
	IPFrom string `json:"ipFrom"`
	IPTo   string `json:"ipTo"`
}
