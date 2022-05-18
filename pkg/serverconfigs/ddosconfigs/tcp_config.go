// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ddosconfigs

type TCPConfig struct {
	IsPrior             bool          `json:"isPrior"`
	IsOn                bool          `json:"isOn"`
	MaxConnections      int32         `json:"maxConnections"`
	MaxConnectionsPerIP int32         `json:"maxConnectionsPerIP"`
	NewConnectionsRate  int32         `json:"newConnectionsRate"`
	AllowIPList         []*IPConfig   `json:"allowIPList"`
	Ports               []*PortConfig `json:"ports"`
}

func (this *TCPConfig) Init() error {
	return nil
}
