// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package firewallconfigs

// SYNFloodConfig Syn flood防护设置
type SYNFloodConfig struct {
	IsPrior        bool  `yaml:"isPrior" json:"isPrior"`
	IsOn           bool  `yaml:"isOn" json:"isOn"`
	MinAttempts    int32 `yaml:"minAttempts" json:"minAttempts"`       // 最小尝试次数/分钟
	TimeoutSeconds int32 `yaml:"timeoutSeconds" json:"timeoutSeconds"` // 拦截超时时间
	IgnoreLocal    bool  `yaml:"ignoreLocal" json:"ignoreLocal"`       // 忽略本地IP
}

func DefaultSYNFloodConfig() *SYNFloodConfig {
	return &SYNFloodConfig{
		IsOn:           true,
		MinAttempts:    10,
		TimeoutSeconds: 600,
		IgnoreLocal:    true,
	}
}

func (this *SYNFloodConfig) Init() error {
	return nil
}
