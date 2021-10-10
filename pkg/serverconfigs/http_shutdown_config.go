package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// HTTPShutdownConfig 关闭页面配置
type HTTPShutdownConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`
	IsOn    bool `yaml:"isOn" json:"isOn"`

	BodyType shared.BodyType `yaml:"bodyType" json:"bodyType"` // 内容类型
	URL      string          `yaml:"url" json:"url"`           // URL
	Body     string          `yaml:"body" json:"body"`         // 输出的内容

	Status int `yaml:"status" json:"status"`

	// TODO 可以自定义Content-Type
	// TODO 可以设置是否立即断开与客户端的连接
}

// Init 校验
func (this *HTTPShutdownConfig) Init() error {
	return nil
}
