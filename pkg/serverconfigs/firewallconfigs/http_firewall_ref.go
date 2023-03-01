package firewallconfigs

type HTTPFirewallRef struct {
	IsPrior           bool  `yaml:"isPrior" json:"isPrior"`
	IsOn              bool  `yaml:"isOn" json:"isOn"`
	FirewallPolicyId  int64 `yaml:"firewallPolicyId" json:"firewallPolicyId"`
	IgnoreGlobalRules bool  `yaml:"ignoreGlobalRules" json:"ignoreGlobalRules"` // 忽略系统定义的全局规则
}

func (this *HTTPFirewallRef) Init() error {
	return nil
}
