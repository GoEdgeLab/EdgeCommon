// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

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

type NodeValueRange = string

const (
	NodeValueRangeMinute NodeValueRange = "minute"
)
