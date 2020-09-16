package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type HealthCheckConfig struct {
	IsOn        bool                `yaml:"isOn" json:"isOn"`               // 是否开启 TODO
	URL         string              `yaml:"url" json:"url"`                 // TODO
	Interval    int                 `yaml:"interval" json:"interval"`       // TODO
	StatusCodes []int               `yaml:"statusCodes" json:"statusCodes"` // TODO
	Timeout     *shared.TimeDuration `yaml:"timeout" json:"timeout"`         // 超时时间 TODO
}
