// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package userconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// UserOrderConfig 用户订单配置
type UserOrderConfig struct {
	EnablePay       bool                 `json:"enablePay"`       // 启用支付
	DisablePageHTML string               `json:"disablePageHTML"` // 禁用支付时的页面提示
	OrderLife       *shared.TimeDuration `json:"orderLife"`       // 过期时间
}

func DefaultUserOrderConfig() *UserOrderConfig {
	return &UserOrderConfig{
		EnablePay:       false,
		DisablePageHTML: "暂不提供在线充值功能，请联系管理员充值。",
		OrderLife: &shared.TimeDuration{
			Count: 1,
			Unit:  shared.TimeDurationUnitHour,
		},
	}
}
