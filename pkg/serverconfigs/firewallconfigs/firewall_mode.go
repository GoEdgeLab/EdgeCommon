// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package firewallconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// FirewallMode 模式
type FirewallMode = string

const (
	FirewallModeDefend  FirewallMode = "defend"  // 防御模式
	FirewallModeObserve FirewallMode = "observe" // 观察模式
	FirewallModePass    FirewallMode = "pass"    // 通过模式
)

func FindAllFirewallModes() []*shared.Definition {
	return []*shared.Definition{
		{
			Name:        "防御模式",
			Description: "执行正常的防御规则和相应动作。",
			Code:        FirewallModeDefend,
		},
		{
			Name:        "观察模式",
			Description: "执行正常的防御规则，但只记录日志，不执行动作。",
			Code:        FirewallModeObserve,
		},
		{
			Name:        "通过模式",
			Description: "不执行任何规则，所有的请求都将会直接通过。",
			Code:        FirewallModePass,
		},
	}
}

func FindFirewallMode(code FirewallMode) *shared.Definition {
	for _, def := range FindAllFirewallModes() {
		if def.Code == code {
			return def
		}
	}
	return nil
}
