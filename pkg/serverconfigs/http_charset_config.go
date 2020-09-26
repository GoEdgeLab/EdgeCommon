package serverconfigs

// 字符集设置
type HTTPCharsetConfig struct {
	IsPrior bool   `yaml:"isPrior" json:"isPrior"` // 是否覆盖
	IsOn    bool   `yaml:"isOn" json:"isOn"`       // 是否启用
	Charset string `yaml:"charset" json:"charset"` // 字符集
	IsUpper bool   `yaml:"isUpper" json:"isUpper"` // 是否要大写

	// TODO 支持自定义字符集
}

// 初始化
func (this *HTTPCharsetConfig) Init() error {
	return nil
}
