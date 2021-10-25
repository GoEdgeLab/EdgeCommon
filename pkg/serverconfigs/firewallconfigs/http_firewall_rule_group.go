package firewallconfigs

// HTTPFirewallRuleGroup 规则组
type HTTPFirewallRuleGroup struct {
	Id          int64                     `yaml:"id" json:"id"`
	IsOn        bool                      `yaml:"isOn" json:"isOn"`
	Name        string                    `yaml:"name" json:"name"`
	Description string                    `yaml:"description" json:"description"`
	Code        string                    `yaml:"code" json:"code"`
	SetRefs     []*HTTPFirewallRuleSetRef `yaml:"setRefs" json:"setRefs"`
	Sets        []*HTTPFirewallRuleSet    `yaml:"sets" json:"sets"`
}

// Init 初始化
func (this *HTTPFirewallRuleGroup) Init() error {
	for _, set := range this.Sets {
		err := set.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// AddRuleSet 添加规则集
func (this *HTTPFirewallRuleGroup) AddRuleSet(ruleSet *HTTPFirewallRuleSet) {
	this.Sets = append(this.Sets, ruleSet)
}

// FindRuleSet 根据ID查找规则集
func (this *HTTPFirewallRuleGroup) FindRuleSet(ruleSetId int64) *HTTPFirewallRuleSet {
	for _, set := range this.Sets {
		if set.Id == ruleSetId {
			return set
		}
	}
	return nil
}

// FindRuleSetWithCode 根据Code查找规则集
func (this *HTTPFirewallRuleGroup) FindRuleSetWithCode(code string) *HTTPFirewallRuleSet {
	for _, set := range this.Sets {
		if set.Code == code {
			return set
		}
	}
	return nil
}
