// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package userconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type UserTicketStatus = string

const (
	UserTicketStatusNone   UserTicketStatus = "none"
	UserTicketStatusSolved UserTicketStatus = "solved"
	UserTicketStatusClosed UserTicketStatus = "closed"
)

func UserTicketStatusName(status UserTicketStatus) string {
	switch status {
	case UserTicketStatusNone:
		return "进行中"
	case UserTicketStatusSolved:
		return "已解决"
	case UserTicketStatusClosed:
		return "已关闭"
	}
	return ""
}

func FindAllUserTicketStatusList() []*shared.Definition {
	return []*shared.Definition{
		{
			Name: "进行中",
			Code: UserTicketStatusNone,
		},
		{
			Name: "已解决",
			Code: UserTicketStatusSolved,
		},
		{
			Name: "已关闭",
			Code: UserTicketStatusClosed,
		},
	}
}
