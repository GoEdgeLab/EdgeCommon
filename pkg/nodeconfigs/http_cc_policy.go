// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package nodeconfigs

// HTTPCCPolicy CC策略
type HTTPCCPolicy struct {
	IsOn bool `json:"isOn" yaml:"isOn"`
}

func NewHTTPCCPolicy() *HTTPCCPolicy {
	return &HTTPCCPolicy{
		IsOn: true,
	}
}

func (this *HTTPCCPolicy) Init() error {
	return nil
}
