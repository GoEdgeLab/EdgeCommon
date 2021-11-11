// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type PlanConfig struct {
	Id int64 `yaml:"id" json:"id"`
}

func (this *PlanConfig) Init() error {
	return nil
}
