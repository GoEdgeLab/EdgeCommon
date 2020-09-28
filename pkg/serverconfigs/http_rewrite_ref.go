package serverconfigs

// 重写规则的引用
type HTTPRewriteRef struct {
	IsOn          bool  `yaml:"isOn" json:"isOn"`                   // 是否启用
	RewriteRuleId int64 `yaml:"rewriteRuleId" json:"rewriteRuleId"` // 规则ID
}
