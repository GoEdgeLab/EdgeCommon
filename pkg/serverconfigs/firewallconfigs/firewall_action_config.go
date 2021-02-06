package firewallconfigs

import "github.com/iwind/TeaGo/maps"

// 防火墙动作配置
type FirewallActionConfig struct {
	Id         int64    `yaml:"id" json:"id"`                 // Id
	Type       string   `yaml:"type" json:"type"`             // 类型
	Params     maps.Map `yaml:"params" json:"params"`         // 参数
	EventLevel string   `yaml:"eventLevel" json:"eventLevel"` // 事件级别
}

// 初始化
func (this *FirewallActionConfig) Init() error {
	return nil
}
