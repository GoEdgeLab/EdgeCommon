package serverconfigs

import "github.com/iwind/TeaGo/maps"

type CachePolicyStorageType = string

const (
	CachePolicyStorageFile   CachePolicyStorageType = "file"
	CachePolicyStorageMemory CachePolicyStorageType = "memory"
)

var AllCachePolicyStorageTypes = []maps.Map{
	{
		"name": "文件缓存",
		"type": CachePolicyStorageFile,
	},
	{
		"name": "内存缓存",
		"type": CachePolicyStorageMemory,
	},
}

// 根据类型查找名称
func FindCachePolicyStorageName(policyType CachePolicyStorageType) string {
	for _, t := range AllCachePolicyStorageTypes {
		if t.GetString("type") == policyType {
			return t.GetString("name")
		}
	}
	return ""
}
