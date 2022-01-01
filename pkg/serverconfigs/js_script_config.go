// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type JSScriptConfig struct {
	IsPrior bool   `yaml:"isPrior" json:"isPrior"`
	IsOn    bool   `yaml:"isOn" json:"isOn"`
	Code    string `yaml:"code" json:"code"`
}

func (this *JSScriptConfig) Init() error {
	return nil
}
