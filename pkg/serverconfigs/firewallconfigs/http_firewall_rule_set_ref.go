package firewallconfigs

type HTTPFirewallRuleSetRef struct {
	IsOn  bool  `yaml:"isOn" json:"isOn"`
	SetId int64 `yaml:"setId" json:"setId"`
}
