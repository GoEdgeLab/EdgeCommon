// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/configutils"

// NewReferersConfig 获取新防盗链配置对象
func NewReferersConfig() *ReferersConfig {
	return &ReferersConfig{
		CheckOrigin: true,
	}
}

// ReferersConfig 防盗链设置
type ReferersConfig struct {
	IsPrior         bool     `yaml:"isPrior" json:"isPrior"`
	IsOn            bool     `yaml:"isOn" json:"isOn"`
	AllowEmpty      bool     `yaml:"allowEmpty" json:"allowEmpty"`           // 来源域名允许为空
	AllowSameDomain bool     `yaml:"allowSameDomain" json:"allowSameDomain"` // 允许来源域名和当前访问的域名一致，相当于在站内访问
	AllowDomains    []string `yaml:"allowDomains" json:"allowDomains"`       // 允许的来源域名列表
	DenyDomains     []string `yaml:"denyDomains" json:"denyDomains"`         // 禁止的来源域名列表
	CheckOrigin     bool     `yaml:"checkOrigin" json:"checkOrigin"`         // 是否检查Origin
}

func (this *ReferersConfig) Init() error {
	return nil
}

func (this *ReferersConfig) MatchDomain(requestDomain string, refererDomain string) bool {
	if len(refererDomain) == 0 {
		if this.AllowEmpty {
			return true
		}
		return false
	}

	if this.AllowSameDomain && requestDomain == refererDomain {
		return true
	}

	if len(this.AllowDomains) == 0 {
		if len(this.DenyDomains) > 0 {
			return !configutils.MatchDomains(this.DenyDomains, refererDomain)
		}
		return false
	}

	if configutils.MatchDomains(this.AllowDomains, refererDomain) {
		if len(this.DenyDomains) > 0 && configutils.MatchDomains(this.DenyDomains, refererDomain) {
			return false
		}
		return true
	}

	return false
}
