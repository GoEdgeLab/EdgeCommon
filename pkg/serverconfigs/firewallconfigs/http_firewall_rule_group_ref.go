package firewallconfigs

type HTTPFirewallRuleGroupRef struct {
	IsOn    bool  `yaml:"isOn" json:"isOn"`
	GroupId int64 `yaml:"groupId" json:"groupId"`
}
