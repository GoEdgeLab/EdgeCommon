// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build plus

package nodeconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
)

type NodeActionParam = string

const (
	NodeActionParamDailyTrafficOut   NodeActionParam = "dailyTrafficOut"   // 节点日出口流量
	NodeActionParamMonthlyTrafficOut NodeActionParam = "monthlyTrafficOut" // 节点月出口流量

	// NodeActionParamTotalTraffic   NodeActionParam = "totalTraffic"   // 节点总流量 TODO 需要实现

	NodeActionParamBandwidthIn  NodeActionParam = "bandwidthIn"  // 节点入口带宽
	NodeActionParamBandwidthOut NodeActionParam = "bandwidthOut" // 节点出口带宽
	NodeActionParamCPUUsage     NodeActionParam = "cpuUsage"     // 节点CPU用量，百分比制，0-100
	NodeActionParamMemoryUsage  NodeActionParam = "memoryUsage"  // 节点内存用量，百分比制，0-100
	NodeActionParamLoad         NodeActionParam = "load"         // 当前节点负载

	// NodeActionParamConnectivity NodeActionParam = "connectivity" // 节点连通性 TODO 需要实现

	NodeActionParamHealthCheckFailure NodeActionParam = "heathCheckFailure" // 节点健康检查失败
)

type NodeActionParamDefinition struct {
	Name        string               `json:"name"`
	Code        string               `json:"code"`
	Description string               `json:"description"`
	Operators   []NodeActionOperator `json:"operators"`
	ValueName   string               `json:"valueName"`
	ValueType   string               `json:"valueType"`
}

func FindAllNodeActionParamDefinitions() []*NodeActionParamDefinition {
	return []*NodeActionParamDefinition{
		{
			Code:        NodeActionParamBandwidthOut,
			Name:        "节点出口带宽",
			Description: "当前节点当前时间点的出口平均带宽（从节点发送到客户端）。",
			Operators:   allNodeActionNumberOperators,
			ValueName:   "对比带宽",
			ValueType:   "bandwidth",
		},
		{
			Code:        NodeActionParamBandwidthIn,
			Name:        "节点入口带宽",
			Description: "当前节点当前时间点的入口平均带宽（从客户端发送到节点）。",
			Operators:   allNodeActionNumberOperators,
			ValueName:   "对比带宽",
			ValueType:   "bandwidth",
		},
		{
			Code:        NodeActionParamMonthlyTrafficOut,
			Name:        "节点当月流量",
			Description: "当前节点在当月的出口流量。",
			Operators:   allNodeActionNumberOperators,
			ValueName:   "对比流量",
			ValueType:   "traffic",
		},
		{
			Code:        NodeActionParamDailyTrafficOut,
			Name:        "节点当日流量",
			Description: "当前节点在当天的出口流量。",
			Operators:   allNodeActionNumberOperators,
			ValueName:   "对比流量",
			ValueType:   "traffic",
		},
		{
			Code:        NodeActionParamCPUUsage,
			Name:        "节点CPU利用率",
			Description: "节点当前CPU利用率，取值范围为0-100。",
			Operators:   allNodeActionNumberOperators,
			ValueName:   "CPU利用率",
			ValueType:   "cpu",
		},
		{
			Code:        NodeActionParamMemoryUsage,
			Name:        "节点内存利用率",
			Description: "节点当前内存利用率，取值范围为0-100。",
			Operators:   allNodeActionNumberOperators,
			ValueName:   "内存利用率",
			ValueType:   "memory",
		},
		{
			Code:        NodeActionParamLoad,
			Name:        "节点负载",
			Description: "节点当前负载，取值范围为0-∞，通常超过10表示系统负载较重。",
			Operators:   allNodeActionNumberOperators,
			ValueName:   "系统负载",
			ValueType:   "load",
		},
		{
			Code:        NodeActionParamHealthCheckFailure,
			Name:        "健康检查失败",
			Description: "当前节点任一IP健康检查失败。",
			Operators:   nil,
		},
	}
}

type NodeActionOperator = string

const (
	NodeActionOperatorGt  NodeActionOperator = "gt"
	NodeActionOperatorGte NodeActionOperator = "gte"
	NodeActionOperatorLt  NodeActionOperator = "lt"
	NodeActionOperatorLte NodeActionOperator = "lte"
	NodeActionOperatorEq  NodeActionOperator = "eq"
)

var allNodeActionNumberOperators = []NodeActionOperator{NodeActionOperatorGt, NodeActionOperatorGte, NodeActionOperatorLt, NodeActionOperatorLte, NodeActionOperatorEq}

