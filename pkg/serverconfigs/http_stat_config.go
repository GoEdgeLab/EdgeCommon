package serverconfigs

type HTTPStatRef struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"` // 是否覆盖
	IsOn    bool `yaml:"isOn" json:"isOn"`       // 是否开启
}

func (this *HTTPStatRef) Init() error {
	return nil
}
