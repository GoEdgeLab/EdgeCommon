// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type PlanPriceType = string

const (
	PlanPriceTypeBandwidth PlanPriceType = "bandwidth"
	PlanPriceTypePeriod    PlanPriceType = "period"
)

func FindPlanPriceTypeName(priceType PlanPriceType) string {
	switch priceType {
	case PlanPriceTypeBandwidth:
		return "带宽用量"
	case PlanPriceTypePeriod:
		return "时间周期"
	}
	return ""
}

type PlanBandwidthPrice struct {
	Base float32 `yaml:"base" json:"base"` // 基础价格，单位是 元/GB
}
