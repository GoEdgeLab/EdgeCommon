package shared

import (
	"regexp"
)

var regexpNamedVariable = regexp.MustCompile("\\${[\\w.-]+}")

// 头部信息定义
type HTTPHeaderConfig struct {
	Id     int64             `yaml:"id" json:"id"`         // ID
	IsOn   bool              `yaml:"isOn" json:"isOn"`     // 是否开启
	Name   string            `yaml:"name" json:"name"`     // Name
	Value  string            `yaml:"value" json:"value"`   // Value
	Status *HTTPStatusConfig `yaml:"status" json:"status"` // 支持的状态码 TODO

	hasVariables bool
}

// 获取新Header对象
func NewHeaderConfig() *HTTPHeaderConfig {
	return &HTTPHeaderConfig{
		IsOn: true,
	}
}

// 校验
func (this *HTTPHeaderConfig) Init() error {
	this.hasVariables = regexpNamedVariable.MatchString(this.Value)

	if this.Status != nil {
		err := this.Status.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 判断是否匹配状态码
func (this *HTTPHeaderConfig) Match(statusCode int) bool {
	if !this.IsOn {
		return false
	}

	if this.Status == nil {
		return false
	}

	return this.Status.Match(statusCode)
}

// 是否有变量
func (this *HTTPHeaderConfig) HasVariables() bool {
	return this.hasVariables
}
