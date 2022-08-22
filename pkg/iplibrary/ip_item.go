// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"encoding/binary"
	"github.com/iwind/TeaGo/types"
)

type ipItem struct {
	IPFrom uint64
	IPTo   uint64

	Region *ipRegion
}

type ipRegion struct {
	CountryId  uint32
	ProvinceId uint32
	CityId     uint32
	TownId     uint32
	ProviderId uint32
}

func (this *ipItem) AsBinary() ([]byte, error) {
	var buf = &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, this)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func HashRegion(countryId uint32, provinceId uint32, cityId uint32, townId uint32, providerId uint32) string {
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
