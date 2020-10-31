package serverconfigs

// 服务相关的全局设置
type GlobalConfig struct {
	HTTPAll struct {
		MatchDomainStrictly bool `yaml:"matchDomainStrictly" json:"matchDomainStrictly"`
	} `yaml:"httpAll" json:"httpAll"`
	HTTP   struct{} `yaml:"http" json:"http"`
	HTTPS  struct{} `yaml:"https" json:"https"`
	TCPAll struct{} `yaml:"tcpAll" json:"tcpAll"`
	TCP    struct{} `yaml:"tcp" json:"tcp"`
	TLS    struct{} `yaml:"tls" json:"tls"`
	Unix   struct{} `yaml:"unix" json:"unix"`
	UDP    struct{} `yaml:"udp" json:"udp"`
}

func (this *GlobalConfig) Init() error {
	return nil
}
