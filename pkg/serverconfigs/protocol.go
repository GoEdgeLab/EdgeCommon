package serverconfigs

type Protocol string

const (
	ProtocolHTTP  Protocol = "http"
	ProtocolHTTPS Protocol = "https"
	ProtocolTCP   Protocol = "tcp"
	ProtocolTLS   Protocol = "tls"
	ProtocolUnix  Protocol = "unix"
	ProtocolUDP   Protocol = "udp"

	// 子协议
	ProtocolHTTP4 Protocol = "http4"
	ProtocolHTTP6 Protocol = "http6"

	ProtocolHTTPS4 Protocol = "https4"
	ProtocolHTTPS6 Protocol = "https6"

	ProtocolTCP4 Protocol = "tcp4"
	ProtocolTCP6 Protocol = "tcp6"

	ProtocolTLS4 Protocol = "tls4"
	ProtocolTLS6 Protocol = "tls6"
)

func AllProtocols() []Protocol {
	return []Protocol{ProtocolHTTP, ProtocolHTTPS, ProtocolTCP, ProtocolTLS, ProtocolUnix, ProtocolUDP, ProtocolHTTP4, ProtocolHTTP6, ProtocolHTTPS4, ProtocolHTTPS6, ProtocolTCP4, ProtocolTCP6, ProtocolTLS4, ProtocolTLS6}
}

func (this Protocol) IsHTTPFamily() bool {
	return this == ProtocolHTTP || this == ProtocolHTTP4 || this == ProtocolHTTP6
}

func (this Protocol) IsHTTPSFamily() bool {
	return this == ProtocolHTTPS || this == ProtocolHTTPS4 || this == ProtocolHTTPS6
}

func (this Protocol) IsTCPFamily() bool {
	return this == ProtocolTCP || this == ProtocolTCP4 || this == ProtocolTCP6
}

func (this Protocol) IsTLSFamily() bool {
	return this == ProtocolTLS || this == ProtocolTLS4 || this == ProtocolTLS6
}

func (this Protocol) IsUnixFamily() bool {
	return this == ProtocolUnix
}

func (this Protocol) IsUDPFamily() bool {
	return this == ProtocolUDP
}

// 主协议
func (this Protocol) Primary() Protocol {
	switch this {
	case ProtocolHTTP, ProtocolHTTP4, ProtocolHTTP6:
		return ProtocolHTTP
	case ProtocolHTTPS, ProtocolHTTPS4, ProtocolHTTPS6:
		return ProtocolHTTPS
	case ProtocolTCP, ProtocolTCP4, ProtocolTCP6:
		return ProtocolTCP
	case ProtocolTLS, ProtocolTLS4, ProtocolTLS6:
		return ProtocolTLS
	case ProtocolUnix:
		return ProtocolUnix
	case ProtocolUDP:
		return ProtocolUDP
	default:
		return this
	}
}

// Scheme
func (this Protocol) Scheme() string {
	return string(this)
}

// 转换为字符串
func (this Protocol) String() string {
	return string(this)
}
