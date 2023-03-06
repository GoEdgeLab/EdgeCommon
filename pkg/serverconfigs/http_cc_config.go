// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

import "strings"

// HTTPCCConfig HTTP CC防护配置
type HTTPCCConfig struct {
	IsPrior            bool   `yaml:"isPrior" json:"isPrior"`                       // 是否覆盖父级
	IsOn               bool   `yaml:"isOn" json:"isOn"`                             // 是否启用
	WithRequestPath    bool   `yaml:"withRequestPath" json:"withRequestPath"`       // 根据URL路径区分请求
	PeriodSeconds      int32  `yaml:"periodSeconds" json:"periodSeconds"`           // 计算周期
	MaxRequests        int32  `yaml:"maxRequests" json:"maxRequests"`               // 请求数最大值
	MaxConnections     int32  `yaml:"maxConnections" json:"maxConnections"`         // 连接数最大值
	IgnoreCommonFiles  bool   `yaml:"ignoreCommonFiles" json:"ignoreCommonFiles"`   // 忽略常用文件，如CSS、JS等
	IgnoreCommonAgents bool   `yaml:"ignoreCommonAgents" json:"ignoreCommonAgents"` // 忽略常见搜索引擎等
	Action             string `yaml:"action" json:"action"`                         // 动作，比如block、captcha等

	fullKey string
}

func NewHTTPCCConfig() *HTTPCCConfig {
	return &HTTPCCConfig{
		WithRequestPath:    false,
		PeriodSeconds:      10,
		MaxRequests:        60,
		MaxConnections:     10,
		IgnoreCommonFiles:  false,
		IgnoreCommonAgents: true,
		Action:             "captcha",
	}
}

func (this *HTTPCCConfig) Init() error {
	// 组合Key
	var keys = []string{"${remoteAddr}"}
	if this.WithRequestPath {
		keys = append(keys, "${requestPath}")
	}
	this.fullKey = strings.Join(keys, "@")

	return nil
}

func (this *HTTPCCConfig) Key() string {
	return this.fullKey
}
