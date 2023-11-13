package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type HTTPPageBodyType = string

const (
	HTTPPageBodyTypeHTML        HTTPPageBodyType = "html"
	HTTPPageBodyTypeURL         HTTPPageBodyType = "url"
	HTTPPageBodyTypeRedirectURL HTTPPageBodyType = "redirectURL"
)

func FindAllHTTPPageBodyTypes() []*shared.Definition {
	return []*shared.Definition{
		{
			Name: "HTML",
			Code: HTTPPageBodyTypeHTML,
		},
		{
			Name: "读取URL",
			Code: HTTPPageBodyTypeURL,
		},
		{
			Name: "跳转URL",
			Code: HTTPPageBodyTypeRedirectURL,
		},
	}
}

// HTTPPageConfig 特殊页面配置
type HTTPPageConfig struct {
	Id        int64    `yaml:"id" json:"id"`               // 页面ID
	IsOn      bool     `yaml:"isOn" json:"isOn"`           // 是否开启 TODO
	Status    []string `yaml:"status" json:"status"`       // 响应支持40x, 50x, 3x2
	NewStatus int      `yaml:"newStatus" json:"newStatus"` // 新状态码

	BodyType HTTPPageBodyType `yaml:"bodyType" json:"bodyType"` // 内容类型
	URL      string           `yaml:"url" json:"url"`           // URL
	Body     string           `yaml:"body" json:"body"`         // 输出的内容

	OnlyURLPatterns   []*shared.URLPattern `yaml:"onlyURLPatterns" json:"onlyURLPatterns"`     // 仅限的URL
	ExceptURLPatterns []*shared.URLPattern `yaml:"exceptURLPatterns" json:"exceptURLPatterns"` // 排除的URL

	statusList    []*WildcardStatus
	hasStatusList bool
}

// NewHTTPPageConfig 获取新对象
func NewHTTPPageConfig() *HTTPPageConfig {
	return &HTTPPageConfig{
		IsOn: true,
	}
}

// Init 校验
func (this *HTTPPageConfig) Init() error {
	this.statusList = []*WildcardStatus{}
	for _, s := range this.Status {
		this.statusList = append(this.statusList, NewWildcardStatus(s))
	}
	this.hasStatusList = len(this.statusList) > 0

	for _, urlPattern := range this.OnlyURLPatterns {
		err := urlPattern.Init()
		if err != nil {
			return err
		}
	}

	for _, urlPattern := range this.ExceptURLPatterns {
		err := urlPattern.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// Match 检查是否匹配
func (this *HTTPPageConfig) Match(status int) bool {
	if !this.hasStatusList {
		return false
	}
	for _, s := range this.statusList {
		if s.Match(status) {
			return true
		}
	}
	return false
}

func (this *HTTPPageConfig) MatchURL(url string) bool {
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