func FindAllNodeActionOperatorDefinitions() []*shared.Definition {
	return []*shared.Definition{
		{
			Code: NodeActionOperatorGt,
			Name: "大于(＞)",
		},
		{
			Code: NodeActionOperatorGte,
			Name: "大于等于(≥)",
		},
		{
			Code: NodeActionOperatorLt,
			Name: "小于(＜)",
		},
		{
			Code: NodeActionOperatorLte,
			Name: "小于等于(≤)",
		},
		{
			Code: NodeActionOperatorEq,
			Name: "等于(=)",
		},
	}
}

// NodeActionCondConnector 条件之间关系
type NodeActionCondConnector = string

const (
	NodeActionCondConnectorAnd NodeActionCondConnector = "and"
	NodeActionCondConnectorOr  NodeActionCondConnector = "or"
)

type NodeActionCond struct {
	Param    NodeActionParam    `json:"param"`    // 参数名
	Operator NodeActionOperator `json:"operator"` // 操作符
	Value    any                `json:"value"`    // 对比值
}

func (this *NodeActionCond) Match(value any) bool {
	var paramDef *NodeActionParamDefinition
	for _, paramDef2 := range FindAllNodeActionParamDefinitions() {
		if paramDef2.Code == this.Param {
			paramDef = paramDef2
			break
		}
	}
	if paramDef == nil {
		return false
	}

	switch paramDef.ValueType {
	case "bandwidth":
		if value != nil && !this.isScalar(value) {
			var value1Map = maps.NewMap(value)
			value = this.toBandwidthBits(value1Map)
		}
		var value2Map = maps.NewMap(this.Value)
		return this.compare(value, this.toBandwidthBits(value2Map))
	case "traffic":
		if value != nil && !this.isScalar(value) {
			var value1Map = maps.NewMap(value)
			value = this.toTrafficBytes(value1Map)
		}
		var value2Map = maps.NewMap(this.Value)
		return this.compare(value, this.toTrafficBytes(value2Map))
	case "cpu":
		return this.compare(value, this.Value)
	case "memory":
		return this.compare(value, this.Value)
	case "load":
		return this.compare(value, this.Value)
	case "":
		return true
	}

	return false
}

func (this *NodeActionCond) compare(value1, value2 any) bool {
	switch this.Operator {
	case NodeActionOperatorGt:
		return types.Int64(value1) > types.Int64(value2)
	case NodeActionOperatorGte:
		return types.Int64(value1) >= types.Int64(value2)
	case NodeActionOperatorLt:
		return types.Int64(value1) < types.Int64(value2)
	case NodeActionOperatorLte:
		return types.Int64(value1) <= types.Int64(value2)
	case NodeActionOperatorEq:
		return types.Int64(value1) == types.Int64(value2)
	}
	return false
}

func (this *NodeActionCond) toBandwidthBits(m maps.Map) int64 {
	var count = m.GetInt64("count")
	var unit = m.GetString("unit")
	if count <= 0 {
		return 0
	}
	switch unit {
	case "b":
		return count
	case "kb":
		return count << 10
	case "mb":
		return count << 20
	case "gb":
		return count << 30
	case "tb":
		return count << 40
	}
	return 0
}

func (this *NodeActionCond) toTrafficBytes(m maps.Map) int64 {
	var count = m.GetInt64("count")
	var unit = m.GetString("unit")
	if count <= 0 {
		return 0
	}
	switch unit {
	case "b":
		return count
	case "kb":
		return count << 10
	case "mb":
		return count << 20
	case "gb":
		return count << 30
	case "tb":
		return count << 40
	case "pb":
		return count << 50
	case "eb":
		return count << 60
	}
	return 0
}

func (this *NodeActionCond) isScalar(value any) bool {
	if value == nil {
		return false
	}
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, float32, float64, string, bool:
		return true
	}
	return false
}

type NodeActionCondsConfig struct {
	Conds     []*NodeActionCond       `json:"conds"`
	Connector NodeActionCondConnector `json:"connector"`
}

func NewNodeActionCondsConfig() *NodeActionCondsConfig {
	return &NodeActionCondsConfig{
		Connector: NodeActionCondConnectorAnd,
	}
}

