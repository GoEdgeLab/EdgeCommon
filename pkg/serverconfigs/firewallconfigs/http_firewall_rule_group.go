package firewallconfigs

// 规则组
type HTTPFirewallRuleGroup struct {
	Id          int64                     `yaml:"id" json:"id"`
	IsOn        bool                      `yaml:"isOn" json:"isOn"`
	Name        string                    `yaml:"name" json:"name"`
	Description string                    `yaml:"description" json:"description"`
	Code        string                    `yaml:"code" json:"code"`
	SetRefs     []*HTTPFirewallRuleSetRef `yaml:"setRefs" json:"setRefs"`
	Sets        []*HTTPFirewallRuleSet    `yaml:"sets" json:"sets"`
}

// 初始化
func (this *HTTPFirewallRuleGroup) Init() error {
	for _, set := range this.Sets {
		err := set.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// 添加规则集
func (this *HTTPFirewallRuleGroup) AddRuleSet(ruleSet *HTTPFirewallRuleSet) {
	this.Sets = append(this.Sets, ruleSet)
}
