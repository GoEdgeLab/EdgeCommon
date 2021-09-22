package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
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
