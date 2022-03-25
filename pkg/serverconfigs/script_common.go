// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

// CommonScript 公共脚本
type CommonScript struct {
	Id       int64  `yaml:"id" json:"id"`
	IsOn     bool   `yaml:"isOn" json:"isOn"`
	Filename string `yaml:"filename" json:"filename"`
	Code     string `yaml:"code" json:"code"`
}
