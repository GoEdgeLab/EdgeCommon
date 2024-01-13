// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// PlanConfig 套餐配置
type PlanConfig struct {
	Id   int64  `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`

	TrafficLimit          *TrafficLimitConfig     `yaml:"trafficLimit" json:"trafficLimit"`
	BandwidthLimitPerNode *shared.BitSizeCapacity `yaml:"bandwidthLimitPerNode" json:"bandwidthLimitPerNode"`
	MaxUploadSize         *shared.SizeCapacity    `yaml:"maxUploadSize" json:"maxUploadSize"`
}

func (this *PlanConfig) Init() error {
	if this.TrafficLimit != nil {
		err := this.TrafficLimit.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
