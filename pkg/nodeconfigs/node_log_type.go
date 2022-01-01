// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

type NodeLogType = string

const (
	NodeLogTypeListenAddressFailed    NodeLogType = "listenAddressFailed"
	NodeLogTypeServerConfigInitFailed NodeLogType = "serverConfigInitFailed"
	NodeLogTypeRunScriptFailed        NodeLogType = "runScriptFailed"
)
