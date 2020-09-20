package serverconfigs

type HTTPStatConfig struct {
	IsOn bool `yaml:"isOn" json:"isOn"` // 是否开启
}

func (this *HTTPStatConfig) Init() error {
	return nil
}
