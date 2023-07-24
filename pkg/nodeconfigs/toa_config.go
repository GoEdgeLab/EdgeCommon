// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package nodeconfigs

// NewTOAConfig 默认的TOA配置
func NewTOAConfig() *TOAConfig {
	return &TOAConfig{}
}

// TOAConfig TOA相关配置
type TOAConfig struct {
	IsOn bool `yaml:"isOn" json:"isOn"`
}

func (this *TOAConfig) Init() error {
	return nil
}
