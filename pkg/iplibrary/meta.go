// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

type Meta struct {
	Version   int         `json:"version"` // IP库版本
	Author    string      `json:"author"`
	Countries []*Country  `json:"countries"`
	Provinces []*Province `json:"provinces"`
	Cities    []*City     `json:"cities"`
	Towns     []*Town     `json:"towns"`
	Providers []*Provider `json:"providers"`
	CreatedAt int64       `json:"createdAt"`
}

type Provider struct {
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

type Country struct {
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

type Province struct {
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

type City struct {
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

type Town struct {
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}
