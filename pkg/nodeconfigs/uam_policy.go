// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package nodeconfigs

func NewUAMPolicy() *UAMPolicy {
	return &UAMPolicy{}
}

type UAMPolicy struct {
	IsOn bool `yaml:"isOn" json:"isOn"`
}

func (this *UAMPolicy) Init() error {
	return nil
}
