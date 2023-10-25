// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// MetricItemCategory 指标分类
type MetricItemCategory = string

const (
	MetricItemCategoryHTTP MetricItemCategory = "http"
	MetricItemCategoryTCP  MetricItemCategory = "tcp"
	MetricItemCategoryUDP  MetricItemCategory = "udp"
)

func FindAllMetricItemCategoryCodes() []MetricItemCategory {
	return []MetricItemCategory{MetricItemCategoryHTTP, MetricItemCategoryTCP, MetricItemCategoryUDP}
}

// MetricItemPeriodUnit 指标周期单位
type MetricItemPeriodUnit = string

const (
	MetricItemPeriodUnitMinute MetricItemPeriodUnit = "minute"
	MetricItemPeriodUnitHour   MetricItemPeriodUnit = "hour"
	MetricItemPeriodUnitDay    MetricItemPeriodUnit = "day"
	MetricItemPeriodUnitWeek   MetricItemPeriodUnit = "week"
	MetricItemPeriodUnitMonth  MetricItemPeriodUnit = "month"
)

// MetricItemValueType 值类型
type MetricItemValueType = string

const (
	MetricItemValueTypeCount MetricItemValueType = "count" // 数量
	MetricItemValueTypeByte  MetricItemValueType = "byte"  //  字节
)

// FindAllMetricKeyDefinitions 指标对象定义
func FindAllMetricKeyDefinitions(category MetricItemCategory) []*shared.Definition {
	switch category {
	case MetricItemCategoryHTTP:
		return []*shared.Definition{
			{
				Name:        "客户端地址（IP）",
				Code:        "${remoteAddr}",
				Description: "会依次根据X-Forwarded-For、X-Real-IP、RemoteAddr获取，适用于前端可能有别的反向代理的情形，存在被伪造的可能",
			},
			{
				Name:        "直接客户端地址（IP）",
				Code:        "${rawRemoteAddr}",
				Description: "返回直接连接服务的客户端原始IP地址",
			},
			{
				Name:        "客户端用户名",
				Code:        "${remoteUser}",
				Description: "通过基本认证填入的用户名",
			},
			{
				Name:        "请求URI",
				Code:        "${requestURI}",
				Description: "包含参数，比如/hello?name=lily",
			},
			{
				Name:        "请求路径",
				Code:        "${requestPath}",
				Description: "不包含参数，比如/hello",
			},
			{
				Name:        "完整URL",
				Code:        "${requestURL}",
				Description: "比如https://example.com/hello?name=lily",
			},
			{
				Name:        "请求方法",
				Code:        "${requestMethod}",
				Description: "比如GET、POST等",
			},
			{
				Name:        "请求协议Scheme",
				Code:        "${scheme}",
				Description: "http或https",
			},
			{
				Name:        "文件扩展名",
				Code:        "${requestPathExtension}",
				Description: "请求路径中的文件扩展名，包括点符号，比如.html、.png",
			},
			{
				Name:        "小写文件扩展名",
				Code:        "${requestPathLowerExtension}",
				Description: "请求路径中的文件扩展名小写形式，包括点符号，比如.html、.png",
			},
			{
				Name:        "主机名",
				Code:        "${host}",
				Description: "通常是请求的域名",
			},

			{
				Name:        "HTTP协议",
				Code:        "${proto}",
				Description: "包含版本的HTTP请求协议，类似于HTTP/1.0",
			},
			{
				Name:        "URL参数值",
				Code:        "${arg.NAME}",
				Description: "单个URL参数值",
			},
			{
				Name:        "请求来源URL",
				Code:        "${referer}",
				Description: "请求来源Referer URL",
			},
			{
				Name:        "请求来源URL域名",
				Code:        "${referer.host}",
				Description: "请求来源Referer URL域名",
			},
			{
				Name:        "Header值",
				Code:        "${header.NAME}",
				Description: "单个Header值，比如${header.User-Agent}",
			},
			{
				Name:        "Cookie值",
				Code:        "${cookie.NAME}",
				Description: "单个cookie值，比如${cookie.sid}",
			},

			// =========  以下是响应  =========
			{
				Name: "状态码",
				Code: "${status}",
			},
			{
				Name: "响应的Content-Type值",
				Code: "${response.contentType}",
			},
		}
	case MetricItemCategoryTCP:
		// TODO
		return []*shared.Definition{}
	case MetricItemCategoryUDP:
		// TODO
		return []*shared.Definition{}
	}
	return []*shared.Definition{}
}

// MetricValueDefinition 指标数值定义
type MetricValueDefinition struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Type        string `json:"type"`
}

// FindAllMetricValueDefinitions 指标数值定义
func FindAllMetricValueDefinitions(category MetricItemCategory) []*MetricValueDefinition {
	switch category {
	case MetricItemCategoryHTTP:
		return []*MetricValueDefinition{
			{
				Name: "请求数",
				Code: "${countRequest}",
				Type: MetricItemValueTypeCount,
			},
			{
				Name: "连接数",
				Code: "${countConnection}",
				Type: MetricItemValueTypeCount,
			},
			{
				Name: "下行流量",
				Code: "${countTrafficOut}",
				Type: MetricItemValueTypeByte,
			},
			{
				Name: "上行流量",
				Code: "${countTrafficIn}",
				Type: MetricItemValueTypeByte,
			},
		}
	case MetricItemCategoryTCP:
		return []*MetricValueDefinition{
			{
				Name: "连接数",
				Code: "${countConnection}",
				Type: MetricItemValueTypeCount,
			},
			{
				Name: "下行流量",
				Code: "${countTrafficOut}",
				Type: MetricItemValueTypeByte,
			},
			{
				Name: "上行流量",
				Code: "${countTrafficIn}",
				Type: MetricItemValueTypeByte,
			},
		}
	case MetricItemCategoryUDP:
		return []*MetricValueDefinition{
			{
				Name: "连接数",
				Code: "${countConnection}",
				Type: MetricItemValueTypeCount,
			},
			{
				Name: "下行流量",
				Code: "${countTrafficOut}",
				Type: MetricItemValueTypeByte,
			},
			{
				Name: "上行流量",
				Code: "${countTrafficIn}",
				Type: MetricItemValueTypeByte,
			},
		}
	}
	return []*MetricValueDefinition{}
}

func FindMetricValueName(category MetricItemCategory, code string) string {
	for _, def := range FindAllMetricValueDefinitions(category) {
		if def.Code == code {
			return def.Name
		}
	}
	return code
}

func FindMetricValueType(category MetricItemCategory, code string) string {
	for _, def := range FindAllMetricValueDefinitions(category) {
		if def.Code == code {
			return def.Type
		}
	}
	return MetricItemValueTypeCount
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

func FindMetricPeriodUnitName(unit string) string {
	switch unit {
	case MetricItemPeriodUnitMonth:
		return "月"
	case MetricItemPeriodUnitWeek:
		return "周"
	case MetricItemPeriodUnitDay:
		return "天"
	case MetricItemPeriodUnitHour:
		return "小时"
	case MetricItemPeriodUnitMinute:
		return "分钟"
	}
	return ""
}
