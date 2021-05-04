// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

import "encoding/json"

// NodeValueItem 监控项
type NodeValueItem = string

const (
	NodeValueItemCPU         NodeValueItem = "cpu"         // CPU
	NodeValueItemMemory      NodeValueItem = "memory"      // 内存
	NodeValueItemLoad        NodeValueItem = "load"        // 负载
	NodeValueItemTrafficIn   NodeValueItem = "trafficIn"   // 上行流量
	NodeValueItemTrafficOut  NodeValueItem = "trafficOut"  // 下行流量
	NodeValueItemConnections NodeValueItem = "connections" // 连接数
	NodeValueItemDisk        NodeValueItem = "disk"        // 磁盘
)

type nodeValueItemDefinition struct {
	Code        string                          `json:"code"`
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	Params      []*nodeValueItemParamDefinition `json:"params"`
}

type nodeValueItemParamDefinition struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var nodeValueItemDefinitions = []*nodeValueItemDefinition{
	{
		Code: NodeValueItemCPU,
		Name: "CPU",
		Params: []*nodeValueItemParamDefinition{
			{
				Code:        "usage",
				Name:        "使用比例",
				Description: "一个不超过1的小数",
			},
		},
	},
	{
		Code: NodeValueItemMemory,
		Name: "内存",
		Params: []*nodeValueItemParamDefinition{
			{
				Code:        "usage",
				Name:        "使用比例",
				Description: "一个不超过1的小数",
			},
		},
	},
	{
		Code: NodeValueItemLoad,
		Name: "负载",
		Params: []*nodeValueItemParamDefinition{
			{
				Code:        "load1m",
				Name:        "1分钟负载",
				Description: "1分钟内的负载",
			},
			{
				Code:        "load5m",
				Name:        "5分钟负载",
				Description: "5分钟内的负载",
			},
			{
				Code:        "load15m",
				Name:        "15分钟负载",
				Description: "15分钟内的负载",
			},
		},
	},
	{
		Code:        NodeValueItemTrafficIn,
		Name:        "上行流量",
		Description: "客户端发送到服务器端的流量。",
		Params: []*nodeValueItemParamDefinition{
			{
				Code:        "total",
				Name:        "流量",
				Description: "单位为字节",
			},
		},
	},
	{
		Code:        NodeValueItemTrafficOut,
		Name:        "下行流量",
		Description: "服务器端发送到客户端的流量。",
		Params: []*nodeValueItemParamDefinition{
			{
				Code:        "total",
				Name:        "流量",
				Description: "单位为字节",
			},
		},
	},
	{
		Code: NodeValueItemConnections,
		Name: "连接数",
		Params: []*nodeValueItemParamDefinition{
			{
				Code:        "total",
				Name:        "总数",
				Description: "连接总数",
			},
		},
	},
	{
		Code: NodeValueItemDisk,
		Name: "硬盘空间",
		Params: []*nodeValueItemParamDefinition{
			{
				Code:        "usage",
				Name:        "使用比例",
				Description: "一个不超过1的小数",
			},
		},
	},
}

// FindAllNodeValueItemDefinitions 获取所有监控项信息
func FindAllNodeValueItemDefinitions() []*nodeValueItemDefinition {
	return nodeValueItemDefinitions
}

// FindNodeValueItemName 获取监控项名称
func FindNodeValueItemName(code NodeValueItem) string {
	for _, def := range nodeValueItemDefinitions {
		if def.Code == code {
			return def.Name
		}
	}
	return ""
}

// FindNodeValueItemParamName 获取监控项某个参数的名称
func FindNodeValueItemParamName(nodeCode NodeValueItem, paramCode string) string {
	for _, def := range nodeValueItemDefinitions {
		if def.Code == nodeCode {
			for _, p := range def.Params {
				if p.Code == paramCode {
					return p.Name
				}
			}
			return ""
		}
	}
	return ""
}

// NodeValueRange 值范围
type NodeValueRange = string

const (
	NodeValueRangeMinute NodeValueRange = "minute"
)

// NodeValueOperator 操作符
type NodeValueOperator = string

const (
	NodeValueOperatorGt  NodeValueOperator = "gt"  // 大于
	NodeValueOperatorGte NodeValueOperator = "gte" // 大于等于
	NodeValueOperatorLt  NodeValueOperator = "lt"  // 小于
	NodeValueOperatorLte NodeValueOperator = "lte" // 小于等于
	NodeValueOperatorEq  NodeValueOperator = "eq"  // 等于
	NodeValueOperatorNeq NodeValueOperator = "neq" // 不等于
)

type nodeValueOperatorDefinition struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var nodeValueOperatorDefinitions = []*nodeValueOperatorDefinition{
	{
		Code: NodeValueOperatorGt,
		Name: "大于",
	},
	{
		Code: NodeValueOperatorGte,
		Name: "大于等于",
	},
	{
		Code: NodeValueOperatorLt,
		Name: "小于",
	},
	{
		Code: NodeValueOperatorLte,
		Name: "小于等于",
	},
	{
		Code: NodeValueOperatorEq,
		Name: "等于",
	},
	{
		Code: NodeValueOperatorNeq,
		Name: "不等于",
	},
}

// FindNodeValueOperatorName 操作符名称
func FindNodeValueOperatorName(operator NodeValueOperator) string {
	for _, def := range nodeValueOperatorDefinitions {
		if def.Code == operator {
			return def.Name
		}
	}
	return ""
}

// FindAllNodeValueOperatorDefinitions 获取所有操作符定义
func FindAllNodeValueOperatorDefinitions() []*nodeValueOperatorDefinition {
	return nodeValueOperatorDefinitions
}

// NodeValueSumMethod 聚合方法
type NodeValueSumMethod = string

const (
	NodeValueSumMethodSum NodeValueSumMethod = "sum" // 相加
	NodeValueSumMethodAvg NodeValueSumMethod = "avg" // 平均
)

// FindNodeValueSumMethodName 聚合方法名称
func FindNodeValueSumMethodName(method NodeValueSumMethod) string {
	switch method {
	case NodeValueSumMethodSum:
		return "和"
	case NodeValueSumMethodAvg:
		return "平均"
	}
	return ""
}

// NodeValueDurationUnit 时间单位
type NodeValueDurationUnit = string

const (
	NodeValueDurationUnitMinute NodeValueDurationUnit = "minute"
)

// FindNodeValueDurationUnitName 时间单位名称
func FindNodeValueDurationUnitName(unit NodeValueDurationUnit) string {
	switch unit {
	case NodeValueDurationUnitMinute:
		return "分钟"
	}
	return ""
}

// UnmarshalNodeValue 对值进行解码
func UnmarshalNodeValue(valueJSON []byte) string {
	var result = ""
	err := json.Unmarshal(valueJSON, &result)
	if err != nil {
		// 暂时不提示错误
	}
	return result
}
