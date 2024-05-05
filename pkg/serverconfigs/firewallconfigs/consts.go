// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package firewallconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/ipconfigs"
	"github.com/iwind/TeaGo/types"
)

const (
	GlobalBlackListId int64 = 2_000_000_000
	GlobalWhiteListId int64 = 2_000_000_001
	GlobalGreyListId  int64 = 2_000_000_002

	DefaultEventLevel = "critical"
)

func FindGlobalListIdWithType(listType ipconfigs.IPListType) int64 {
	switch listType {
	case ipconfigs.IPListTypeBlack:
		return GlobalBlackListId
	case ipconfigs.IPListTypeWhite:
		return GlobalWhiteListId
	case ipconfigs.IPListTypeGrey:
		return GlobalGreyListId
	}

	return 0
}

func FindGlobalListNameWithType(listType ipconfigs.IPListType) string {
	switch listType {
	case ipconfigs.IPListTypeBlack:
		return "全局黑名单"
	case ipconfigs.IPListTypeWhite:
		return "全局白名单"
	case ipconfigs.IPListTypeGrey:
		return "全局灰名单"
	}
	return "全局黑名单"
}

func IsGlobalListId(listId int64) bool {
	return listId == GlobalBlackListId || listId == GlobalWhiteListId || listId == GlobalGreyListId
}

func FindGlobalListIds() []int64 {
	return []int64{GlobalBlackListId, GlobalWhiteListId, GlobalGreyListId}
}

func FindGlobalListIdStrings() []string {
	return []string{types.String(GlobalBlackListId), types.String(GlobalWhiteListId), types.String(GlobalGreyListId)}
}
