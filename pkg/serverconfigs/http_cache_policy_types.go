package serverconfigs

import "github.com/iwind/TeaGo/maps"

const (
	CachePolicyTypeFile   CachePolicyType = "file"
	CachePolicyTypeMemory CachePolicyType = "memory"
)

var AllCachePolicyTypes = []maps.Map{
	{
		"name": "文件缓存",
		"type": CachePolicyTypeFile,
	},
	{
		"name": "内存缓存",
		"type": CachePolicyTypeMemory,
	},
}

// 根据类型查找名称
func FindCachePolicyTypeName(policyType CachePolicyType) string {
	for _, t := range AllCachePolicyTypes {
		if t.GetString("type") == policyType {
			return t.GetString("name")
		}
	}
	return ""
}
