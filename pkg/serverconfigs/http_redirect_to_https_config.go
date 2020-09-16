package serverconfigs

// 跳转到HTTPS配置
// TODO 支持跳转的状态码选择
type HTTPRedirectToHTTPSConfig struct {
	IsOn bool `yaml:"isOn" json:"isOn"`
}
