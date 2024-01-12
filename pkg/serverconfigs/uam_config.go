// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
)

// UAMConfig UAM配置
type UAMConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`
	IsOn    bool `yaml:"isOn" json:"isOn"`

	AddToWhiteList    bool                           `yaml:"addToWhiteList" json:"addToWhiteList"`       // 是否将IP加入到白名单
	OnlyURLPatterns   []*shared.URLPattern           `yaml:"onlyURLPatterns" json:"onlyURLPatterns"`     // 仅限的URL
	ExceptURLPatterns []*shared.URLPattern           `yaml:"exceptURLPatterns" json:"exceptURLPatterns"` // 排除的URL
	MinQPSPerIP       int                            `yaml:"minQPSPerIP" json:"minQPSPerIP"`             // 启用要求的单IP最低平均QPS
	Conds             *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"`                         // 匹配条件
}

func NewUAMConfig() *UAMConfig {
	return &UAMConfig{
		AddToWhiteList: true,
	}
}

func (this *UAMConfig) Init() error {
	// only urls
	for _, pattern := range this.OnlyURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	// except urls
	for _, pattern := range this.ExceptURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	// conds
	if this.Conds != nil {
		err := this.Conds.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *UAMConfig) MatchURL(url string) bool {
	// except
	if len(this.ExceptURLPatterns) > 0 {
		for _, pattern := range this.ExceptURLPatterns {
			if pattern.Match(url) {
				return false
			}
		}
	}

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

func (this *UAMConfig) MatchRequest(formatter func(s string) string) bool {
	if this.Conds == nil {
		return true
	}
	return this.Conds.MatchRequest(formatter)
}
