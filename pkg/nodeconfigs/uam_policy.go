// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodeconfigs

func init() {
	_ = DefaultUAMPolicy.Init()
}

var DefaultUAMPolicy = &UAMPolicy{
	IsOn:               true,
	AllowSearchEngines: true,
	DenySpiders:        true,
	UITitle:            "",
	UIBody:             "",
}

type UAMPolicy struct {
	IsOn               bool `yaml:"isOn" json:"isOn"`
	AllowSearchEngines bool `yaml:"allowSearchEngines" json:"allowSearchEngines"` // 直接跳过常见搜索引擎
	DenySpiders        bool `yaml:"denySpiders" json:"denySpiders"`               // 拦截常见爬虫

	UITitle string `yaml:"uiTitle" json:"uiTitle"` // 页面标题
	UIBody  string `yaml:"uiBody" json:"uiBody"`   // 页面内容
}

func (this *UAMPolicy) Init() error {
	return nil
}
