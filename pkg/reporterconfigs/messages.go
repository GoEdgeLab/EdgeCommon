package reporterconfigs

type MessageCode = string

// 节点相关消息
const (
	MessageCodeConnectedAPINode    MessageCode = "connectedAPINode"    // 节点连接API节点成功
	MessageCodeCheckSystemdService MessageCode = "checkSystemdService" // 检查Systemd服务
	MessageCodeNewNodeTask         MessageCode = "newNodeTask"         // 有新的节点任务产生
)

// ConnectedAPINodeMessage 连接API节点成功
type ConnectedAPINodeMessage struct {
	APINodeId int64 `json:"apiNodeId"`
}

// CheckSystemdServiceMessage Systemd服务
type CheckSystemdServiceMessage struct {
}

// NewNodeTaskMessage 有新的节点任务
type NewNodeTaskMessage struct {
}
