package messageconfigs

type NSMessageCode = string

// NS节点相关消息
const (
	NSMessageCodeConnectedAPINode    NSMessageCode = "connectedAPINode"    // NS节点连接API节点成功
	NSMessageCodeCheckSystemdService NSMessageCode = "checkSystemdService" // 检查Systemd服务
	NSMessageCodeNewNodeTask         MessageCode   = "newNodeTask"         // 有新的节点任务产生
)

// NSConnectedAPINodeMessage 连接API节点成功
type NSConnectedAPINodeMessage struct {
	APINodeId int64 `json:"apiNodeId"`
}

// NSCheckSystemdServiceMessage Systemd服务
type NSCheckSystemdServiceMessage struct {
}

// NewNSNodeTaskMessage 有新的节点任务
type NewNSNodeTaskMessage struct {
}
