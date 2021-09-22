package serverconfigs

// ReverseProxyRef 反向代理引用
type ReverseProxyRef struct {
	IsPrior        bool  `yaml:"isPrior" json:"isPrior"`               // 是否覆盖
	IsOn           bool  `yaml:"isOn" json:"isOn"`                     // 是否启用
	ReverseProxyId int64 `yaml:"reverseProxyId" json:"reverseProxyId"` // 反向代理ID
}

// Init 初始化
func (this *ReverseProxyRef) Init() error {
	return nil
}
