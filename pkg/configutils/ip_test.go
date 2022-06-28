// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package configutils_test

import (
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/assert"
	"net"
	"testing"
)

func TestParseCIDR(t *testing.T) {
	t.Log(configutils.ParseCIDR("192.168.1.1/32"))
	t.Log(configutils.ParseCIDR("192.168.1.1/24"))
	t.Log(configutils.ParseCIDR("192.168.1.1/16"))
}

func TestIPString2Long(t *testing.T) {
	for _, ip := range []string{"127.0.0.1", "192.168.1.100", "::1", "fd00:6868:6868:0:10ac:d056:3bf6:7452", "fd00:6868:6868:0:10ac:d056:3bf6:7453", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "wrong ip"} {
		t.Log(fmt.Sprintf("%42s", ip), "=>", configutils.IPString2Long(ip))
	}
}

func TestIsIPv4(t *testing.T) {
	t.Log(configutils.IsIPv4(net.ParseIP("192.168.1.100")))
	t.Log(configutils.IsIPv4(net.ParseIP("::1")))
}

func TestIsIPv6(t *testing.T) {
	t.Log(configutils.IsIPv6(net.ParseIP("192.168.1.100")))
	t.Log(configutils.IsIPv6(net.ParseIP("::1")))
}

func TestIPVersion(t *testing.T) {
	var a = assert.NewAssertion(t)
	a.IsTrue(configutils.IPVersion(net.ParseIP("192.168.1.100")) == 4)
	a.IsTrue(configutils.IPVersion(net.ParseIP("1.2.3")) == 0)
	a.IsTrue(configutils.IPVersion(net.ParseIP("::1")) == 6)
	a.IsTrue(configutils.IPVersion(net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334")) == 6)
}
