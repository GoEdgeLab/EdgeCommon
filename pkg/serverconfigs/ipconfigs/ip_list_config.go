package ipconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// IP名单配置
type IPListConfig struct {
	Id      int64                `yaml:"id" json:"id"`           // ID
	IsOn    bool                 `yaml:"isOn" json:"isOn"`       // 是否启用
	Version int64                `yaml:"version" json:"version"` // 版本号
	Timeout *shared.TimeDuration `yaml:"timeout" json:"timeout"` // 默认超时时间
	Code    string               `yaml:"code" json:"code"`       // 代号
	Type    string               `yaml:"type" json:"type"`       // 类型
}
