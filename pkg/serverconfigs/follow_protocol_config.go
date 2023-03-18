// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

// FollowProtocolConfig 协议跟随配置
type FollowProtocolConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"` // 是否覆盖父级配置
	IsOn    bool `yaml:"isOn" json:"isOn"`       // 是否启用
	HTTP    struct {
		Port       int  `yaml:"port" json:"port"`             // 端口
		FollowPort bool `yaml:"followPort" json:"followPort"` // 跟随端口
	} `yaml:"http" json:"http"` // HTTP配置
	HTTPS struct {
		Port       int  `yaml:"port" json:"port"`             // 端口
		FollowPort bool `yaml:"followPort" json:"followPort"` // 跟随端口
	} `yaml:"https" json:"https"` // HTTPS配置
}

func NewFollowProtocolConfig() *FollowProtocolConfig {
	var p = &FollowProtocolConfig{}
	p.HTTP.FollowPort = true
	p.HTTPS.FollowPort = true
	return p
}

func (this *FollowProtocolConfig) Init() error {
	return nil
}
