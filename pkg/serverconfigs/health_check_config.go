package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
)

// 健康检查设置
type HealthCheckConfig struct {
	IsOn           bool                 `yaml:"isOn" json:"isOn"`                     // 是否开启
	URL            string               `yaml:"url" json:"url"`                       // 读取的URL
	Interval       *shared.TimeDuration `yaml:"interval" json:"interval"`             // 检测周期
	StatusCodes    []int                `yaml:"statusCodes" json:"statusCodes"`       // 返回的状态码要求
	Timeout        *shared.TimeDuration `yaml:"timeout" json:"timeout"`               // 超时时间
	CountTries     int64                `yaml:"countTries" json:"countTries"`         // 尝试次数
	TryDelay       *shared.TimeDuration `yaml:"tryDelay" json:"tryDelay"`             // 尝试间隔
	FailActions    []maps.Map           `yaml:"failActions" json:"failActions"`       // 失败采取的动作 TODO
	RecoverActions []maps.Map           `yaml:"recoverActions" json:"recoverActions"` // 恢复采取的动作 TODO
	AutoDown       bool                 `yaml:"autoDown" json:"autoDown"`             // 是否自动下线
	CountUp        int                  `yaml:"countUp" json:"countUp"`               // 连续在线认定次数
	CountDown      int                  `yaml:"countDown" json:"countDown"`           // 连续离线认定次数
}

// 初始化
func (this *HealthCheckConfig) Init() error {
	return nil
}
