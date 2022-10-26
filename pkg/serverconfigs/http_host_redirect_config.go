package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/types"
	"net/url"
	"regexp"
	"strings"
)

type HTTPHostRedirectType = string

const (
	HTTPHostRedirectTypeURL    HTTPHostRedirectType = "url"
	HTTPHostRedirectTypeDomain HTTPHostRedirectType = "domain"
	HTTPHostRedirectTypePort   HTTPHostRedirectType = "port"
)

// HTTPHostRedirectConfig 主机名跳转设置
type HTTPHostRedirectConfig struct {
	IsOn   bool `yaml:"isOn" json:"isOn"`     // 是否开启
	Status int  `yaml:"status" json:"status"` // 跳转用的状态码

	Type string               `yaml:"type" json:"type"` // 类型
	Mode HTTPHostRedirectType `yaml:"mode" json:"mode"` // 模式

	// URL跳转
	BeforeURL string `yaml:"beforeURL" json:"beforeURL"` // 跳转前的地址
	AfterURL  string `yaml:"afterURL" json:"afterURL"`   // 跳转后的地址

	MatchPrefix bool `yaml:"matchPrefix" json:"matchPrefix"` // 只匹配前缀部分
	MatchRegexp bool `yaml:"matchRegexp" json:"matchRegexp"` // 匹配正则表达式

	KeepRequestURI bool                           `yaml:"keepRequestURI" json:"keepRequestURI"` // 保留请求URI
	KeepArgs       bool                           `yaml:"keepArgs" json:"keepArgs"`             // 保留参数
	Conds          *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"`                   // 匹配条件

	realBeforeURL   string
	beforeURLRegexp *regexp.Regexp

	// 域名跳转
	DomainsAll        bool     `yaml:"domainAll" json:"domainsAll"`                // 所有域名都跳转
	DomainsBefore     []string `yaml:"domainsBefore" json:"domainsBefore"`         // 指定跳转之前的域名
	DomainAfter       string   `yaml:"domainAfter" json:"domainAfter"`             // 跳转之后的域名
	DomainAfterScheme string   `yaml:"domainAfterScheme" json:"domainAfterScheme"` // 跳转之后的协议
	//DomainRegexp  bool     `yaml:"domainRegexp" json:"domainRegexp"` // 使用正则匹配域名 TODO 暂时不实现
	//DomainKeepPort bool `yaml:"domainKeepPort" json:"domainKeepPort"` // 是否保持端口 TODO 暂时不实现
	//DomainNewPort int `yaml:"domainNewPort" json:"domainNewPort"` // 是否使用新端口 TODO 暂时不实现

	// 端口跳转
	PortsAll        bool     `yaml:"portsAll" json:"portsAll"`               // 所有端口
	PortsBefore     []string `yaml:"portsBefore" json:"portsBefore"`         // 跳转之前的端口：8080, 8080-8090
	PortAfter       int      `yaml:"port" json:"portAfter"`                  // 跳转之后的端口
	PortAfterScheme string   `yaml:"portAfterScheme" json:"portAfterScheme"` // 跳转之后的协议

	beforePortRanges [][2]int // [[from, to], {from2, to2}, ...]
}

// Init 初始化
func (this *HTTPHostRedirectConfig) Init() error {
	if len(this.Type) == 0 {
		this.Type = HTTPHostRedirectTypeURL
	}

	if this.Type == HTTPHostRedirectTypeURL {
		if !this.MatchRegexp {
			u, err := url.Parse(this.BeforeURL)
			if err != nil {
				return err
			}
			if len(u.Path) == 0 {
				this.realBeforeURL = this.BeforeURL + "/"
			} else {
				this.realBeforeURL = this.BeforeURL
			}
		} else if this.MatchRegexp {
			reg, err := regexp.Compile(this.BeforeURL)
			if err != nil {
				return err
			}
			this.beforeURLRegexp = reg
		}
	} else if this.Type == HTTPHostRedirectTypePort {
		this.beforePortRanges = [][2]int{}
		var portReg = regexp.MustCompile(`^\d+$`)
		var portRangeReg = regexp.MustCompile(`^\d+-\d+$`)
		for _, port := range this.PortsBefore {
			if portReg.MatchString(port) {
				this.beforePortRanges = append(this.beforePortRanges, [2]int{types.Int(port), types.Int(port)})
			} else if portRangeReg.MatchString(port) {
				var pieces = strings.Split(port, "-")
				if len(pieces) == 2 {
					var portFrom = types.Int(pieces[0])
					var portTo = types.Int(pieces[1])
					if portFrom > portTo {
						portFrom, portTo = portTo, portFrom
					}
					this.beforePortRanges = append(this.beforePortRanges, [2]int{portFrom, portTo})
				}
			}
		}
	}

	if this.Conds != nil {
		err := this.Conds.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// RealBeforeURL 跳转前URL
func (this *HTTPHostRedirectConfig) RealBeforeURL() string {
	return this.realBeforeURL
}

// BeforeURLRegexp 跳转前URL正则表达式
func (this *HTTPHostRedirectConfig) BeforeURLRegexp() *regexp.Regexp {
	return this.beforeURLRegexp
}

// MatchRequest 判断请求是否符合条件
func (this *HTTPHostRedirectConfig) MatchRequest(formatter func(source string) string) bool {
	if this.Conds == nil || !this.Conds.HasRequestConds() {
		return true
	}
	return this.Conds.MatchRequest(formatter)
}

// ContainsPort 是否包含端口
func (this *HTTPHostRedirectConfig) ContainsPort(reqPort int) bool {
	for _, port := range this.beforePortRanges {
		if port[0] <= reqPort && port[1] >= reqPort {
			return true
		}
	}
	return false
}
