// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dnsconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// NSDomainStatus 域名状态
type NSDomainStatus = string

const (
	NSDomainStatusNone      NSDomainStatus = "none"      // 初始状态
	NSDomainStatusVerified  NSDomainStatus = "verified"  // 已验证
	NSDomainStatusRejected  NSDomainStatus = "rejected"  // 已驳回（可以重新提交）
	NSDomainStatusForbidden NSDomainStatus = "forbidden" // 已禁止（禁止继续使用此域名）
)

func FindAllNSDomainStatusList() []*shared.Definition {
	return []*shared.Definition{
		{
			Name: "未验证",
			Code: NSDomainStatusNone,
		},
		{
			Name: "已验证",
			Code: NSDomainStatusVerified,
		},
		{
			Name: "已驳回",
			Code: NSDomainStatusRejected,
		},
		{
			Name: "已禁止",
			Code: NSDomainStatusForbidden,
		},
	}
}

func NSDomainStatusIsValid(status string) bool {
	for _, def := range FindAllNSDomainStatusList() {
		if def.Code == status {
			return true
		}
	}
	return false
}

func NSDomainStatusName(status string) string {
	for _, def := range FindAllNSDomainStatusList() {
		if def.Code == status {
			return def.Name
		}
	}
	return ""
}
