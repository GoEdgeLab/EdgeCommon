// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import timeutil "github.com/iwind/TeaGo/utils/time"

type TrafficLimitTarget = string

const (
	TrafficLimitTargetTraffic TrafficLimitTarget = "traffic"
	TrafficLimitTargetRequest TrafficLimitTarget = "request"
)

// TrafficLimitStatus 流量限制状态
type TrafficLimitStatus struct {
	UntilDay   string `yaml:"untilDay" json:"untilDay"`     // 有效日期，格式YYYYMMDD
	PlanId     int64  `yaml:"planId" json:"planId"`         // 套餐ID
	DateType   string `yaml:"dateType" json:"dateType"`     // 日期类型 day|month
	TargetType string `yaml:"targetType" json:"targetType"` // 限制类型：traffic|request|...
}

func (this *TrafficLimitStatus) IsValid() bool {
	if len(this.UntilDay) == 0 {
		return false
	}
	return this.UntilDay >= timeutil.Format("Ymd")
}
