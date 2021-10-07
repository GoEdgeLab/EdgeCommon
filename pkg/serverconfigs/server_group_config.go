// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

// ServerGroupConfig 服务分组配置
type ServerGroupConfig struct {
	Id   int64  `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
	IsOn bool   `yaml:"isOn" json:"isOn"`

	// 反向代理配置
	HTTPReverseProxyRef *ReverseProxyRef    `yaml:"httpReverseProxyRef" json:"httpReverseProxyRef"`
	HTTPReverseProxy    *ReverseProxyConfig `yaml:"httpReverseProxy" json:"httpReverseProxy"`
	TCPReverseProxyRef  *ReverseProxyRef    `yaml:"tcpReverseProxyRef" json:"tcpReverseProxyRef"`
	TCPReverseProxy     *ReverseProxyConfig `yaml:"tcpReverseProxy" json:"tcpReverseProxy"`
	UDPReverseProxyRef  *ReverseProxyRef    `yaml:"udpReverseProxyRef" json:"udpReverseProxyRef"`
	UDPReverseProxy     *ReverseProxyConfig `yaml:"udpReverseProxy" json:"udpReverseProxy"`

	Web *HTTPWebConfig `yaml:"web" json:"web"`
}
