package serverconfigs

// 反向代理引用
type ReverseProxyRef struct {
	IsOn           bool  `yaml:"isOn" json:"isOn"`                     // 是否启用
	ReverseProxyId int64 `yaml:"reverseProxyId" json:"reverseProxyId"` // 反向代理ID
}
