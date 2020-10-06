package firewallconfigs

import "github.com/iwind/TeaGo/maps"

// 规则集定义
type HTTPFirewallRuleSet struct {
	Id            int64                  `yaml:"id" json:"id"`
	IsOn          bool                   `yaml:"isOn" json:"isOn"`
	Name          string                 `yaml:"name" json:"name"`
	Code          string                 `yaml:"code" json:"code"`
	Description   string                 `yaml:"description" json:"description"`
	Connector     string                 `yaml:"connector" json:"connector"`
	RuleRefs      []*HTTPFirewallRuleRef `yaml:"ruleRefs" json:"ruleRefs"`
	Rules         []*HTTPFirewallRule    `yaml:"rules" json:"rules"`
	Action        string                 `yaml:"action" json:"action"`
	ActionOptions maps.Map               `yaml:"actionOptions" json:"actionOptions"`
}

// 初始化
func (this *HTTPFirewallRuleSet) Init() error {
	for _, rule := range this.Rules {
		err := rule.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// 添加规则
func (this *HTTPFirewallRuleSet) AddRule(rule *HTTPFirewallRule) {
	this.Rules = append(this.Rules, rule)
}
