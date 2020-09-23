package serverconfigs

// 跳转到HTTPS配置
// TODO 支持跳转的状态码选择
type HTTPRedirectToHTTPSConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"` // 是否覆盖
	IsOn    bool `yaml:"isOn" json:"isOn"`       // 是否开启
}
