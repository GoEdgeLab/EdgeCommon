// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/configutils"

type RouteRangeType = string

const (
	RouteRangeTypeIP RouteRangeType = "ipRange"
)

type RouteRangeInterface interface {
	Init() error
	Contains(ip uint64) bool
}

// RouteRangeIPRange IP范围配置
type RouteRangeIPRange struct {
	IPFrom string `json:"ipFrom"`
	IPTo   string `json:"ipTo"`

	ipFromLong uint64
	ipToLong   uint64
}

func (this *RouteRangeIPRange) Init() error {
	this.ipFromLong = configutils.IP2Long(this.IPFrom)
	this.ipToLong = configutils.IP2Long(this.IPTo)

	if this.ipFromLong > this.ipToLong {
		this.ipFromLong, this.ipToLong = this.ipToLong, this.ipFromLong
	}

	return nil
}

func (this *RouteRangeIPRange) Contains(ip uint64) bool {
	return this.ipFromLong <= ip && this.ipToLong >= ip
}
