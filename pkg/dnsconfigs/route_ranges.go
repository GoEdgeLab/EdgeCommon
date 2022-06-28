// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

import (
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
	"net"
)

type RouteRangeType = string

const (
	RouteRangeTypeIP     RouteRangeType = "ipRange" // IP范围
	RouteRangeTypeCIDR   RouteRangeType = "cidr"    // CIDR
	RouteRangeTypeRegion RouteRangeType = "region"  // 区域
)

func AllRouteRangeTypes() []*shared.Definition {
	return []*shared.Definition{
		{
			Name: "IP范围",
			Code: RouteRangeTypeIP,
		},
		{
			Name: "CIDR",
			Code: RouteRangeTypeCIDR,
		},
		{
			Name: "区域",
			Code: RouteRangeTypeRegion,
		},
	}
}

// RouteRegionResolver 解析IP接口
type RouteRegionResolver interface {
	Resolve(ip net.IP) (countryId int64, provinceId int64, cityId int64, providerId int64)
}

// RouteRangeInterface 线路范围接口
type RouteRangeInterface interface {
	// Init 初始化
	Init() error

	// Contains 判断是否包含
	Contains(ip net.IP) bool

	// SetRegionResolver 设置IP解析接口
	SetRegionResolver(resolver RouteRegionResolver)

	// IsExcluding 是否为排除
	IsExcluding() bool
}

type BaseRouteRange struct {
	IsReverse bool `json:"isReverse"`

	routeRegionResolver RouteRegionResolver
}

func (this *BaseRouteRange) SetRegionResolver(resolver RouteRegionResolver) {
	this.routeRegionResolver = resolver
}

func (this *BaseRouteRange) IsExcluding() bool {
	return this.IsReverse
}

// RouteRangeIPRange IP范围配置
// IPv4和IPv6不能混用
type RouteRangeIPRange struct {
	BaseRouteRange

	IPFrom string `json:"ipFrom"`
	IPTo   string `json:"ipTo"`

	ipFromLong uint64
	ipToLong   uint64

	ipVersion int // 4|6
}

func (this *RouteRangeIPRange) Init() error {
	var ipFrom = net.ParseIP(this.IPFrom)
	var ipTo = net.ParseIP(this.IPTo)
	if ipFrom == nil {
		return errors.New("invalid ipFrom '" + this.IPFrom + "'")
	}
	if ipTo == nil {
		return errors.New("invalid ipTo '" + this.IPTo + "'")
	}

	var ipFromVersion = configutils.IPVersion(ipFrom)
	var ipToVersion = configutils.IPVersion(ipTo)
	if ipFromVersion != ipToVersion {
		return errors.New("ipFrom and ipTo version are not same")
	}
	this.ipVersion = ipFromVersion

	this.ipFromLong = configutils.IP2Long(ipFrom)
	this.ipToLong = configutils.IP2Long(ipTo)

	if this.ipFromLong > this.ipToLong {
		this.ipFromLong, this.ipToLong = this.ipToLong, this.ipFromLong
	}

	return nil
}

func (this *RouteRangeIPRange) Contains(netIP net.IP) bool {
	if len(netIP) == 0 {
		return false
	}

	var version = configutils.IPVersion(netIP)
	if version != this.ipVersion {
		return false
	}

	var ipLong = configutils.IP2Long(netIP)
	return ipLong >= this.ipFromLong && ipLong <= this.ipToLong
}

// RouteRangeCIDR CIDR范围配置
type RouteRangeCIDR struct {
	BaseRouteRange

	CIDR string `json:"cidr"`

	cidr *net.IPNet
}

func (this *RouteRangeCIDR) Init() error {
	_, ipNet, err := net.ParseCIDR(this.CIDR)
	if err != nil {
		return errors.New("parse cidr failed: " + err.Error())
	}

	this.cidr = ipNet

	return nil
}

func (this *RouteRangeCIDR) Contains(netIP net.IP) bool {
	if netIP == nil {
		return false
	}

	if this.cidr == nil {
		return false
	}

	return this.cidr.Contains(netIP)
}

// RouteRangeRegion 区域范围
// country:ID, province:ID, city:ID, isp:ID
type RouteRangeRegion struct {
	BaseRouteRange

	Regions []*routeRegion `json:"regions"`
}

func (this *RouteRangeRegion) Init() error {
	return nil
}

func (this *RouteRangeRegion) Contains(netIP net.IP) bool {
	if this.routeRegionResolver == nil {
		return false
	}

	if len(this.Regions) == 0 {
		return false
	}

	countryId, provinceId, cityId, providerId := this.routeRegionResolver.Resolve(netIP)
	if countryId <= 0 && provinceId <= 0 && cityId <= 0 && providerId <= 0 {
		return false
	}

	for _, region := range this.Regions {
		if region.Id <= 0 {
			continue
		}

		switch region.Type {
		case "country":
			if region.Id == countryId {
				return true
			}
		case "province":
			if region.Id == provinceId {
				return true
			}
		case "city":
			if region.Id == cityId {
				return true
			}
		case "isp":
			if region.Id == providerId {
				return true
			}
		}
	}

	return false
}

type routeRegion struct {
	Type string `json:"type"` // country|province|city|isp
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// InitRangesFromJSON 从JSON中初始化线路范围
func InitRangesFromJSON(rangesJSON []byte) (ranges []RouteRangeInterface, err error) {
	if len(rangesJSON) == 0 {
		return
	}

	var rangeMaps = []maps.Map{}
	err = json.Unmarshal(rangesJSON, &rangeMaps)
	if err != nil {
		return nil, err
	}
	for _, rangeMap := range rangeMaps {
		var rangeType = rangeMap.GetString("type")
		paramsJSON, err := json.Marshal(rangeMap.Get("params"))
		if err != nil {
			return nil, err
		}

		var r RouteRangeInterface

		switch rangeType {
		case RouteRangeTypeIP:
			r = &RouteRangeIPRange{}
		case RouteRangeTypeCIDR:
			r = &RouteRangeCIDR{}
		case RouteRangeTypeRegion:
			r = &RouteRangeRegion{}
		default:
			return nil, errors.New("invalid route line type '" + rangeType + "'")
		}

		err = json.Unmarshal(paramsJSON, r)
		if err != nil {
			return nil, err
		}
		err = r.Init()
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, r)
	}
	return
}
