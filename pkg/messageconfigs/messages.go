package messageconfigs

type MessageCode = string

const (
	MessageCodeConnectedAPINode    MessageCode = "connectedAPINode"    // 边缘节点连接API节点成功
	MessageCodeWriteCache          MessageCode = "writeCache"          // 写入缓存
	MessageCodeReadCache           MessageCode = "readCache"           // 读取缓存
	MessageCodeStatCache           MessageCode = "statCache"           // 统计缓存
	MessageCodePurgeCache          MessageCode = "purgeCache"          // 删除缓存
	MessageCodeCleanCache          MessageCode = "cleanCache"          // 清理缓存
	MessageCodePreheatCache        MessageCode = "preheatCache"        // 预热缓存
	MessageCodeCheckSystemdService MessageCode = "checkSystemdService" // 检查Systemd服务
	MessageCodeCheckLocalFirewall  MessageCode = "checkLocalFirewall"  // 检查本地防火墙
	MessageCodeNewNodeTask         MessageCode = "newNodeTask"         // 有新的节点任务产生
	MessageCodeChangeAPINode       MessageCode = "changeAPINode"       // 改变新的API节点
)

// ConnectedAPINodeMessage 连接API节点成功
type ConnectedAPINodeMessage struct {
	APINodeId int64 `json:"apiNodeId"`
}

// WriteCacheMessage 写入缓存
type WriteCacheMessage struct {
	CachePolicyJSON []byte `json:"cachePolicyJSON"`
	Key             string `json:"key"`
	Value           []byte `json:"value"`
	LifeSeconds     int64  `json:"lifeSeconds"`
}

// ReadCacheMessage 读取缓存
type ReadCacheMessage struct {
	CachePolicyJSON []byte `json:"cachePolicyJSON"`
	Key             string `json:"key"`
}

// StatCacheMessage 统计缓存
type StatCacheMessage struct {
	CachePolicyJSON []byte `json:"cachePolicyJSON"`
}

// 清除缓存

type CleanCacheMessage struct {
	CachePolicyJSON []byte `json:"cachePolicyJSON"`
}

// PurgeCacheMessageType 删除缓存
type PurgeCacheMessageType = string

const (
	PurgeCacheMessageTypeFile PurgeCacheMessageType = "file"
	PurgeCacheMessageTypeDir  PurgeCacheMessageType = "dir"
)

type PurgeCacheMessage struct {
	CachePolicyJSON []byte                `json:"cachePolicyJSON"`
	Keys            []string              `json:"keys"`
	Type            PurgeCacheMessageType `json:"type"` // 清理类型
}

// PreheatCacheMessage 预热缓存
type PreheatCacheMessage struct {
	CachePolicyJSON []byte   `json:"cachePolicyJSON"`
	Keys            []string `json:"keys"`
}

// CheckSystemdServiceMessage Systemd服务
type CheckSystemdServiceMessage struct {
}

// CheckLocalFirewallMessage 检查本地防火墙
type CheckLocalFirewallMessage struct {
	Name string `json:"name"`
}

// NewNodeTaskMessage 有新的节点任务
type NewNodeTaskMessage struct {
}

// ChangeAPINodeMessage 修改API地址
type ChangeAPINodeMessage struct {
	Addr string `json:"addr"`
}
