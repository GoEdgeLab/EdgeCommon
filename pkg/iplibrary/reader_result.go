// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"github.com/iwind/TeaGo/lists"
	"strings"
)

type QueryResult struct {
	item any
	meta *Meta
}

func (this *QueryResult) IsOk() bool {
	return this.item != nil
}

func (this *QueryResult) CountryId() int64 {
	return int64(this.realCountryId())
}

func (this *QueryResult) CountryName() string {
	if this.item == nil {
		return ""
	}
	var countryId = this.realCountryId()
	if countryId > 0 {
		var country = this.meta.CountryWithId(countryId)
		if country != nil {
			return country.Name
		}
	}
	return ""
}

func (this *QueryResult) CountryCodes() []string {
	if this.item == nil {
		return nil
	}
	var countryId = this.realCountryId()
	if countryId > 0 {
		var country = this.meta.CountryWithId(countryId)
		if country != nil {
			return country.Codes
		}
	}
	return nil
}

func (this *QueryResult) ProvinceId() int64 {
	return int64(this.realProvinceId())
}

func (this *QueryResult) ProvinceName() string {
	if this.item == nil {
		return ""
	}
	var provinceId = this.realProvinceId()
	if provinceId > 0 {
		var province = this.meta.ProvinceWithId(provinceId)
		if province != nil {
			return province.Name
		}
	}
	return ""
}

func (this *QueryResult) ProvinceCodes() []string {
	if this.item == nil {
		return nil
	}
	var provinceId = this.realProvinceId()
	if provinceId > 0 {
		var province = this.meta.ProvinceWithId(provinceId)
		if province != nil {
			return province.Codes
		}
	}
	return nil
}

func (this *QueryResult) CityId() int64 {
	return int64(this.realCityId())
}

func (this *QueryResult) CityName() string {
	if this.item == nil {
		return ""
	}
	var cityId = this.realCityId()
	if cityId > 0 {
		var city = this.meta.CityWithId(cityId)
		if city != nil {
			return city.Name
		}
	}
	return ""
}

func (this *QueryResult) TownId() int64 {
	return int64(this.realTownId())
}

func (this *QueryResult) TownName() string {
	if this.item == nil {
		return ""
	}
	var townId = this.realTownId()
	if townId > 0 {
		var town = this.meta.TownWithId(townId)
		if town != nil {
			return town.Name
		}
	}
	return ""
}

func (this *QueryResult) ProviderId() int64 {
	return int64(this.realProviderId())
}

func (this *QueryResult) ProviderName() string {
	if this.item == nil {
		return ""
	}
	var providerId = this.realProviderId()
	if providerId > 0 {
		var provider = this.meta.ProviderWithId(providerId)
		if provider != nil {
			return provider.Name
		}
	}
	return ""
}

func (this *QueryResult) ProviderCodes() []string {
	if this.item == nil {
		return nil
	}
	var providerId = this.realProviderId()
	if providerId > 0 {
		var provider = this.meta.ProviderWithId(providerId)
		if provider != nil {
			return provider.Codes
		}
	}
	return nil
}

func (this *QueryResult) Summary() string {
	if this.item == nil {
		return ""
	}

	var pieces = []string{}
	var countryName = this.CountryName()
	var provinceName = this.ProvinceName()
	var cityName = this.CityName()
	var townName = this.TownName()
	var providerName = this.ProviderName()

	if len(countryName) > 0 {
		pieces = append(pieces, countryName)
	}
	if len(provinceName) > 0 && !lists.ContainsString(pieces, provinceName) {
		pieces = append(pieces, provinceName)
	}
	if len(cityName) > 0 && !lists.ContainsString(pieces, cityName) && !lists.ContainsString(pieces, strings.TrimSuffix(cityName, "市")) {
		pieces = append(pieces, cityName)
	}
	if len(townName) > 0 && !lists.ContainsString(pieces, townName) && !lists.ContainsString(pieces, strings.TrimSuffix(townName, "县")) {
		pieces = append(pieces, townName)
	}

	if len(providerName) > 0 && !lists.ContainsString(pieces, providerName) {
		if len(pieces) > 0 {
			pieces = append(pieces, "|")
		}
		pieces = append(pieces, providerName)
	}

	return strings.Join(pieces, " ")
}

func (this *QueryResult) realCountryId() uint16 {
	if this.item != nil {
		switch item := this.item.(type) {
		case *ipv4Item:
			return item.Region.CountryId
		case *ipv6Item:
			return item.Region.CountryId
		}

	}
	return 0
}

func (this *QueryResult) realProvinceId() uint16 {
	if this.item != nil {
		switch item := this.item.(type) {
		case *ipv4Item:
			return item.Region.ProvinceId
		case *ipv6Item:
			return item.Region.ProvinceId
		}

	}
	return 0
}

func (this *QueryResult) realCityId() uint32 {
	if this.item != nil {
		switch item := this.item.(type) {
		case *ipv4Item:
			return item.Region.CityId
		case *ipv6Item:
			return item.Region.CityId
		}

	}
	return 0
}

func (this *QueryResult) realTownId() uint32 {
	if this.item != nil {
		switch item := this.item.(type) {
		case *ipv4Item:
			return item.Region.TownId
		case *ipv6Item:
			return item.Region.TownId
		}

	}
	return 0
}

func (this *QueryResult) realProviderId() uint16 {
	if this.item != nil {
		switch item := this.item.(type) {
		case *ipv4Item:
			return item.Region.ProviderId
		case *ipv6Item:
			return item.Region.ProviderId
		}

	}
	return 0
}
