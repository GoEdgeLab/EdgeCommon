// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package firewallconfigs

type HTTPFirewallPost307Action struct {
	Life  int32         `yaml:"life" json:"life"`
	Scope FirewallScope `yaml:"scope" json:"scope"`
}
