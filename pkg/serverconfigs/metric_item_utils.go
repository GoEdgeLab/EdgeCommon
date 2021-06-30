// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

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

// HTTP相关指标对象
type metricKeyDefinition struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type metricValueDefinition struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func FindAllHTTPMetricKeyDefinitions() []*metricKeyDefinition {
	// TODO

	return nil
}

func FindAllMetricValueDefinitions(category MetricItemCategory) []*metricValueDefinition {
	switch category {
	case MetricItemCategoryHTTP:
		return []*metricValueDefinition{
			{
				Name: "请求数",
				Code: "${countRequest}",
			},
			{
				Name: "连接数",
				Code: "${countConnection}",
			},
			{
				Name: "下行流量",
				Code: "${countTrafficOut}",
			},
			{
				Name: "上行流量",
				Code: "${countTrafficIn}",
			},
		}
	case MetricItemCategoryTCP:
		return []*metricValueDefinition{
			{
				Name: "连接数",
				Code: "${countConnection}",
			},
			{
				Name: "下行流量",
				Code: "${countTrafficOut}",
			},
			{
				Name: "上行流量",
				Code: "${countTrafficIn}",
			},
		}
	case MetricItemCategoryUDP:
		return []*metricValueDefinition{
			{
				Name: "连接数",
				Code: "${countConnection}",
			},
			{
				Name: "下行流量",
				Code: "${countTrafficOut}",
			},
			{
				Name: "上行流量",
				Code: "${countTrafficIn}",
			},
		}
	}
	return []*metricValueDefinition{}
}

// FindAllTCPMetricKeyDefinitions TCP相关指标对象
func FindAllTCPMetricKeyDefinitions() []*metricKeyDefinition {
	// TODO

	return nil
}

// FindAllUDPMetricKeyDefinitions UDP相关指标对象
func FindAllUDPMetricKeyDefinitions() []*metricKeyDefinition {
	// TODO

	return nil
}

// HumanMetricTime 格式化时间，让时间更易读
func HumanMetricTime(periodUnit MetricItemPeriodUnit, time string) string {
	switch periodUnit {
	case MetricItemPeriodUnitMonth:
		if len(time) != 6 {
			return time
		}
		return time[:4] + "-" + time[4:]
	case MetricItemPeriodUnitWeek:
		if len(time) != 6 {
			return time
		}
		return time[:4] + "-" + time[4:]
	case MetricItemPeriodUnitDay:
		if len(time) != 8 {
			return time
		}
		return time[:4] + "-" + time[4:6] + "-" + time[6:]
	case MetricItemPeriodUnitHour:
		if len(time) != 10 {
			return time
		}
		return time[:4] + "-" + time[4:6] + "-" + time[6:8] + " " + time[8:]
	case MetricItemPeriodUnitMinute:
		if len(time) != 12 {
			return time
		}
		return time[:4] + "-" + time[4:6] + "-" + time[6:8] + " " + time[8:10] + ":" + time[10:]
	}
	return time
}
