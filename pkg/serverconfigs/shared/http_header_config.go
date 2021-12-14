package shared

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"regexp"
	"strings"
)

// HTTPHeaderReplaceValue 值替换定义
type HTTPHeaderReplaceValue struct {
	Pattern           string `yaml:"pattern" json:"pattern"`
	Replacement       string `yaml:"replacement" json:"replacement"`
	IsCaseInsensitive bool   `yaml:"isCaseInsensitive" json:"isCaseInsensitive"` // TODO
	IsRegexp          bool   `yaml:"isRegexp" json:"isRegexp"`                   // TODO

	patternReg *regexp.Regexp
}

func (this *HTTPHeaderReplaceValue) Init() error {
	if this.IsRegexp {
		var pattern = this.Pattern
		if this.IsCaseInsensitive && !strings.HasPrefix(pattern, "(?i)") {
			pattern = "(?i)" + pattern
		}

		reg, err := regexp.Compile(pattern)
		if err != nil {
			return err
		}

		// TODO 支持匹配名（${name}）和反向引用${1}。。。
		this.patternReg = reg
	} else {
		if this.IsCaseInsensitive {
			var pattern = "(?i)" + regexp.QuoteMeta(this.Pattern)
			reg, err := regexp.Compile(pattern)
			if err != nil {
				return err
			}
			this.patternReg = reg
		}
	}
	return nil
}

func (this *HTTPHeaderReplaceValue) Replace(value string) string {
	if this.patternReg != nil {
		return this.patternReg.ReplaceAllString(value, this.Replacement)
	} else {
		return strings.ReplaceAll(value, this.Pattern, this.Replacement)
	}
}

// HTTPHeaderConfig 头部信息定义
type HTTPHeaderConfig struct {
	Id    int64  `yaml:"id" json:"id"`       // ID
	IsOn  bool   `yaml:"isOn" json:"isOn"`   // 是否开启
	Name  string `yaml:"name" json:"name"`   // Name
	Value string `yaml:"value" json:"value"` // Value

	Status          *HTTPStatusConfig         `yaml:"status" json:"status"`                   // 支持的状态码
	DisableRedirect bool                      `yaml:"disableRedirect" json:"disableRedirect"` // 在跳转时不调用
	ShouldAppend    bool                      `yaml:"shouldAppend" json:"shouldAppend"`       // 是否为附加
	ShouldReplace   bool                      `yaml:"shouldReplace" json:"shouldReplace"`     // 是否替换值
	ReplaceValues   []*HTTPHeaderReplaceValue `yaml:"replaceValues" json:"replaceValues"`     // 替换值
	Methods         []string                  `yaml:"methods" json:"methods"`                 // 请求方法
	Domains         []string                  `yaml:"domains" json:"domains"`                 // 专属域名

	hasVariables bool
}

// NewHeaderConfig 获取新Header对象
func NewHeaderConfig() *HTTPHeaderConfig {
	return &HTTPHeaderConfig{
		IsOn: true,
	}
}

// Init 校验
func (this *HTTPHeaderConfig) Init() error {
	this.hasVariables = configutils.HasVariables(this.Value)

	if this.Status != nil {
		err := this.Status.Init()
		if err != nil {
			return err
		}
	}

	if this.ShouldReplace {
		for _, v := range this.ReplaceValues {
			err := v.Init()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// HasVariables 是否有变量
func (this *HTTPHeaderConfig) HasVariables() bool {
	return this.hasVariables
}
