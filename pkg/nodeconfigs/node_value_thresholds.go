// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

import "github.com/iwind/TeaGo/maps"

type IPAddressThresholdItem = string

const (
	IPAddressThresholdItemAvgRequests   IPAddressThresholdItem = "avgRequests"
	IPAddressThresholdItemAvgTrafficOut IPAddressThresholdItem = "avgTrafficOut" // M
	IPAddressThresholdItemAvgTrafficIn  IPAddressThresholdItem = "avgTrafficIn"  // M
	IPAddressThresholdItemConnectivity  IPAddressThresholdItem = "connectivity"  // 0-100
)

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
	NodeValueThresholdActionUp     NodeValueThresholdAction = "up"     // 上线
	NodeValueThresholdActionDown   NodeValueThresholdAction = "down"   // 下线
	NodeValueThresholdActionNotify NodeValueThresholdAction = "notify" // 通知
	NodeValueThresholdActionSwitch NodeValueThresholdAction = "switch" // 切换到备用IP
)
