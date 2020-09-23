package serverconfigs

type HTTPFirewallRef struct {
	IsPrior          bool  `yaml:"isPrior" json:"isPrior"`
	IsOn             bool  `yaml:"isOn" json:"isOn"`
	FirewallPolicyId int64 `yaml:"firewallPolicyId" json:"firewallPolicyId"`
}
