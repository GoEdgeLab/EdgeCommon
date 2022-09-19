// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

import (
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/ddosconfigs"
)

type NSNodeConfig struct {
	Id              int64                         `yaml:"id" json:"id"`
	NodeId          string                        `yaml:"nodeId" json:"nodeId"`
	Secret          string                        `yaml:"secret" json:"secret"`
	ClusterId       int64                         `yaml:"clusterId" json:"clusterId"`
	AccessLogRef    *NSAccessLogRef               `yaml:"accessLogRef" json:"accessLogRef"`
	RecursionConfig *NSRecursionConfig            `yaml:"recursionConfig" json:"recursionConfig"`
	DDoSProtection  *ddosconfigs.ProtectionConfig `yaml:"ddosProtection" json:"ddosProtection"`
	AllowedIPs      []string                      `yaml:"allowedIPs" json:"allowedIPs"`
	TimeZone        string                        `yaml:"timeZone" json:"timeZone"` // 自动设置时区

	TCP *serverconfigs.TCPProtocolConfig `yaml:"tcp" json:"tcp"` // TCP配置
	TLS *serverconfigs.TLSProtocolConfig `yaml:"tls" json:"tls"` // TLS配置
	UDP *serverconfigs.UDPProtocolConfig `yaml:"udp" json:"udp"` // UDP配置

	paddedId string
}

func (this *NSNodeConfig) Init() error {
	this.paddedId = fmt.Sprintf("%08d", this.Id)

	// accessLog
	if this.AccessLogRef != nil {
		err := this.AccessLogRef.Init()
		if err != nil {
			return err
		}
	}

	// 递归DNS
	if this.RecursionConfig != nil {
		err := this.RecursionConfig.Init()
		if err != nil {
			return err
		}
	}

	// DDoS
	if this.DDoSProtection != nil {
		err := this.DDoSProtection.Init()
		if err != nil {
			return err
		}
	}

	// tcp
	if this.TCP != nil {
		err := this.TCP.Init()
		if err != nil {
			return err
		}
	}

	// tls
	if this.TLS != nil {
		err := this.TLS.Init()
		if err != nil {
			return err
		}
	}

	// udp
	if this.UDP != nil {
		err := this.UDP.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *NSNodeConfig) PaddedId() string {
	return this.paddedId
}
