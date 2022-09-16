package serverconfigs

const (
	DefaultTCPPortRangeMin = 10000
	DefaultTCPPortRangeMax = 40000
)

// GlobalConfig 服务相关的全局设置
// Deprecated
type GlobalConfig struct {
	// HTTP & HTTPS相关配置
	HTTPAll struct {
		//MatchDomainStrictly  bool                  `yaml:"matchDomainStrictly" json:"matchDomainStrictly"`   // 是否严格匹配域名
		//AllowMismatchDomains []string              `yaml:"allowMismatchDomains" json:"allowMismatchDomains"` // 允许的不匹配的域名
		//DefaultDomain        string                `yaml:"defaultDomain" json:"defaultDomain"`               // 默认的域名
		//DomainMismatchAction *DomainMismatchAction `yaml:"domainMismatchAction" json:"domainMismatchAction"` // 不匹配时采取的动作
		DomainAuditingIsOn   bool   `yaml:"domainAuditingIsOn" json:"domainAuditingIsOn"`     // 域名是否需要审核
		DomainAuditingPrompt string `yaml:"domainAuditingPrompt" json:"domainAuditingPrompt"` // 域名审核的提示
	} `yaml:"httpAll" json:"httpAll"`

	TCPAll struct {
		PortRangeMin int   `yaml:"portRangeMin" json:"portRangeMin"` // 最小端口
		PortRangeMax int   `yaml:"portRangeMax" json:"portRangeMax"` // 最大端口
		DenyPorts    []int `yaml:"denyPorts" json:"denyPorts"`       // 禁止使用的端口
	} `yaml:"tcpAll" json:"tcpAll"`
}

func (this *GlobalConfig) Init() error {
	return nil
}
