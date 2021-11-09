// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type PlanPriceType = string

const (
	PlanPriceTypeTraffic PlanPriceType = "traffic"
	PlanPriceTypePeriod  PlanPriceType = "period"
)

func FindPlanPriceTypeName(priceType PlanPriceType) string {
	switch priceType {
	case PlanPriceTypeTraffic:
		return "带宽用量"
	case PlanPriceTypePeriod:
		return "时间周期"
	}
	return ""
}

type PlanTrafficPrice struct {
	Base float32 `yaml:"base" json:"base"` // 基础价格，单位是 元/GB
}
