package serverconfigs

import (
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"regexp"
)

type HTTPRewriteMode = string

const (
	HTTPRewriteTargetProxy = 1
	HTTPRewriteTargetURL   = 2
)

const (
	HTTPRewriteModeRedirect HTTPRewriteMode = "redirect" // 跳转
	HTTPRewriteModeProxy    HTTPRewriteMode = "proxy"    // 代理
)

// 重写规则定义
//
// 参考
// - http://nginx.org/en/docs/http/ngx_http_rewrite_module.html
// - https://httpd.apache.org/docs/current/mod/mod_rewrite.html
// - https://httpd.apache.org/docs/2.4/rewrite/flags.html
// TODO 实现对其他代理服务的引用
type HTTPRewriteRule struct {
	Id   int64 `yaml:"id" json:"id"`     // ID
	IsOn bool  `yaml:"isOn" json:"isOn"` // 是否开启

	// 开启的条件
	// 语法为：cond param operator value 比如：
	// - cond ${status} gte 200
	// - cond ${arg.name} eq lily
	// - cond ${requestPath} regexp .*\.png
	Conds *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"` // 匹配条件 TODO

	// 规则
	// 语法为：pattern regexp 比如：
	// - pattern ^/article/(\d+).html
	Pattern string `yaml:"pattern" json:"pattern"`

	// 模式
	Mode           HTTPRewriteMode `yaml:"mode" json:"mode"`
	RedirectStatus int             `yaml:"redirectStatus" json:"redirectStatus"` // 跳转的状态码
	ProxyHost      string          `yaml:"proxyHost" json:"proxyHost"`           // 代理模式下的Host

	// 要替换成的URL
	// 支持反向引用：${0}, ${1}, ...，也支持?P<NAME>语法
	// - 如果以 proxy:// 开头，表示目标为代理，首先会尝试作为代理ID请求，如果找不到，会尝试作为代理Host请求
	Replace string `yaml:"replace" json:"replace"`

	// 其他选项
	IsBreak   bool `yaml:"isBreak" json:"isBreak"`     // 终止向下解析
	WithQuery bool `yaml:"withQuery" json:"withQuery"` // 是否保留QueryString

	reg                   *regexp.Regexp
	replaceHasVariables   bool
	proxyHostHasVariables bool
}

// 校验
func (this *HTTPRewriteRule) Init() error {
	this.replaceHasVariables = configutils.HasVariables(this.Replace)
	this.proxyHostHasVariables = configutils.HasVariables(this.ProxyHost)

	reg, err := regexp.Compile(this.Pattern)
	if err != nil {
		return err
	}
	this.reg = reg

	// 校验条件
	if this.Conds != nil {
		err := this.Conds.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 对某个请求执行规则
func (this *HTTPRewriteRule) MatchRequest(requestPath string, formatter func(source string) string) (replace string, varMapping map[string]string, matched bool) {
	if this.reg == nil {
		return "", nil, false
	}

	matches := this.reg.FindStringSubmatch(requestPath)
	if len(matches) == 0 {
		return "", nil, false
	}

	// 判断条件
	if this.Conds != nil && !this.Conds.MatchRequest(formatter) {
		return "", nil, false
	}

	varMapping = map[string]string{}
	subNames := this.reg.SubexpNames()
	for index, match := range matches {
		varMapping[fmt.Sprintf("%d", index)] = match
		subName := subNames[index]
		if len(subName) > 0 {
			varMapping[subName] = match
		}
	}

	replace = configutils.ParseVariables(this.Replace, func(varName string) string {
		v, ok := varMapping[varName]
		if ok {
			return v
		}
		return "${" + varName + "}"
	})

	if this.replaceHasVariables {
		replace = formatter(replace)
	}

	return replace, varMapping, true
}

// 判断是否是外部URL
func (this *HTTPRewriteRule) IsExternalURL(url string) bool {
	return shared.RegexpExternalURL.MatchString(url)
}

// 判断ProxyHost是否有变量
func (this *HTTPRewriteRule) ProxyHostHasVariables() bool {
	return this.proxyHostHasVariables
}
