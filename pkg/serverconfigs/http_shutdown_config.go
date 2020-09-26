package serverconfigs

// 关闭页面配置
type HTTPShutdownConfig struct {
	IsPrior bool   `yaml:"isPrior" json:"isPrior"`
	IsOn    bool   `yaml:"isOn" json:"isOn"`
	URL     string `yaml:"url" json:"url"`
	Status  int    `yaml:"status" json:"status"`

	// TODO 可以自定义文本
	// TODO 可以自定义Content-Type
	// TODO 可以设置是否立即断开与客户端的连接
}

// 校验
func (this *HTTPShutdownConfig) Init() error {
	return nil
}
