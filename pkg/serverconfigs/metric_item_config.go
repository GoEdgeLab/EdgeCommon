// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/cespare/xxhash/v2"
	"github.com/iwind/TeaGo/types"
	"strconv"
)

// MetricItemCategory 指标分类
type MetricItemCategory = string

const (
	MetricItemCategoryHTTP MetricItemCategory = "http"
	MetricItemCategoryTCP  MetricItemCategory = "tcp"
	MetricItemCategoryUDP  MetricItemCategory = "udp"
)

// MetricItemPeriodUnit 指标周期单位
type MetricItemPeriodUnit = string

const (
	MetricItemPeriodUnitMinute MetricItemPeriodUnit = "minute"
	MetricItemPeriodUnitHour   MetricItemPeriodUnit = "hour"
	MetricItemPeriodUnitDay    MetricItemPeriodUnit = "day"
	MetricItemPeriodUnitWeek   MetricItemPeriodUnit = "week"
	MetricItemPeriodUnitMonth  MetricItemPeriodUnit = "month"
)

// MetricItemConfig 指标配置
type MetricItemConfig struct {
	Id         int64                `yaml:"id" json:"id"`
	IsOn       bool                 `yaml:"isOn" json:"isOn"`
	Category   MetricItemCategory   `yaml:"category" json:"category"`
	Period     int                  `yaml:"period" json:"period"`
	PeriodUnit MetricItemPeriodUnit `yaml:"periodUnit" json:"periodUnit"`
	Keys       []string             `yaml:"keys" json:"keys"`
	Value      string               `yaml:"value" json:"value"`

	sumType string // 统计类型
}

// Init 初始化
func (this *MetricItemConfig) Init() error {
	return nil
}

// ParseRequest 处理请求
func (this *MetricItemConfig) ParseRequest(format func(string) string) (key string, hash string, value float64) {
	for _, k := range this.Keys {
		key += "@" + format(k)
	}
	hash = strconv.FormatUint(xxhash.Sum64String(key), 10)

	// TODO value将来支持复杂运算，比如 ${request.traffic.bytes} * 8
	if len(this.Value) == 0 {
		value = 1
	} else {
		value = types.Float64(format(this.Value))
	}

	return
}
