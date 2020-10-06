package firewallconfigs

type HTTPFirewallRuleRef struct {
	IsOn   bool  `yaml:"isOn" json:"isOn"`
	RuleId int64 `yaml:"ruleId" json:"ruleId"`
}
