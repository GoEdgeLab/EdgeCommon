package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// HTTPPageConfig 特殊页面配置
// TODO 需要支持Header定义
type HTTPPageConfig struct {
	Id        int64    `yaml:"id" json:"id"`               // 页面ID
	IsOn      bool     `yaml:"isOn" json:"isOn"`           // 是否开启 TODO
	Status    []string `yaml:"status" json:"status"`       // 响应支持40x, 50x, 3x2
	NewStatus int      `yaml:"newStatus" json:"newStatus"` // 新状态码

	BodyType shared.BodyType `yaml:"bodyType" json:"bodyType"` // 内容类型
	URL      string          `yaml:"url" json:"url"`           // URL
	Body     string          `yaml:"body" json:"body"`         // 输出的内容

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
