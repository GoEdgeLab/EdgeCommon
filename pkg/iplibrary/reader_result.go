// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

type QueryResult struct {
	item *ipItem
	meta *Meta
}

func (this *QueryResult) IsOk() bool {
	return this.item != nil
}

func (this *QueryResult) CountryId() int64 {
	if this.item != nil {
		return this.item.countryId
	}
	return 0
}

func (this *QueryResult) CountryName() string {
	if this.item.countryId > 0 {
		var country = this.meta.CountryWithId(this.item.countryId)
		if country != nil {
			return country.Name
		}
	}
	return ""
}

func (this *QueryResult) ProvinceId() int64 {
	if this.item != nil {
		return this.item.provinceId
	}
	return 0
}

func (this *QueryResult) ProvinceName() string {
	if this.item.provinceId > 0 {
		var province = this.meta.ProvinceWithId(this.item.provinceId)
		if province != nil {
			return province.Name
		}
	}
	return ""
}

func (this *QueryResult) CityId() int64 {
	if this.item != nil {
		return this.item.cityId
	}
	return 0
}

func (this *QueryResult) CityName() string {
	if this.item.cityId > 0 {
		var city = this.meta.CityWithId(this.item.cityId)
		if city != nil {
			return city.Name
		}
	}
	return ""
}

func (this *QueryResult) TownId() int64 {
	if this.item != nil {
		return this.item.townId
	}
	return 0
}

func (this *QueryResult) TownName() string {
	if this.item.townId > 0 {
		var town = this.meta.TownWithId(this.item.townId)
		if town != nil {
			return town.Name
		}
	}
	return ""
}

func (this *QueryResult) ProviderId() int64 {
	if this.item != nil {
		return this.item.providerId
	}
	return 0
}

func (this *QueryResult) ProviderName() string {
	if this.item.providerId > 0 {
		var provider = this.meta.ProviderWithId(this.item.providerId)
		if provider != nil {
			return provider.Name
		}
	}
	return ""
}
