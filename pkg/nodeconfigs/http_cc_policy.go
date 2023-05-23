// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package nodeconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"

// HTTPCCPolicy CC策略
type HTTPCCPolicy struct {
	IsOn       bool                             `json:"isOn" yaml:"isOn"`
	Thresholds []*serverconfigs.HTTPCCThreshold `json:"thresholds" yaml:"thresholds"` // 阈值
}

func NewHTTPCCPolicy() *HTTPCCPolicy {
	return &HTTPCCPolicy{
		IsOn: true,
	}
}

func (this *HTTPCCPolicy) Init() error {
	return nil
}
