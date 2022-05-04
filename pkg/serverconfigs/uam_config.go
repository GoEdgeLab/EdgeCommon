// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type UAMConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`
	IsOn    bool `yaml:"isOn" json:"isOn"`
}

func (this *UAMConfig) Init() error {
	return nil
}
