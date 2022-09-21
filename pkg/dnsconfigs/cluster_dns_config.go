package dnsconfigs

// ClusterDNSConfig 集群的DNS设置
type ClusterDNSConfig struct {
	CNAMERecords     []string `yaml:"cnameRecords" json:"cnameRecords"`         // 自动加入的CNAME
	TTL              int32    `yaml:"ttl" json:"ttl"`                           // 默认TTL，各个DNS服务商对记录的TTL的限制各有不同
	CNAMEAsDomain    bool     `yaml:"cnameAsDomain" json:"cnameAsDomain"`       // 是否可以像域名一样直接访问CNAME
	IncludingLnNodes bool     `yaml:"includingLnNodes" json:"includingLnNodes"` // 是否包含Ln节点

	NodesAutoSync   bool `yaml:"nodesAutoSync" json:"nodesAutoSync"`     // 是否自动同步节点状态
	ServersAutoSync bool `yaml:"serversAutoSync" json:"serversAutoSync"` // 是否自动同步服务状态
}
