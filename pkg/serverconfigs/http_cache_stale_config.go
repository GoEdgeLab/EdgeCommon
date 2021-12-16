// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// HTTPCacheStaleConfig Stale策略配置
type HTTPCacheStaleConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`
	IsOn    bool `yaml:"isOn" json:"isOn"` // 是否启用

	Status                    []int                `yaml:"status" json:"status"`                                       // 状态列表
	SupportStaleIfErrorHeader bool                 `yaml:"supportStaleIfErrorHeader" json:"supportStaleIfErrorHeader"` // 是否支持stale-if-error
	Life                      *shared.TimeDuration `yaml:"life" json:"life"`                                           // 陈旧内容生命周期
}
