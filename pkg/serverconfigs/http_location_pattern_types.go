package serverconfigs

import "github.com/iwind/TeaGo/maps"

// HTTPLocationPatternType 匹配类型
type HTTPLocationPatternType = int

// 内置的匹配类型定义
const (
	HTTPLocationPatternTypePrefix HTTPLocationPatternType = 1
	HTTPLocationPatternTypeSuffix HTTPLocationPatternType = 4
	HTTPLocationPatternTypeExact  HTTPLocationPatternType = 2
	HTTPLocationPatternTypeRegexp HTTPLocationPatternType = 3
)

// AllLocationPatternTypes 取得所有的匹配类型信息
func AllLocationPatternTypes() []maps.Map {
	return []maps.Map{
		{
			"name":        "匹配前缀",
			"type":        HTTPLocationPatternTypePrefix,
			"description": "带有此前缀的路径才会被匹配",
		},
		{
			"name":        "匹配后缀",
			"type":        HTTPLocationPatternTypeSuffix,
			"description": "带有此后缀的路径才会被匹配",
		},
		{
			"name":        "精准匹配",
			"type":        HTTPLocationPatternTypeExact,
			"description": "带此路径完全一样的路径才会被匹配",
		},
		{
			"name":        "正则表达式匹配",
			"type":        HTTPLocationPatternTypeRegexp,
			"description": "通过正则表达式来匹配路径，<a href=\"https://goedge.cn/docs/Appendix/Regexp/Index.md\" target=\"_blank\">正则表达式语法 &raquo;</a>",
		},
	}
}

// FindLocationPatternType 查找单个匹配类型信息
func FindLocationPatternType(patternType int) maps.Map {
	for _, t := range AllLocationPatternTypes() {
		if t["type"] == patternType {
			return t
		}
	}
	return nil
}

// FindLocationPatternTypeName 查找单个匹配类型名称
func FindLocationPatternTypeName(patternType int) string {
	t := FindLocationPatternType(patternType)
	if t == nil {
		return ""
	}
	return t["name"].(string)
}
