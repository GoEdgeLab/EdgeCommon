// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package monitorconfigs

type RateConfig struct {
	Minutes int32 `json:"minutes"` // 周期分钟
	Count   int32 `json:"count"`   // 数量
}
