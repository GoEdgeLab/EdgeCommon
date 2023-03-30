// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"github.com/iwind/TeaGo/types"
)

type ipv4Item struct {
	IPFrom uint32
	IPTo   uint32

	Region *ipRegion
}

type ipv6Item struct {
	IPFrom uint64
	IPTo   uint64

	Region *ipRegion
}

type ipRegion struct {
	CountryId  uint16
	ProvinceId uint16
	CityId     uint32
	TownId     uint32
	ProviderId uint16
}

func HashRegion(countryId uint16, provinceId uint16, cityId uint32, townId uint32, providerId uint16) string {
	var providerHash = ""
	if providerId > 0 {
		providerHash = "_" + types.String(providerId)
	}

	if townId > 0 {
		return "t" + types.String(townId) + providerHash
	}
	if cityId > 0 {
		return "c" + types.String(cityId) + providerHash
	}
	if provinceId > 0 {
		return "p" + types.String(provinceId) + providerHash
	}
	return "a" + types.String(countryId) + providerHash
}
