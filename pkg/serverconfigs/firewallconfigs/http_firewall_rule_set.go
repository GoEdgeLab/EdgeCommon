package firewallconfigs

import (
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/maps"
)

// HTTPFirewallActionConfig 单个动作配置
type HTTPFirewallActionConfig struct {
	Code    HTTPFirewallActionString `yaml:"code" json:"code"`
	Options maps.Map                 `yaml:"options" json:"options"`
}

// HTTPFirewallRuleSet 规则集定义
type HTTPFirewallRuleSet struct {
	Id          int64                  `yaml:"id" json:"id"`
	IsOn        bool                   `yaml:"isOn" json:"isOn"`
	Name        string                 `yaml:"name" json:"name"`
	Code        string                 `yaml:"code" json:"code"`
	Description string                 `yaml:"description" json:"description"`
	Connector   string                 `yaml:"connector" json:"connector"`
	RuleRefs    []*HTTPFirewallRuleRef `yaml:"ruleRefs" json:"ruleRefs"`
	Rules       []*HTTPFirewallRule    `yaml:"rules" json:"rules"`

	Actions []*HTTPFirewallActionConfig `yaml:"actions" json:"actions"`

	//Action        string   `yaml:"action" json:"action"`               // deprecated, v0.2.5
	//ActionOptions maps.Map `yaml:"actionOptions" json:"actionOptions"` // deprecated, v0.2.5
}

// Init 初始化
func (this *HTTPFirewallRuleSet) Init() error {
	for _, rule := range this.Rules {
		err := rule.Init()
		if err != nil {
			logs.Println("ERROR", "validate rule '"+rule.Summary()+"' failed: "+err.Error())

			// 这里不阻断执行，因为先前有些用户填写了错误的规则
		}
	}

	return nil
}

// AddRule 添加规则
func (this *HTTPFirewallRuleSet) AddRule(rule *HTTPFirewallRule) {
	this.Rules = append(this.Rules, rule)
}
