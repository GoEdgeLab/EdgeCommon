// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "testing"

func TestMetricItemConfig_CurrentTime_Month(t *testing.T) {
	for _, period := range []int{1, 2, 3, 4, 5, 100} {
		var item = &MetricItemConfig{
			Period:     period,
			PeriodUnit: MetricItemPeriodUnitMonth,
		}
		_ = item.Init()
		t.Logf(item.CurrentTime())
	}
}

func TestMetricItemConfig_CurrentTime_Week(t *testing.T) {
	for _, period := range []int{1, 2, 3, 4, 5} {
		var item = &MetricItemConfig{
			Period:     period,
			PeriodUnit: MetricItemPeriodUnitWeek,
		}
		_ = item.Init()
		t.Log(period, ":", item.CurrentTime())
	}
}

func TestMetricItemConfig_CurrentTime_Day(t *testing.T) {
	for _, period := range []int{1, 2, 3, 4, 5, 13} {
		var item = &MetricItemConfig{
			Period:     period,
			PeriodUnit: MetricItemPeriodUnitDay,
		}
		_ = item.Init()
		t.Log(period, ":", item.CurrentTime())
	}
}

func TestMetricItemConfig_CurrentTime_Hour(t *testing.T) {
	for _, period := range []int{1, 2, 3, 4, 5, 13} {
		var item = &MetricItemConfig{
			Period:     period,
			PeriodUnit: MetricItemPeriodUnitHour,
		}
		_ = item.Init()
		t.Log(period, ":", item.CurrentTime())
	}
}

func TestMetricItemConfig_CurrentTime_Minute(t *testing.T) {
	for _, period := range []int{1, 2, 3, 4, 5, 13} {
		var item = &MetricItemConfig{
			Period:     period,
			PeriodUnit: MetricItemPeriodUnitMinute,
		}
		_ = item.Init()
		t.Log(period, ":", item.CurrentTime())
	}
}
