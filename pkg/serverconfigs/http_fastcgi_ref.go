// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type HTTPFastcgiRef struct {
	IsPrior    bool    `yaml:"isPrior" json:"isPrior"`       // 是否覆盖
	IsOn       bool    `yaml:"isOn" json:"isOn"`             // 是否开启
	FastcgiIds []int64 `yaml:"fastcgiIds" json:"fastcgiIds"` // Fastcgi ID列表
}
