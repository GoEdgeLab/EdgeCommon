// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package userconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type OrderType = string

const (
	OrderTypeCharge OrderType = "charge"
)

func IsValidOrderType(s string) bool {
	return s == OrderTypeCharge
}

func OrderTypeName(orderType OrderType) string {
	switch orderType {
	case OrderTypeCharge:
		return "充值"
	}
	return ""
}

type OrderStatus = string

const (
	OrderStatusNone      OrderStatus = "none"
	OrderStatusCancelled OrderStatus = "cancelled"
	OrderStatusFinished  OrderStatus = "finished"
)

func OrderStatusName(orderStatus OrderStatus) string {
	switch orderStatus {
	case OrderStatusNone:
		return "未支付"
	case OrderStatusCancelled:
		return "已取消"
	case OrderStatusFinished:
		return "已完成"
	}
	return ""
}

func FindAllOrderStatusList() []*shared.Definition {
	return []*shared.Definition{
		{
			Name: "已完成",
			Code: OrderStatusFinished,
		},
		{
			Name: "未支付",
			Code: OrderStatusNone,
		},
		{
			Name: "已取消",
			Code: OrderStatusCancelled,
		},
	}
}
