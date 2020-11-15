package dnsconfigs

// 集群的DNS设置
type ClusterDNSConfig struct {
	NodesAutoSync   bool `yaml:"nodesAutoSync" json:"nodesAutoSync"`     // 是否自动同步节点状态
	ServersAutoSync bool `yaml:"serversAutoSync" json:"serversAutoSync"` // 是否自动同步服务状态
}
