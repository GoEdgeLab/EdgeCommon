// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

type Country struct {
	Id    uint32   `json:"id"`
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

type Province struct {
	Id    uint32   `json:"id"`
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

type City struct {
	Id    uint32   `json:"id"`
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

type Town struct {
	Id    uint32   `json:"id"`
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

type Provider struct {
	Id    uint32   `json:"id"`
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

type Meta struct {
	Version   int         `json:"version"` // IP库版本
	Author    string      `json:"author"`
	Countries []*Country  `json:"countries"`
	Provinces []*Province `json:"provinces"`
	Cities    []*City     `json:"cities"`
	Towns     []*Town     `json:"towns"`
	Providers []*Provider `json:"providers"`
	CreatedAt int64       `json:"createdAt"`

	countryMap  map[uint32]*Country  // id => *Country
	provinceMap map[uint32]*Province // id => *Province
	cityMap     map[uint32]*City     // id => *City
	townMap     map[uint32]*Town     // id => *Town
	providerMap map[uint32]*Provider // id => *Provider
}

func (this *Meta) Init() {
	this.countryMap = map[uint32]*Country{}
	this.provinceMap = map[uint32]*Province{}
	this.cityMap = map[uint32]*City{}
	this.townMap = map[uint32]*Town{}
	this.providerMap = map[uint32]*Provider{}

	for _, country := range this.Countries {
		this.countryMap[country.Id] = country
	}
	for _, province := range this.Provinces {
		this.provinceMap[province.Id] = province
	}
	for _, city := range this.Cities {
		this.cityMap[city.Id] = city
	}
	for _, town := range this.Towns {
		this.townMap[town.Id] = town
	}
	for _, provider := range this.Providers {
		this.providerMap[provider.Id] = provider
	}
}

func (this *Meta) CountryWithId(countryId uint32) *Country {
	return this.countryMap[countryId]
}

func (this *Meta) ProvinceWithId(provinceId uint32) *Province {
	return this.provinceMap[provinceId]
}

func (this *Meta) CityWithId(cityId uint32) *City {
	return this.cityMap[cityId]
}

func (this *Meta) TownWithId(townId uint32) *Town {
	return this.townMap[townId]
}

func (this *Meta) ProviderWithId(providerId uint32) *Provider {
	return this.providerMap[providerId]
}
