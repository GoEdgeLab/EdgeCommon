// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

type QueryResult struct {
	item *ipItem
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

func (this *QueryResult) ProvinceId() int64 {
	if this.item != nil {
		return this.item.provinceId
	}
	return 0
}

func (this *QueryResult) CityId() int64 {
	if this.item != nil {
		return this.item.cityId
	}
	return 0
}

func (this *QueryResult) TownId() int64 {
	if this.item != nil {
		return this.item.townId
	}
	return 0
}

func (this *QueryResult) ProviderId() int64 {
	if this.item != nil {
		return this.item.providerId
	}
	return 0
}
