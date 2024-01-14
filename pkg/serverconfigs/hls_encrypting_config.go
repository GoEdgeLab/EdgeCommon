// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// HLSEncryptingConfig HLS加密配置
type HLSEncryptingConfig struct {
	IsOn bool `yaml:"isOn" json:"isOn"` // 是否启用

	OnlyURLPatterns   []*shared.URLPattern `yaml:"onlyURLPatterns" json:"onlyURLPatterns"`     // 仅限的URL
	ExceptURLPatterns []*shared.URLPattern `yaml:"exceptURLPatterns" json:"exceptURLPatterns"` // 排除的URL
}

func (this *HLSEncryptingConfig) Init() error {
	for _, pattern := range this.OnlyURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	for _, pattern := range this.ExceptURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *HLSEncryptingConfig) MatchURL(url string) bool {
	// except
	if len(this.ExceptURLPatterns) > 0 {
		for _, pattern := range this.ExceptURLPatterns {
			if pattern.Match(url) {
				return false
			}
		}
	}

	// only
	if len(this.OnlyURLPatterns) > 0 {
		for _, pattern := range this.OnlyURLPatterns {
			if pattern.Match(url) {
				return true
			}
		}
		return false
	}

	return true
}
