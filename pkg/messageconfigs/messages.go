package messageconfigs

type MessageCode = string

const (
	MessageCodeConnectedAPINode MessageCode = "connectedAPINode" // 边缘节点连接API节点成功
	MessageCodeWriteCache       MessageCode = "writeCache"       // 写入缓存
	MessageCodeReadCache        MessageCode = "readCache"        // 读取缓存
	MessageCodeStatCache        MessageCode = "statCache"        // 统计缓存
	MessageCodePurgeCache       MessageCode = "purgeCache"       // 删除缓存
	MessageCodeCleanCache       MessageCode = "cleanCache"       // 清理缓存
	MessageCodePreheatCache     MessageCode = "preheatCache"     // 预热缓存
	MessageCodeConfigChanged    MessageCode = "configChanged"    // 配置已改变
	MessageCodeIPListChanged    MessageCode = "ipListChanged"    // IP列表变化
)

// 连接API节点成功
type ConnectedAPINodeMessage struct {
	APINodeId int64 `json:"apiNodeId"`
}

// 写入缓存
type WriteCacheMessage struct {
	CachePolicyJSON []byte `json:"cachePolicyJSON"`
	Key             string `json:"key"`
	Value           []byte `json:"value"`
	LifeSeconds     int64  `json:"lifeSeconds"`
}

// 读取缓存
type ReadCacheMessage struct {
	CachePolicyJSON []byte `json:"cachePolicyJSON"`
	Key             string `json:"key"`
}

// 统计缓存
type StatCacheMessage struct {
	CachePolicyJSON []byte `json:"cachePolicyJSON"`
}

// 清除缓存
type CleanCacheMessage struct {
	CachePolicyJSON []byte `json:"cachePolicyJSON"`
}

// 删除缓存
type PurgeCacheMessage struct {
	CachePolicyJSON []byte   `json:"cachePolicyJSON"`
	Keys            []string `json:"keys"`
}

// 预热缓存
type PreheatCacheMessage struct {
	CachePolicyJSON []byte   `json:"cachePolicyJSON"`
	Keys            []string `json:"keys"`
}

// 配置已改变
type ConfigChangedMessage struct {
}

// IPList变化
type IPListChangedMessage struct {
}
