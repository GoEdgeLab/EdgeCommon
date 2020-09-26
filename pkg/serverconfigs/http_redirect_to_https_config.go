package serverconfigs

// 跳转到HTTPS配置
type HTTPRedirectToHTTPSConfig struct {
	IsPrior bool   `yaml:"isPrior" json:"isPrior"` // 是否覆盖
	IsOn    bool   `yaml:"isOn" json:"isOn"`       // 是否开启
	Status  int    `yaml:"status" json:"status"`   // 跳转用的状态码
	Host    string `yaml:"host" json:"host"`       // 跳转后的Host
	Port    int    `yaml:"port" json:"port"`       // 跳转后的端口
}

func (this *HTTPRedirectToHTTPSConfig) Init() error {
	return nil
}
