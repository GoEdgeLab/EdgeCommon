// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodeconfigs

func DefaultClockConfig() *ClockConfig {
	return &ClockConfig{
		AutoSync:    true,
		Server:      "",
		CheckChrony: true,
	}
}

// ClockConfig 时钟相关配置
type ClockConfig struct {
	AutoSync    bool   `yaml:"autoSync" json:"autoSync"`       // 自动尝试同步时钟
	Server      string `yaml:"server" json:"server"`           // 时钟同步服务器
	CheckChrony bool   `yaml:"checkChrony" json:"checkChrony"` // 检查 chronyd 是否在运行
}

func (this *ClockConfig) Init() error {
	return nil
}
