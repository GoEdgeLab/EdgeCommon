// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"strings"
)

type HTTPBaseOptimizationConfig struct {
	OnlyURLPatterns   []*shared.URLPattern `yaml:"onlyURLPatterns" json:"onlyURLPatterns"`     // 仅限的URL
	ExceptURLPatterns []*shared.URLPattern `yaml:"exceptURLPatterns" json:"exceptURLPatterns"` // 排除的URL
}

func (this *HTTPBaseOptimizationConfig) Init() error {
	// only url
	for _, pattern := range this.OnlyURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	// except url
	for _, pattern := range this.ExceptURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *HTTPBaseOptimizationConfig) MatchURL(url string) bool {
	// 去除问号
	var index = strings.Index(url, "?")
	if index >= 0 {
		url = url[:index]
	}

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
