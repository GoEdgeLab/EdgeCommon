// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

// PlanPriceType 套餐类型
type PlanPriceType = string

const (
	PlanPriceTypeTraffic   PlanPriceType = "traffic"   // 流量
	PlanPriceTypePeriod    PlanPriceType = "period"    // 周期
	PlanPriceTypeBandwidth PlanPriceType = "bandwidth" // 百分位
)

func FindPlanPriceTypeName(priceType PlanPriceType) string {
	switch priceType {
	case PlanPriceTypeTraffic:
		return "流量"
	case PlanPriceTypePeriod:
		return "时间周期"
	case PlanPriceTypeBandwidth:
		return "带宽"
	}
	return ""
}

// PlanTrafficPriceConfig 按流量计费价格配置
type PlanTrafficPriceConfig struct {
	Base float32 `yaml:"base" json:"base"` // 基础价格，单位是 元/GB
}

// PlanBandwidthPriceConfig 按百分位带宽计费配置
type PlanBandwidthPriceConfig struct {
	Percentile int                              `yaml:"percentile" json:"percentile"` // 百分位
	Ranges     []*PlanBandwidthPriceRangeConfig `yaml:"ranges" json:"ranges"`
}

func (this *PlanBandwidthPriceConfig) LookupRange(sizeMB float32) *PlanBandwidthPriceRangeConfig {
	if len(this.Ranges) == 0 {
		return nil
	}
	for _, r := range this.Ranges {
		if sizeMB >= r.MinMB && (r.MaxMB <= 0 || r.MaxMB >= sizeMB) {
			return r
		}
	}

	// 寻找最接近的
	for index, r := range this.Ranges {
		if r.MinMB >= sizeMB {
			if index > 0 {
				return this.Ranges[index-1]
			}
			return r
		}
	}

	for _, r := range this.Ranges {
		if r.MaxMB <= 0 || r.MaxMB >= sizeMB {
			return r
		}
	}

	// 获取最大值
	return this.Ranges[len(this.Ranges)-1]
}

func (this *PlanBandwidthPriceConfig) LookupPrice(sizeMB float32) float32 {
	var r = this.LookupRange(sizeMB)
	if r == nil {
		return 0
	}

	if r.TotalPrice > 0 {
		return r.TotalPrice
	}
	return r.PricePerMB * sizeMB
}

type PlanBandwidthPriceRangeConfig struct {
	MinMB      float32 `yaml:"minMB" json:"minMB"`
	MaxMB      float32 `yaml:"maxMB" json:"maxMB"`
	PricePerMB float32 `yaml:"pricePerMB" json:"pricePerMB"` // 单位价格，元/MB
	TotalPrice float32 `yaml:"totalPrice" json:"totalPrice"` // 总价格
}
