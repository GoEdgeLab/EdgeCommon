// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

// 一组系统默认值

const (
	DefaultMaxThreads    = 20000   // 单节点最大线程数
	DefaultMaxThreadsMin = 1000    // 单节点最大线程数最小值
	DefaultMaxThreadsMax = 100_000 // 单节点最大线程数最大值

	DefaultTCPMaxConnections   = 100_000 // 单节点TCP最大连接数
	DefaultTCPLinger           = 3       // 单节点TCP Linger值
	DefaultTLSHandshakeTimeout = 3       // TLS握手超时时间
)
