package firewallconfigs

type HTTPFirewallOutboundConfig struct {
	IsOn      bool                        `yaml:"isOn" json:"isOn"`
	GroupRefs []*HTTPFirewallRuleGroupRef `yaml:"groupRefs" json:"groupRefs"`
	Groups    []*HTTPFirewallRuleGroup    `yaml:"groups" json:"groups"`
}

// 初始化
func (this *HTTPFirewallOutboundConfig) Init() error {
	for _, group := range this.Groups {
		err := group.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// 根据Code查找Group
func (this *HTTPFirewallOutboundConfig) FindGroupWithCode(code string) *HTTPFirewallRuleGroup {
	for _, group := range this.Groups {
		if group.Code == code {
			return group
		}
	}
	return nil
}
