package serverconfigs

type HTTPStatRef struct {
	IsOn bool `yaml:"isOn" json:"isOn"` // 是否开启
}

func (this *HTTPStatRef) Init() error {
	return nil
}
