// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"

// UserFinanceConfig 财务相关设置
type UserFinanceConfig struct {
	IsOn                 bool                                    `yaml:"isOn" json:"isOn"`
	PriceType            serverconfigs.PlanPriceType             `yaml:"priceType" json:"priceType"`
	TrafficPriceConfig   *serverconfigs.PlanTrafficPriceConfig   `yaml:"trafficPrice" json:"trafficPrice"`
	BandwidthPriceConfig *serverconfigs.PlanBandwidthPriceConfig `yaml:"bandwidthPrice" json:"bandwidthPrice"`
}

func DefaultUserFinanceConfig() *UserFinanceConfig {
	return &UserFinanceConfig{
		PriceType: serverconfigs.PlanPriceTypeBandwidth,
	}
}
