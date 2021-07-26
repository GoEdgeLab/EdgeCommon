package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/configutils"

// HTTPRedirectToHTTPSConfig 跳转到HTTPS配置
type HTTPRedirectToHTTPSConfig struct {
	IsPrior bool   `yaml:"isPrior" json:"isPrior"` // 是否覆盖
	IsOn    bool   `yaml:"isOn" json:"isOn"`       // 是否开启
	Status  int    `yaml:"status" json:"status"`   // 跳转用的状态码
	Host    string `yaml:"host" json:"host"`       // 跳转后的Host
	Port    int    `yaml:"port" json:"port"`       // 跳转后的端口

	OnlyDomains   []string `yaml:"onlyDomains" json:"onlyDomains"`     // 允许的域名
	ExceptDomains []string `yaml:"exceptDomains" json:"exceptDomains"` // 排除的域名
}

// Init 初始化
func (this *HTTPRedirectToHTTPSConfig) Init() error {
	return nil
}

// MatchDomain 检查域名是否匹配
func (this *HTTPRedirectToHTTPSConfig) MatchDomain(domain string) bool {
	if len(this.ExceptDomains) > 0 && configutils.MatchDomains(this.ExceptDomains, domain) {
		return false
	}

	if len(this.OnlyDomains) > 0 && !configutils.MatchDomains(this.OnlyDomains, domain) {
		return false
	}

	return true
}
