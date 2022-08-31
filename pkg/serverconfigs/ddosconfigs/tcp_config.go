// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ddosconfigs

type TCPConfig struct {
	IsPrior             bool  `json:"isPrior"`
	IsOn                bool  `json:"isOn"`
	MaxConnections      int32 `json:"maxConnections"`
	MaxConnectionsPerIP int32 `json:"maxConnectionsPerIP"`

	// 分钟级速率
	NewConnectionsMinutelyRate             int32 `json:"newConnectionsRate"`             // 分钟
	NewConnectionsMinutelyRateBlockTimeout int32 `json:"newConnectionsRateBlockTimeout"` // 拦截时间

	// 秒级速率
	NewConnectionsSecondlyRate             int32 `json:"newConnectionsSecondlyRate"`
	NewConnectionsSecondlyRateBlockTimeout int32 `json:"newConnectionsSecondlyRateBlockTimeout"`

	AllowIPList []*IPConfig   `json:"allowIPList"`
	Ports       []*PortConfig `json:"ports"`
}

func (this *TCPConfig) Init() error {
	return nil
}
