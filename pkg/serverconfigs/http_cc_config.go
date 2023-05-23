// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package serverconfigs

type HTTPCCThreshold struct {
	// stub
}

func NewHTTPCCThreshold() *HTTPCCThreshold {
	return &HTTPCCThreshold{}
}

func (this *HTTPCCThreshold) Merge(threshold *HTTPCCThreshold) {
	// stub
}

func (this *HTTPCCThreshold) MergeIfEmpty(threshold *HTTPCCThreshold) {
	// stub
}

func (this *HTTPCCThreshold) Clone() *HTTPCCThreshold {
	return &HTTPCCThreshold{}
}

var DefaultHTTPCCThresholds = []*HTTPCCThreshold{} // stub

// DefaultHTTPCCConfig 默认的CC配置
func DefaultHTTPCCConfig() *HTTPCCConfig {
	return &HTTPCCConfig{}
}

// HTTPCCConfig HTTP CC防护配置
type HTTPCCConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"` // 是否覆盖父级
	IsOn    bool `yaml:"isOn" json:"isOn"`       // 是否启用
}

func NewHTTPCCConfig() *HTTPCCConfig {
	return &HTTPCCConfig{}
}

func (this *HTTPCCConfig) Init() error {
	return nil
}

func (this *HTTPCCConfig) MatchURL(url string) bool {
	return false
}
