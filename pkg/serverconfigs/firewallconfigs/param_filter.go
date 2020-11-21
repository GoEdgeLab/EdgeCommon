package firewallconfigs

import "github.com/iwind/TeaGo/maps"

// 对参数的过滤器
type ParamFilter struct {
	Code    string   `yaml:"code" json:"code"`       // 过滤器编号
	Name    string   `yaml:"name" json:"name"`       // 名称
	Options maps.Map `yaml:"options" json:"options"` // 过滤器选项
}
