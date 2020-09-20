package serverconfigs

type HTTPFirewallRef struct {
	IsOn             bool  `yaml:"isOn" json:"isOn"`
	FirewallPolicyId int64 `yaml:"firewallPolicyId" json:"firewallPolicyId"`
}
