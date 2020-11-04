package serverconfigs

import "github.com/iwind/TeaGo/maps"

type IPLibraryType = string

const (
	IPLibraryTypeIP2Region IPLibraryType = "ip2region"
)

// 所有的IP库类型
var IPLibraryTypes = []maps.Map{
	{
		"name":        "ip2region",
		"code":        IPLibraryTypeIP2Region,
		"description": "一个开源的IP库：https://github.com/lionsoul2014/ip2region",
		"ext":         ".db",
	},
}

var DefaultIPLibraryType = IPLibraryTypeIP2Region

// 根据类型查找IP库
func FindIPLibraryWithType(libraryType string) maps.Map {
	for _, t := range IPLibraryTypes {
		if t.GetString("code") == libraryType {
			return t
		}
	}
	return nil
}
