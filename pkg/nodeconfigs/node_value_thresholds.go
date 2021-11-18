// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

import "github.com/iwind/TeaGo/maps"

type IPAddressThresholdItem = string

const (
	IPAddressThresholdItemNodeAvgRequests       IPAddressThresholdItem = "nodeAvgRequests"       // 个
	IPAddressThresholdItemNodeAvgTrafficOut     IPAddressThresholdItem = "nodeAvgTrafficOut"     // 节点下行流量 M
	IPAddressThresholdItemNodeAvgTrafficIn      IPAddressThresholdItem = "nodeAvgTrafficIn"      // 节点上行流量 M
	IPAddressThresholdItemNodeHealthCheckFailed IPAddressThresholdItem = "nodeHealthCheckFailed" // 节点健康检查失败
	IPAddressThresholdItemGroupAvgRequests      IPAddressThresholdItem = "groupAvgRequests"      // 个
	IPAddressThresholdItemGroupAvgTrafficIn     IPAddressThresholdItem = "groupAvgTrafficIn"     // 分组上行流量 M
	IPAddressThresholdItemGroupAvgTrafficOut    IPAddressThresholdItem = "groupAvgTrafficOut"    // 分组下行流量 M
	IPAddressThresholdItemClusterAvgRequests    IPAddressThresholdItem = "clusterAvgRequests"    // 个
	IPAddressThresholdItemClusterAvgTrafficIn   IPAddressThresholdItem = "clusterAvgTrafficIn"   // 集群上行流量 M
	IPAddressThresholdItemClusterAvgTrafficOut  IPAddressThresholdItem = "clusterAvgTrafficOut"  // 集群下行流量 M
	IPAddressThresholdItemConnectivity          IPAddressThresholdItem = "connectivity"          // 0-100
)

// FindAllIPAddressThresholdItems IP相关阈值项目
func FindAllIPAddressThresholdItems() []maps.Map {
	return []maps.Map{
		{
			"name":        "节点平均请求数",
			"code":        IPAddressThresholdItemNodeAvgRequests,
			"description": "当前节点在单位时间内接收到的平均请求数。",
			"unit":        "个",
		},
		{
			"name":        "节点平均下行流量",
			"code":        IPAddressThresholdItemNodeAvgTrafficOut,
			"description": "当前节点在单位时间内发送的下行流量。",
			"unit":        "M",
		},
		{
			"name":        "节点平均上行流量",
			"code":        IPAddressThresholdItemNodeAvgTrafficIn,
			"description": "当前节点在单位时间内接收的上行流量。",
			"unit":        "M",
		},
		{
			"name":        "节点健康检查失败",
			"code":        IPAddressThresholdItemNodeHealthCheckFailed,
			"description": "当前节点健康检查失败",
			"unit":        "",
		},

		{
			"name":        "IP连通性",
			"code":        IPAddressThresholdItemConnectivity,
			"description": "通过区域监控得到的当前IP地址的连通性数值，取值在0和100之间。",
			"unit":        "%",
		},

		{
			"name":        "分组平均请求数",
			"code":        IPAddressThresholdItemGroupAvgRequests,
			"description": "当前节点所在分组在单位时间内接收到的平均请求数。",
			"unit":        "个",
		},
		{
			"name":        "分组平均下行流量",
			"code":        IPAddressThresholdItemGroupAvgTrafficOut,
			"description": "当前节点所在分组在单位时间内发送的下行流量。",
			"unit":        "M",
		},
		{
			"name":        "分组平均上行流量",
			"code":        IPAddressThresholdItemGroupAvgTrafficIn,
			"description": "当前节点所在分组在单位时间内接收的上行流量。",
			"unit":        "M",
		},

		{
			"name":        "集群平均请求数",
			"code":        IPAddressThresholdItemClusterAvgRequests,
			"description": "当前节点所在集群在单位时间内接收到的平均请求数。",
			"unit":        "个",
		},
		{
			"name":        "集群平均下行流量",
			"code":        IPAddressThresholdItemClusterAvgTrafficOut,
			"description": "当前节点所在集群在单位时间内发送的下行流量。",
			"unit":        "M",
		},
		{
			"name":        "集群平均上行流量",
			"code":        IPAddressThresholdItemClusterAvgTrafficIn,
			"description": "当前节点所在集群在单位时间内接收的上行流量。",
			"unit":        "M",
		},
	}
}

// IPAddressThresholdConfig 阈值列表
type IPAddressThresholdConfig struct {
	Id      int64                             `json:"id"`
	Items   []*IPAddressThresholdItemConfig   `json:"items"`
	Actions []*IPAddressThresholdActionConfig `json:"actions"`
}

// IPAddressThresholdItemConfig 阈值项目
type IPAddressThresholdItemConfig struct {
	Item         IPAddressThresholdItem `json:"item"`
	Operator     NodeValueOperator      `json:"operator"`
	Value        float64                `json:"value"`
	Duration     int                    `json:"duration"`
	DurationUnit NodeValueDurationUnit  `json:"durationUnit"`
	Options      maps.Map               `json:"options"` // 附加选项
}

type IPAddressThresholdActionConfig struct {
	Action  string   `json:"action"`
	Options maps.Map `json:"options"`
}

// IPAddressThresholdAction 动作
type IPAddressThresholdAction = string

const (
	IPAddressThresholdActionUp      IPAddressThresholdAction = "up"      // 上线
	IPAddressThresholdActionDown    IPAddressThresholdAction = "down"    // 下线
	IPAddressThresholdActionNotify  IPAddressThresholdAction = "notify"  // 通知
	IPAddressThresholdActionSwitch  IPAddressThresholdAction = "switch"  // 切换到备用IP
	IPAddressThresholdActionWebHook IPAddressThresholdAction = "webHook" // 调用外部Webhook
)

// FindAllIPAddressThresholdActions IP相关阈值动作
func FindAllIPAddressThresholdActions() []maps.Map {
	return []maps.Map{
		{
			"name":        "上线",
			"code":        IPAddressThresholdActionUp,
			"description": "上线当前IP。",
		},
		{
			"name":        "下线",
			"code":        IPAddressThresholdActionDown,
			"description": "下线当前IP。",
		},
		{
			"name":        "通知",
			"code":        IPAddressThresholdActionNotify,
			"description": "发送已达到阈值通知。",
		},
		{
			"name":        "切换",
			"code":        IPAddressThresholdActionSwitch,
			"description": "在DNS中记录中将IP切换到指定的备用IP。",
		},
		{
			"name":        "WebHook",
			"code":        IPAddressThresholdActionWebHook,
			"description": "调用外部的WebHook。",
		},
	}
}
