package serverconfigs

import (
	"github.com/iwind/TeaGo/logs"
	"testing"
)

func TestNormalizeServerNames(t *testing.T) {
	serverNames := []*ServerNameConfig{
		{
			Name:     "Hello.com",
			SubNames: []string{"WoRld.com", "XYZ.com"},
		},
	}
	NormalizeServerNames(serverNames)
	logs.PrintAsJSON(serverNames, t)
}

func TestPlainServerNames(t *testing.T) {
	serverNames := []*ServerNameConfig{
		{
			Name:     "Hello.com",
			SubNames: nil,
		},
		{
			Name:     "world.com",
			SubNames: nil,
		},
		{
			Name:     "",
			SubNames: []string{"WoRld.com", "XYZ.com"},
		},
	}
	logs.PrintAsJSON(PlainServerNames(serverNames), t)
}
