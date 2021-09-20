package serverconfigs

import "github.com/iwind/TeaGo/maps"

// SchedulingConfig 调度算法配置
type SchedulingConfig struct {
	Code    string   `yaml:"code" json:"code"`       // 类型
	Options maps.Map `yaml:"options" json:"options"` // 选项
}

// NewSchedulingConfig 获取新对象
func NewSchedulingConfig() *SchedulingConfig {
	return &SchedulingConfig{}
}

// Clone 克隆
func (this *SchedulingConfig) Clone() *SchedulingConfig {
	return &SchedulingConfig{
		Code:    this.Code,
		Options: maps.NewMap(this.Options),
	}
}
