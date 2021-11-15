package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"github.com/iwind/TeaGo/types"
	"testing"
	"time"
)

func TestServerAddressGroup_Protocol(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		group := NewServerAddressGroup("tcp://127.0.0.1:1234")
		a.IsTrue(group.Protocol() == ProtocolTCP)
		a.IsTrue(group.Addr() == "127.0.0.1:1234")
	}

	{
		group := NewServerAddressGroup("http4://127.0.0.1:1234")
		a.IsTrue(group.Protocol() == ProtocolHTTP4)
		a.IsTrue(group.Addr() == "127.0.0.1:1234")
	}

	{
		group := NewServerAddressGroup("127.0.0.1:1234")
		a.IsTrue(group.Protocol() == ProtocolHTTP)
		a.IsTrue(group.Addr() == "127.0.0.1:1234")
	}

	{
		group := NewServerAddressGroup("unix:/tmp/my.sock")
		a.IsTrue(group.Protocol() == ProtocolUnix)
		a.IsTrue(group.Addr() == "/tmp/my.sock")
	}
}

func TestServerAddressGroup_MatchServerName(t *testing.T) {
	var group = NewServerAddressGroup("")
	for i := 0; i < 1_000_000; i++ {
		group.Add(&ServerConfig{
			ServerNames: []*ServerNameConfig{
				{
					Name:     "hello" + types.String(i) + ".com",
					SubNames: []string{},
				},
			},
		})
	}
	group.Add(&ServerConfig{
		ServerNames: []*ServerNameConfig{
			{
				Name:     "hello.com",
				SubNames: []string{},
			},
		},
	})
	group.Add(&ServerConfig{
		ServerNames: []*ServerNameConfig{
			{
				Name:     "*.hello.com",
				SubNames: []string{},
			},
		},
	})

	var before = time.Now()
	defer func() {
		t.Log(time.Since(before).Seconds()*1000, "ms")
	}()

	t.Log(group.MatchServerName("hello99999.com").AllStrictNames())
	t.Log(group.MatchServerName("hello.com").AllStrictNames())
	t.Log(group.MatchServerName("world.hello.com").AllFuzzyNames())
	for i := 0; i < 100_000; i++ {
		_ = group.MatchServerName("world.hello.com")
	}
}

func TestServerAddressGroup_MatchServerCNAME(t *testing.T) {
	var group = NewServerAddressGroup("")
	group.Add(&ServerConfig{
		ServerNames: []*ServerNameConfig{
			{
				Name:     "hello.com",
				SubNames: []string{},
			},
		},
		SupportCNAME: true,
	})
	group.Add(&ServerConfig{
		ServerNames: []*ServerNameConfig{
			{
				Name:     "*.hello.com",
				SubNames: []string{},
			},
		},
	})

	var before = time.Now()
	defer func() {
		t.Log(time.Since(before).Seconds()*1000, "ms")
	}()

	server := group.MatchServerCNAME("hello.com")
	if server != nil {
		t.Log(server.AllStrictNames())
	} else {
		t.Log(server)
	}
	t.Log(group.MatchServerCNAME("world.hello.com"))
}
