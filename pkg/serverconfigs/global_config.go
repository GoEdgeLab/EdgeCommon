package serverconfigs

const (
	DefaultTCPPortRangeMin = 10000
	DefaultTCPPortRangeMax = 40000
)

// GlobalConfig 服务相关的全局设置
type GlobalConfig struct {
	// HTTP & HTTPS相关配置
	HTTPAll struct {
		MatchDomainStrictly  bool                  `yaml:"matchDomainStrictly" json:"matchDomainStrictly"`   // 是否严格匹配域名
		AllowMismatchDomains []string              `yaml:"allowMismatchDomains" json:"allowMismatchDomains"` // 允许的不匹配的域名
		DefaultDomain        string                `yaml:"defaultDomain" json:"defaultDomain"`               // 默认的域名
		DomainMismatchAction *DomainMismatchAction `yaml:"domainMismatchAction" json:"domainMismatchAction"` // 不匹配时采取的动作
		DomainAuditingIsOn   bool                  `yaml:"domainAuditingIsOn" json:"domainAuditingIsOn"`     // 域名是否需要审核
		DomainAuditingPrompt string                `yaml:"domainAuditingPrompt" json:"domainAuditingPrompt"` // 域名审核的提示
	} `yaml:"httpAll" json:"httpAll"`

	HTTP   struct{} `yaml:"http" json:"http"`
	HTTPS  struct{} `yaml:"https" json:"https"`
	TCPAll struct {
		PortRangeMin int   `yaml:"portRangeMin" json:"portRangeMin"` // 最小端口
		PortRangeMax int   `yaml:"portRangeMax" json:"portRangeMax"` // 最大端口
		DenyPorts    []int `yaml:"denyPorts" json:"denyPorts"`       // 禁止使用的端口
	} `yaml:"tcpAll" json:"tcpAll"`
	TCP  struct{} `yaml:"tcp" json:"tcp"`
	TLS  struct{} `yaml:"tls" json:"tls"`
	Unix struct{} `yaml:"unix" json:"unix"`
	UDP  struct{} `yaml:"udp" json:"udp"`

	// IP库相关配置
	IPLibrary struct {
		Code string `yaml:"code" json:"code"` // 当前使用的IP库代号
	} `yaml:"ipLibrary" json:"ipLibrary"`
}

func (this *GlobalConfig) Init() error {
	// HTTPAll
	if this.HTTPAll.DomainMismatchAction != nil {
		err := this.HTTPAll.DomainMismatchAction.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
