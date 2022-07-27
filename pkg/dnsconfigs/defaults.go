// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

import "github.com/iwind/TeaGo/maps"

// 一组系统默认值
// 修改单个IP相关限制值时要考虑到NAT中每个IP会代表很多个主机，并非1对1的关系

const (
	DefaultMaxThreads    = 20000   // 单节点最大线程数
	DefaultMaxThreadsMin = 1000    // 单节点最大线程数最小值
	DefaultMaxThreadsMax = 100_000 // 单节点最大线程数最大值

	DefaultTCPMaxConnections        = 100_000 // 单节点TCP最大连接数
	DefaultTCPMaxConnectionsPerIP   = 1000    // 单IP最大连接数
	DefaultTCPMinConnectionsPerIP   = 5       // 单IP最小连接数
	DefaultTCPNewConnectionsRate    = 500     // 单IP连接速率限制（按分钟）
	DefaultTCPNewConnectionsMinRate = 5       // 单IP最小连接速率
	DefaultTCPLinger                = 3       // 单节点TCP Linger值
	DefaultTLSHandshakeTimeout      = 3       // TLS握手超时时间
)

var DefaultConfigs = maps.Map{
	"tcpMaxConnections":        DefaultTCPMaxConnections,
	"tcpMaxConnectionsPerIP":   DefaultTCPMaxConnectionsPerIP,
	"tcpMinConnectionsPerIP":   DefaultTCPMinConnectionsPerIP,
	"tcpNewConnectionsRate":    DefaultTCPNewConnectionsRate,
	"tcpNewConnectionsMinRate": DefaultTCPNewConnectionsMinRate,
}
