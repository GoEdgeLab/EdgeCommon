// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

import "github.com/iwind/TeaGo/maps"

type IPAddressThresholdItem = string

const (
	IPAddressThresholdItemNodeAvgRequests      IPAddressThresholdItem = "nodeAvgRequests"      // 个
	IPAddressThresholdItemNodeAvgTrafficOut    IPAddressThresholdItem = "nodeAvgTrafficOut"    // 节点下行流量 M
	IPAddressThresholdItemNodeAvgTrafficIn     IPAddressThresholdItem = "nodeAvgTrafficIn"     // 节点上行流量 M
	IPAddressThresholdItemGroupAvgRequests     IPAddressThresholdItem = "groupAvgRequests"     // 个
	IPAddressThresholdItemGroupAvgTrafficIn    IPAddressThresholdItem = "groupAvgTrafficIn"    // 分组上行流量 M
	IPAddressThresholdItemGroupAvgTrafficOut   IPAddressThresholdItem = "groupAvgTrafficOut"   // 分组下行流量 M
	IPAddressThresholdItemClusterAvgRequests   IPAddressThresholdItem = "clusterAvgRequests"   // 个
	IPAddressThresholdItemClusterAvgTrafficIn  IPAddressThresholdItem = "clusterAvgTrafficIn"  // 集群上行流量 M
	IPAddressThresholdItemClusterAvgTrafficOut IPAddressThresholdItem = "clusterAvgTrafficOut" // 集群下行流量 M
	IPAddressThresholdItemConnectivity         IPAddressThresholdItem = "connectivity"         // 0-100
)

func FindAllIPAddressThresholdItems() []maps.Map {
	return []maps.Map{
		{
			"name":        "节点平均请求数",
			"code":        "nodeAvgRequests",
			"description": "当前节点在单位时间内接收到的平均请求数。",
			"unit":        "个",
		},
		{
			"name":        "节点平均下行流量",
			"code":        "nodeAvgTrafficOut",
			"description": "当前节点在单位时间内发送的下行流量。",
			"unit":        "M",
		},
		{
			"name":        "节点平均上行流量",
			"code":        "nodeAvgTrafficIn",
			"description": "当前节点在单位时间内接收的上行流量。",
			"unit":        "M",
		},

		{
			"name":        "IP连通性",
			"code":        "connectivity",
			"description": "通过区域监控得到的当前IP地址的连通性数值，取值在0和100之间。",
			"unit":        "%",
		},

		{
			"name":        "分组平均请求数",
			"code":        "groupAvgRequests",
			"description": "当前节点所在分组在单位时间内接收到的平均请求数。",
			"unit":        "个",
		},
		{
			"name":        "分组平均下行流量",
			"code":        "groupAvgTrafficOut",
			"description": "当前节点所在分组在单位时间内发送的下行流量。",
			"unit":        "M",
		},
		{
			"name":        "分组平均上行流量",
			"code":        "groupAvgTrafficIn",
			"description": "当前节点所在分组在单位时间内接收的上行流量。",
			"unit":        "M",
		},

		{
			"name":        "集群平均请求数",
			"code":        "clusterAvgRequests",
			"description": "当前节点所在集群在单位时间内接收到的平均请求数。",
			"unit":        "个",
		},
		{
			"name":        "集群平均下行流量",
			"code":        "clusterAvgTrafficOut",
			"description": "当前节点所在集群在单位时间内发送的下行流量。",
			"unit":        "M",
		},
		{
			"name":        "集群平均上行流量",
			"code":        "clusterAvgTrafficIn",
			"description": "当前节点所在集群在单位时间内接收的上行流量。",
			"unit":        "M",
		},
	}
}

// NodeValueThresholdConfig 阈值列表
type NodeValueThresholdConfig struct {
	Id      int64                             `json:"id"`
	Items   []*NodeValueThresholdItemConfig   `json:"items"`
	Actions []*NodeValueThresholdActionConfig `json:"actions"`
}

// NodeValueThresholdItemConfig 阈值项目
type NodeValueThresholdItemConfig struct {
	Item         NodeValueItem         `json:"item"`
	Operator     NodeValueOperator     `json:"operator"`
	Value        float64               `json:"value"`
	Duration     int                   `json:"duration"`
	DurationUnit NodeValueDurationUnit `json:"durationUnit"`
	Options      maps.Map              `json:"options"` // 附加选项
}

type NodeValueThresholdActionConfig struct {
	Action  string   `json:"action"`
	Options maps.Map `json:"options"`
}

// NodeValueThresholdAction 动作
type NodeValueThresholdAction = string

const (
	NodeValueThresholdActionUp      NodeValueThresholdAction = "up"      // 上线
	NodeValueThresholdActionDown    NodeValueThresholdAction = "down"    // 下线
	NodeValueThresholdActionNotify  NodeValueThresholdAction = "notify"  // 通知
	NodeValueThresholdActionSwitch  NodeValueThresholdAction = "switch"  // 切换到备用IP
	NodeValueThresholdActionWebHook NodeValueThresholdAction = "webHook" // 调用外部Webhook
)
