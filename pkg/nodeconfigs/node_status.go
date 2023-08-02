package nodeconfigs

// NodeStatus 节点状态
type NodeStatus struct {
	BuildVersion     string `json:"buildVersion"`     // 编译版本
	BuildVersionCode uint32 `json:"buildVersionCode"` // 版本数字
	ConfigVersion    int64  `json:"configVersion"`    // 节点配置版本

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
	DiskWritingSpeedMB    int     `json:"diskWritingSpeedMB"` // 硬盘写入速度
	UpdatedAt             int64   `json:"updatedAt"`
	Timestamp             int64   `json:"timestamp"` // 当前节点时间戳
	Load1m                float64 `json:"load1m"`
	Load5m                float64 `json:"load5m"`
	Load15m               float64 `json:"load15m"`
	ConnectionCount       int     `json:"connectionCount"`   // 连接数
	ExePath               string  `json:"exePath"`           // 可执行文件路径
	APISuccessPercent     float64 `json:"apiSuccessPercent"` // API成功比例
	APIAvgCostSeconds     float64 `json:"apiAvgCostSeconds"` // API平均耗时

	TrafficInBytes  uint64 `json:"trafficInBytes"`
	TrafficOutBytes uint64 `json:"trafficOutBytes"`

	CacheTotalDiskSize   int64 `json:"cacheTotalDiskSize"`
	CacheTotalMemorySize int64 `json:"cacheTotalMemorySize"`

	LocalFirewallName string `json:"localFirewallName"`

	IsActive bool   `json:"isActive"`
	Error    string `json:"error"`
}
