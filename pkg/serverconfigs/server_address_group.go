package serverconfigs

import "strings"

type ServerAddressGroup struct {
	fullAddr string
	Servers  []*ServerConfig
}

func NewServerAddressGroup(fullAddr string) *ServerAddressGroup {
	return &ServerAddressGroup{fullAddr: fullAddr}
}

// Add 添加服务
func (this *ServerAddressGroup) Add(server *ServerConfig) {
	this.Servers = append(this.Servers, server)
}

// FullAddr 获取完整的地址
func (this *ServerAddressGroup) FullAddr() string {
	return this.fullAddr
}

// Protocol 获取当前分组的协议
func (this *ServerAddressGroup) Protocol() Protocol {
	for _, p := range AllProtocols() {
		if strings.HasPrefix(this.fullAddr, p.String()+":") {
			return p
		}
	}
	return ProtocolHTTP
}

// Addr 获取当前分组的地址
func (this *ServerAddressGroup) Addr() string {
	protocol := this.Protocol()
	if protocol == ProtocolUnix {
		return strings.TrimPrefix(this.fullAddr, protocol.String()+":")
	}
	return strings.TrimPrefix(this.fullAddr, protocol.String()+"://")
}

// IsHTTP 判断当前分组是否为HTTP
func (this *ServerAddressGroup) IsHTTP() bool {
	p := this.Protocol()
	return p == ProtocolHTTP || p == ProtocolHTTP4 || p == ProtocolHTTP6
}

// IsHTTPS 判断当前分组是否为HTTPS
func (this *ServerAddressGroup) IsHTTPS() bool {
	p := this.Protocol()
	return p == ProtocolHTTPS || p == ProtocolHTTPS4 || p == ProtocolHTTPS6
}

// IsTCP 判断当前分组是否为TCP
func (this *ServerAddressGroup) IsTCP() bool {
	p := this.Protocol()
	return p == ProtocolTCP || p == ProtocolTCP4 || p == ProtocolTCP6
}

// IsTLS 判断当前分组是否为TLS
func (this *ServerAddressGroup) IsTLS() bool {
	p := this.Protocol()
	return p == ProtocolTLS || p == ProtocolTLS4 || p == ProtocolTLS6
}

// IsUnix 判断当前分组是否为Unix
func (this *ServerAddressGroup) IsUnix() bool {
	p := this.Protocol()
	return p == ProtocolUnix
}

// IsUDP 判断当前分组是否为UDP
func (this *ServerAddressGroup) IsUDP() bool {
	p := this.Protocol()
	return p == ProtocolUDP
}

// FirstServer 获取第一个Server
func (this *ServerAddressGroup) FirstServer() *ServerConfig {
	if len(this.Servers) > 0 {
		return this.Servers[0]
	}
	return nil
}
