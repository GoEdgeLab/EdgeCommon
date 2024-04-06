// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package configutils_test

import (
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

func TestQuoteIP(t *testing.T) {
	t.Log(configutils.QuoteIP(configutils.QuoteIP("2001:da8:22::10")))
}
