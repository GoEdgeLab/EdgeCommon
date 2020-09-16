package serverconfigs

// 关闭页面配置
type HTTPShutdownConfig struct {
	IsOn   bool   `yaml:"isOn" json:"isOn"`
	URL    string `yaml:"url" json:"url"`
	Status int    `yaml:"status" json:"status"`
}

// 获取新对象
func NewHTTPShutdownConfig() *HTTPShutdownConfig {
	return &HTTPShutdownConfig{}
}

// 校验
func (this *HTTPShutdownConfig) Init() error {
	return nil
}
