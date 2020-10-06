package firewallconfigs

// 防火墙策略
type HTTPFirewallPolicy struct {
	Id          int64                       `yaml:"id" json:"id"`
	IsOn        bool                        `yaml:"isOn" json:"isOn"`
	Name        string                      `yaml:"name" json:"name"`
	Description string                      `yaml:"description" json:"description"`
	Inbound     *HTTPFirewallInboundConfig  `yaml:"inbound" json:"inbound"`
	Outbound    *HTTPFirewallOutboundConfig `yaml:"outbound" json:"outbound"`
}

// 初始化
func (this *HTTPFirewallPolicy) Init() error {
	if this.Inbound != nil {
		err := this.Inbound.Init()
		if err != nil {
			return err
		}
	}
	if this.Outbound != nil {
		err := this.Outbound.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 获取所有分组
func (this *HTTPFirewallPolicy) AllRuleGroups() []*HTTPFirewallRuleGroup {
	result := []*HTTPFirewallRuleGroup{}
	if this.Inbound != nil {
		result = append(result, this.Inbound.Groups...)
	}
	if this.Outbound != nil {
		result = append(result, this.Outbound.Groups...)
	}
	return result
}

// 根据代号获取分组
func (this *HTTPFirewallPolicy) FindRuleGroupWithCode(code string) *HTTPFirewallRuleGroup {
	for _, g := range this.AllRuleGroups() {
		if g.Code == code {
			return g
		}
	}
	return nil
}
