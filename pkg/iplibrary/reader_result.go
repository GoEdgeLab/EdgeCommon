// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"github.com/iwind/TeaGo/lists"
	"strings"
)

type QueryResult struct {
	item *ipItem
	meta *Meta
}

func (this *QueryResult) IsOk() bool {
	return this.item != nil
}

func (this *QueryResult) CountryId() int64 {
	if this.item != nil {
		return int64(this.item.CountryId)
	}
	return 0
}

func (this *QueryResult) CountryName() string {
	if this.item == nil {
		return ""
	}
	if this.item.CountryId > 0 {
		var country = this.meta.CountryWithId(this.item.CountryId)
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
	if this.item.CountryId > 0 {
		var country = this.meta.CountryWithId(this.item.CountryId)
		if country != nil {
			return country.Codes
		}
	}
	return nil
}

func (this *QueryResult) ProvinceId() int64 {
	if this.item != nil {
		return int64(this.item.ProvinceId)
	}
	return 0
}

func (this *QueryResult) ProvinceName() string {
	if this.item == nil {
		return ""
	}
	if this.item.ProvinceId > 0 {
		var province = this.meta.ProvinceWithId(this.item.ProvinceId)
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
	if this.item.ProvinceId > 0 {
		var province = this.meta.ProvinceWithId(this.item.ProvinceId)
		if province != nil {
			return province.Codes
		}
	}
	return nil
}

func (this *QueryResult) CityId() int64 {
	if this.item != nil {
		return int64(this.item.CityId)
	}
	return 0
}

func (this *QueryResult) CityName() string {
	if this.item == nil {
		return ""
	}
	if this.item.CityId > 0 {
		var city = this.meta.CityWithId(this.item.CityId)
		if city != nil {
			return city.Name
		}
	}
	return ""
}

func (this *QueryResult) TownId() int64 {
	if this.item != nil {
		return int64(this.item.TownId)
	}
	return 0
}

func (this *QueryResult) TownName() string {
	if this.item == nil {
		return ""
	}
	if this.item.TownId > 0 {
		var town = this.meta.TownWithId(this.item.TownId)
		if town != nil {
			return town.Name
		}
	}
	return ""
}

func (this *QueryResult) ProviderId() int64 {
	if this.item != nil {
		return int64(this.item.ProviderId)
	}
	return 0
}

func (this *QueryResult) ProviderName() string {
	if this.item == nil {
		return ""
	}
	if this.item.ProviderId > 0 {
		var provider = this.meta.ProviderWithId(this.item.ProviderId)
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
	if this.item.ProviderId > 0 {
		var provider = this.meta.ProviderWithId(this.item.ProviderId)
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
		pieces = append(pieces, cityName)
	}

	if len(providerName) > 0 && !lists.ContainsString(pieces, providerName) {
		if len(pieces) > 0 {
			pieces = append(pieces, "|")
		}
		pieces = append(pieces, providerName)
	}

	return strings.Join(pieces, " ")
}
