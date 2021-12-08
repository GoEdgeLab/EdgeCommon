// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
)

// HTTPExpiresTimeConfig 发送到客户端的过期时间设置
type HTTPExpiresTimeConfig struct {
	IsPrior       bool                 `yaml:"isPrior" json:"isPrior"`             // 是否覆盖父级设置
	IsOn          bool                 `yaml:"isOn" json:"isOn"`                   // 是否启用
	Overwrite     bool                 `yaml:"overwrite" json:"overwrite"`         // 是否覆盖
	AutoCalculate bool                 `yaml:"autoCalculate" json:"autoCalculate"` // 是否自动计算
	Duration      *shared.TimeDuration `yaml:"duration" json:"duration"`           // 周期
}
