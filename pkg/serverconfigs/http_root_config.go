package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
)

func NewHTTPRootConfig() *HTTPRootConfig {
	return &HTTPRootConfig{
		ExceptHiddenFiles: true,
	}
}

// HTTPRootConfig Web文档目录配置
type HTTPRootConfig struct {
	IsPrior           bool                 `yaml:"isPrior" json:"isPrior"`                     // 是否优先
	IsOn              bool                 `yaml:"isOn" json:"isOn"`                           // 是否启用
	Dir               string               `yaml:"dir" json:"dir"`                             // 目录
	Indexes           []string             `yaml:"indexes" json:"indexes"`                     // 默认首页文件
	StripPrefix       string               `yaml:"stripPrefix" json:"stripPrefix"`             // 去除URL前缀
	DecodePath        bool                 `yaml:"decodePath" json:"decodePath"`               // 是否对请求路径进行解码
	IsBreak           bool                 `yaml:"isBreak" json:"isBreak"`                     // 找不到文件的情况下是否终止
	ExceptHiddenFiles bool                 `yaml:"exceptHiddenFiles" json:"exceptHiddenFiles"` // 排除隐藏文件
	OnlyURLPatterns   []*shared.URLPattern `yaml:"onlyURLPatterns" json:"onlyURLPatterns"`     // 仅限的URL
	ExceptURLPatterns []*shared.URLPattern `yaml:"exceptURLPatterns" json:"exceptURLPatterns"` // 排除的URL

	hasVariables bool
}

// Init 初始化
func (this *HTTPRootConfig) Init() error {
	this.hasVariables = configutils.HasVariables(this.Dir)

	for _, pattern := range this.OnlyURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	for _, pattern := range this.ExceptURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// HasVariables 判断是否有变量
func (this *HTTPRootConfig) HasVariables() bool {
	return this.hasVariables
}

func (this *HTTPRootConfig) MatchURL(url string) bool {
	// except
	if len(this.ExceptURLPatterns) > 0 {
		for _, pattern := range this.ExceptURLPatterns {
			if pattern.Match(url) {
				return false
			}
		}
	}

	if len(this.OnlyURLPatterns) > 0 {
		for _, pattern := range this.OnlyURLPatterns {
			if pattern.Match(url) {
				return true
			}
		}
		return false
	}

	return true
}
