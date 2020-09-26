package serverconfigs

type HTTPCharsetConfig struct {
	IsPrior bool   `yaml:"isPrior" json:"isPrior"` // 是否覆盖
	IsOn    bool   `yaml:"isOn" json:"isOn"`       // 是否启用
	Charset string `yaml:"charset" json:"charset"` // 字符集
}

func (this *HTTPCharsetConfig) Init() error {
	return nil
}
