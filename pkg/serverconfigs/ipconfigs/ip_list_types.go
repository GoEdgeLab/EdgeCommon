package ipconfigs

import "github.com/iwind/TeaGo/maps"

type IPListType = string

const (
	IPListTypeWhite IPListType = "white"
	IPListTypeBlack IPListType = "black"
	IPListTypeGrey  IPListType = "grey"
)

var IPListTypes = []maps.Map{
	{
		"name": "白名单",
		"code": IPListTypeWhite,
	},
	{
		"name": "黑名单",
		"code": IPListTypeBlack,
	},
	{
		"name": "灰名单",
		"code": IPListTypeGrey,
	},
}
