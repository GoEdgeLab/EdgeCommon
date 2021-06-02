// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type AccessLogRef struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"` // 是否覆盖
	IsOn    bool `yaml:"isOn" json:"isOn"`       // 是否启用
}

func (this *AccessLogRef) Init() error {
	return nil
}