func (this *NodeActionCondsConfig) Match(valueGetter func(param NodeActionParam) (value any, err error)) (bool, error) {
	if len(this.Conds) == 0 {
		return false, nil
	}

	for index, cond := range this.Conds {
		value, err := valueGetter(cond.Param)
		if err != nil {
			return false, err
		}

		var b = cond.Match(value)
		if !b {
			if this.Connector == NodeActionCondConnectorAnd {
				return false, nil
			}

			// 如果是最后一个OR条件，则直接返回false
			if index == len(this.Conds)-1 {
				return false, nil
			}
		} else {
			if this.Connector == NodeActionCondConnectorOr {
				return true, nil
			}

			// 如果是最后一个AND条件，则直接返回true
			if index == len(this.Conds)-1 {
				return true, nil
			}
		}
	}

	return true, nil
}

// NodeActionCode 动作代号
type NodeActionCode = string

const (
	NodeActionCodeUp                           NodeActionCode = "up"                           // 上线
	NodeActionCodeDown                         NodeActionCode = "down"                         // 下线
	NodeActionCodeSwitchToBackupNodesInCluster NodeActionCode = "switchToBackupNodesInCluster" // 切换到集群备用节点
	NodeActionCodeSwitchToBackupNodesInGroup   NodeActionCode = "switchToBackupNodesInGroup"   // 切换到分组备用节点
	NodeActionCodeSwitchToBackupIP             NodeActionCode = "switchToBackupIP"             // 切换到备用IP
	NodeActionCodeEnableBackupNodesInCluster   NodeActionCode = "enableBackupNodesInCluster"   // 启用集群备用节点
	NodeActionCodeEnableBackupNodesInGroup     NodeActionCode = "enableBackupNodesInGroup"     // 启用分组备用节点
	NodeActionCodeEnableBackupIP               NodeActionCode = "enableBackupIP"               // 启用备用IP
	NodeActionCodeWebHook                      NodeActionCode = "webHook"                      // WebHook
)

func FindAllNodeActionDefinitions() []*shared.Definition {
	return []*shared.Definition{
		{
			Code:        NodeActionCodeUp,
			Name:        "上线当前节点",
			Description: "将当前节点状态设置为在线。",
		},
		{
			Code:        NodeActionCodeDown,
			Name:        "下线当前节点",
			Description: "将当前节点状态设置为离线。",
		},
		{
			Code:        NodeActionCodeSwitchToBackupNodesInCluster,
			Name:        "切换到集群备用节点",
			Description: "下线当前节点并启用节点所在集群备用节点。",
		},
		{
			Code:        NodeActionCodeSwitchToBackupNodesInGroup,
			Name:        "切换到分组备用节点",
			Description: "下线当前节点并启用节点所在分组备用节点。",
		},
		{
			Code:        NodeActionCodeSwitchToBackupIP,
			Name:        "切换到备用IP",
			Description: "将当前节点的IP切换到当前节点配置的备用IP",
		},

		{
			Code:        NodeActionCodeEnableBackupNodesInCluster,
			Name:        "启用集群备用节点",
			Description: "保持当前节点并启用节点所在集群备用节点。",
		},
		{
			Code:        NodeActionCodeEnableBackupNodesInGroup,
			Name:        "启用分组备用节点",
			Description: "保持当前节点并启用节点所在分组备用节点。",
		},
		{
			Code:        NodeActionCodeEnableBackupIP,
			Name:        "启用备用IP",
			Description: "保持当前节点的IP并启用当前节点配置的备用IP",
		},

		{
			Code:        NodeActionCodeWebHook,
			Name:        "WebHook",
			Description: "通过WebHook发送通知到URL",
		},
	}
}

func FindNodeActionDefinition(code NodeActionCode) *shared.Definition {
	for _, def := range FindAllNodeActionDefinitions() {
		if def.Code == code {
			return def
		}
	}
	return nil
}

func FindNodeActionName(code NodeActionCode) string {
	var def = FindNodeActionDefinition(code)
	if def != nil {
		return def.Name
	}
	return ""
}

// NodeActionConfig 动作配置
type NodeActionConfig struct {
	Code   NodeActionCode `json:"code"`   // 动作代号
	Params any            `json:"params"` // 动作参数
}

func NewNodeActionConfig() *NodeActionConfig {
	return &NodeActionConfig{}
}

type NodeActionCodeWebHookParams struct {
	URL string `json:"url"` // URL路径
}

// NodeActionStatus 动作状态
type NodeActionStatus struct {
	ActionId  int64                  `json:"actionId"`  // 动作ID
	CreatedAt int64                  `json:"createdAt"` // 状态创建时间
	Conds     *NodeActionCondsConfig `json:"conds"`     // 动作条件
	Action    *NodeActionConfig      `json:"action"`    // 动作配置
	ExpiresAt int64                  `json:"expiresAt"` // 过期时间
}
