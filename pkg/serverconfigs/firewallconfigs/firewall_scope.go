// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package firewallconfigs

type FirewallScope = string

const (
	FirewallScopeGlobal FirewallScope = "global"
	FirewallScopeServer FirewallScope = "service" // 历史原因，代号为 service 而非 server
)
