package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"net/url"
	"regexp"
)

// HTTPHostRedirectConfig 主机名跳转设置
type HTTPHostRedirectConfig struct {
	IsOn   bool `yaml:"isOn" json:"isOn"`     // 是否开启
	Status int  `yaml:"status" json:"status"` // 跳转用的状态码

	BeforeURL string `yaml:"beforeURL" json:"beforeURL"` // 跳转前的地址
	AfterURL  string `yaml:"afterURL" json:"afterURL"`   // 跳转后的地址

	MatchPrefix    bool                           `yaml:"matchPrefix" json:"matchPrefix"`       // 只匹配前缀部分
	MatchRegexp    bool                           `yaml:"matchRegexp" json:"matchRegexp"`       // 匹配正则表达式
	KeepRequestURI bool                           `yaml:"keepRequestURI" json:"keepRequestURI"` // 保留请求URI
	Conds          *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"`                   // 匹配条件

	realBeforeURL   string
	beforeURLRegexp *regexp.Regexp
}

// Init 初始化
func (this *HTTPHostRedirectConfig) Init() error {
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
