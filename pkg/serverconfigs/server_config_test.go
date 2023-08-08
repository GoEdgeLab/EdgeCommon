package serverconfigs

import (
	"context"
	"testing"
)

func TestServerConfig_Protocols(t *testing.T) {
	{
		server := NewServerConfig()
		t.Log(server.FullAddresses())
	}

	{
		server := NewServerConfig()
		server.HTTP = &HTTPProtocolConfig{BaseProtocol: BaseProtocol{
			IsOn: true,
			Listen: []*NetworkAddressConfig{
				{
					Protocol:  ProtocolHTTP,
					PortRange: "1234",
				},
			},
		}}
		server.HTTPS = &HTTPSProtocolConfig{BaseProtocol: BaseProtocol{
			IsOn: true,
			Listen: []*NetworkAddressConfig{
				{
					Protocol:  ProtocolUnix,
					Host:      "/hello.sock",
					PortRange: "1235",
				},
			},
		}}
		server.TCP = &TCPProtocolConfig{BaseProtocol: BaseProtocol{
			IsOn: true,
			Listen: []*NetworkAddressConfig{
				{
					Protocol:  ProtocolHTTPS,
					PortRange: "1236",
				},
			},
		}}
		server.TLS = &TLSProtocolConfig{BaseProtocol: BaseProtocol{
			IsOn: true,
			Listen: []*NetworkAddressConfig{
				{
					Protocol:  ProtocolTCP,
					PortRange: "1234",
				},
			},
		}}
		server.Unix = &UnixProtocolConfig{BaseProtocol: BaseProtocol{
			IsOn: true,
			Listen: []*NetworkAddressConfig{
				{
					Protocol:  ProtocolTLS,
					PortRange: "1234",
				},
			},
		}}
		server.UDP = &UDPProtocolConfig{BaseProtocol: BaseProtocol{
			IsOn: true,
			Listen: []*NetworkAddressConfig{
				{
					Protocol:  ProtocolUDP,
					PortRange: "1234",
				},
			},
		}}
		err := server.Init(context.TODO())
		if err != nil {
			t.Fatal(err)
		}
		t.Log(server.FullAddresses())
	}
}

func TestServerConfig_AllStrictNames(t *testing.T) {
	var config = &ServerConfig{
		AliasServerNames: []string{"hello.com", ".hello.com"},
		ServerNames: []*ServerNameConfig{
			{
				Name: "hello2.com",
			},
			{
				SubNames: []string{"hello3.com", "hello4.com", "*.hello5.com"},
			},
			{
				Name: "~hello.com",
			},
		},
	}
	t.Log(config.AllStrictNames())
	t.Log(config.AllFuzzyNames())
}
