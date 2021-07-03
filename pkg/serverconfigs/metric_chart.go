// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type MetricChartType = string

const (
	MetricChartTypePie      MetricChartType = "pie"
	MetricChartTypeBar      MetricChartType = "bar"
	MetricChartTypeTimeBar  MetricChartType = "timeBar"
	MetricChartTypeTimeLine MetricChartType = "timeLine"
	MetricChartTypeTable    MetricChartType = "table"
)

func FindAllMetricChartTypes() []*shared.Definition {
	return []*shared.Definition{
		{
			Name:        "柱图",
			Code:        MetricChartTypeBar,
			Description: "通过柱图展示各个对象的排行。",
			Icon:        "chart bar",
		},
		{
			Name:        "饼图",
			Code:        MetricChartTypePie,
			Description: "通过饼图展示各个对象的占比。",
			Icon:        "chart pie",
		},
		{
			Name:        "时间柱图",
			Code:        MetricChartTypeTimeBar,
			Description: "通过柱图展示各个对象在不同时间段的变化。",
			Icon:        "chart bar",
		},
		{
			Name:        "时间线图",
			Code:        MetricChartTypeTimeLine,
			Description: "通过线图展示各个对象在不同时间段的变化。",
			Icon:        "chart line area",
		},
		{
			Name:        "表格",
			Code:        MetricChartTypeTable,
			Description: "以表格的形式展示数据。",
			Icon:        "table",
		},
	}
}

func FindAllMetricChartTypeName(chartType MetricChartType) string {
	for _, def := range FindAllMetricChartTypes() {
		if def.Code == chartType {
			return def.Name
		}
	}
	return ""
}
