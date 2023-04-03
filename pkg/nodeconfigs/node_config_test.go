package nodeconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/logs"
	"testing"
	"time"
)

func TestSharedNodeConfig(t *testing.T) {
	{
		config, err := SharedNodeConfig()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(config)
	}

	// read from memory cache
	{
		config, err := SharedNodeConfig()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(config)
	}
}

func TestNodeConfig_Groups(t *testing.T) {
	config := &NodeConfig{}
	config.Servers = []*serverconfigs.ServerConfig{
		{
			IsOn: true,
			HTTP: &serverconfigs.HTTPProtocolConfig{
				BaseProtocol: serverconfigs.BaseProtocol{
					IsOn: true,
					Listen: []*serverconfigs.NetworkAddressConfig{
						{
							Protocol:  serverconfigs.ProtocolHTTP,
							Host:      "127.0.0.1",
							PortRange: "1234",
						},
						{
							Protocol:  serverconfigs.ProtocolHTTP,
							PortRange: "8080",
						},
					},
				},
			},
		},
		{
			HTTP: &serverconfigs.HTTPProtocolConfig{
				BaseProtocol: serverconfigs.BaseProtocol{
					IsOn: true,
					Listen: []*serverconfigs.NetworkAddressConfig{
						{
							Protocol:  serverconfigs.ProtocolHTTP,
							PortRange: "8080",
						},
					},
				},
			},
		},
	}
	logs.PrintAsJSON(config.AvailableGroups(), t)
}

func TestCloneNodeConfig(t *testing.T) {
	var config = &NodeConfig{Id: 1, NodeId: "1", IsOn: true}
	for i := 0; i < 100_000; i++ {
		config.Servers = append(config.Servers, &serverconfigs.ServerConfig{})
	}
	var before = time.Now()
	newConfig, err := CloneNodeConfig(config)
	t.Log(time.Since(before))
	if err != nil {
		t.Fatal(err)
	}
	newConfig.Servers = []*serverconfigs.ServerConfig{}
	logs.PrintAsJSON(newConfig, t)
}

func TestNodeConfig_AddServer(t *testing.T) {
	var config = &NodeConfig{Id: 1, NodeId: "1", IsOn: true}
	config.AddServer(&serverconfigs.ServerConfig{Id: 1})
	config.AddServer(&serverconfigs.ServerConfig{Id: 2})

	t.Log("===before===")
	for _, s := range config.Servers {
		t.Log(s.Id)
	}

	t.Log("===after===")
	config.AddServer(&serverconfigs.ServerConfig{Id: 3})
	config.RemoveServer(2)
	for _, s := range config.Servers {
		t.Log(s.Id)
	}
}

func TestCloneNodeConfig_UAMPolicies(t *testing.T) {
	var config = &NodeConfig{}
	config.UAMPolicies = map[int64]*UAMPolicy{}
	t.Logf("%p", config.UAMPolicies)

	clonedConfig, err := CloneNodeConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%p", clonedConfig.UAMPolicies)
}

func BenchmarkNodeConfig(b *testing.B) {
	var config = &NodeConfig{}

	for i := 0; i < b.N; i++ {
		_, _ = CloneNodeConfig(config)
	}
}
