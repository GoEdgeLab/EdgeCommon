package nodeconfigs

// 节点状态
type NodeStatus struct {
	BuildVersion  string `json:"buildVersion"`  // 编译版本
	ConfigVersion int64  `json:"configVersion"` // 节点配置版本

	OS                    string  `json:"os"`
	Arch                  string  `json:"arch"`
	Hostname              string  `json:"hostname"`
	HostIP                string  `json:"hostIP"`
	CPUUsage              float64 `json:"cpuUsage"`
	CPULogicalCount       int     `json:"cpuLogicalCount"`
	CPUPhysicalCount      int     `json:"cpuPhysicalCount"`
	MemoryUsage           float64 `json:"memoryUsage"`
	MemoryTotal           uint64  `json:"memoryTotal"`
	DiskUsage             float64 `json:"diskUsage"`
	DiskMaxUsage          float64 `json:"diskMaxUsage"`
	DiskMaxUsagePartition string  `json:"diskMaxUsagePartition"`
	DiskTotal             uint64  `json:"diskTotal"`
	UpdatedAt             int64   `json:"updatedAt"`
	Load1m                float64 `json:"load1m"`
	Load5m                float64 `json:"load5m"`
	Load15m               float64 `json:"load15m"`
	ConnectionCount       int     `json:"connectionCount"` // 连接数

	IsActive bool   `json:"isActive"`
	Error    string `json:"error"`
}
