// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package firewallconfigs

type HTTPFirewallRecordIPAction struct {
	Type     string        `yaml:"type" json:"type"`
	IPListId int64         `yaml:"ipListId" json:"ipListId"`
	Level    string        `yaml:"level" json:"level"`
	Timeout  int32         `yaml:"timeout" json:"timeout"`
	Scope    FirewallScope `yaml:"scope" json:"scope"`
}
