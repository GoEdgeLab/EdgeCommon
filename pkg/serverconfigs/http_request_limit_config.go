// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// HTTPRequestLimitConfig HTTP请求限制相关限制配置
type HTTPRequestLimitConfig struct {
	IsPrior             bool                 `yaml:"isPrior" json:"isPrior"`                         // 是否覆盖父级
	IsOn                bool                 `yaml:"isOn" json:"isOn"`                               // 是否启用
	MaxConns            int                  `yaml:"maxConns" json:"maxConns"`                       // 并发连接数
	MaxConnsPerIP       int                  `yaml:"maxConnsPerIP" json:"maxConnsPerIP"`             // 单个IP并发连接数
	OutBandwidthPerConn *shared.SizeCapacity `yaml:"outBandwidthPerConn" json:"outBandwidthPerConn"` // 下行流量限制
	MaxBodySize         *shared.SizeCapacity `yaml:"maxBodySize" json:"maxBodySize"`                 // 单个请求最大尺寸

	outBandwidthPerConnBytes int64
	maxBodyBytes             int64
}

func (this *HTTPRequestLimitConfig) Init() error {
	if this.OutBandwidthPerConn != nil {
		this.outBandwidthPerConnBytes = this.OutBandwidthPerConn.Bytes()
	}
	if this.MaxBodySize != nil {
		this.maxBodyBytes = this.MaxBodySize.Bytes()
	}

	return nil
}

func (this *HTTPRequestLimitConfig) OutBandwidthPerConnBytes() int64 {
	return this.outBandwidthPerConnBytes
}

func (this *HTTPRequestLimitConfig) MaxBodyBytes() int64 {
	return this.maxBodyBytes
}
