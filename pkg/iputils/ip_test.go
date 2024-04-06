// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iputils_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/iputils"
	"github.com/iwind/TeaGo/assert"
	"runtime"
	"testing"
)

func TestIP_ParseIP(t *testing.T) {
	var a = assert.NewAssertion(t)

	{
		var i = iputils.ParseIP("127.0.0.1")
		a.IsTrue(i.IsIPv4())
		a.IsFalse(i.IsIPv6())
		a.IsTrue(i.IsValid())
		a.IsTrue(iputils.IsIPv4("127.0.0.1"))
		a.IsFalse(iputils.IsIPv6("127.0.0.1"))
		t.Log(i.String(), i.ToLong())
		t.Log("raw:", i.Raw())
	}

	{
		var i = iputils.ParseIP("0.0.0.1")
		a.IsTrue(i.IsIPv4())
		a.IsFalse(i.IsIPv6())
		t.Log(i.String(), i.ToLong())
	}

	for j := 0; j < 3; j++ /** repeat test **/ {
		var i = iputils.ParseIP("::1")
		a.IsFalse(i.IsIPv4())
		a.IsTrue(i.IsIPv6())
		a.IsTrue(i.IsValid())
		t.Log(i.String(), i.ToLong())
	}

	{
		{
			var i = iputils.ParseIP("2001:db8:0:1::1:101")
			t.Log(i.String(), i.ToLong())
			a.IsFalse(i.IsIPv4())
			a.IsTrue(i.IsIPv6())
			a.IsFalse(iputils.IsIPv4("2001:db8:0:1::1:101"))
			a.IsTrue(iputils.IsIPv6("2001:db8:0:1::1:101"))
			a.IsTrue(i.IsValid())
		}

		{
			var i = iputils.ParseIP("2001:db8:0:1::1:102")
			t.Log(i.String(), i.ToLong())
			a.IsFalse(i.IsIPv4())
			a.IsTrue(i.IsIPv6())
			a.IsTrue(i.IsValid())
		}

		{
			var i = iputils.ParseIP("2001:db8:0:1::2:101")
			t.Log(i.String(), i.ToLong())
			a.IsFalse(i.IsIPv4())
			a.IsTrue(i.IsIPv6())
			a.IsTrue(i.IsValid())
		}
	}

	{
		var i = iputils.ParseIP("WRONG IP")
		t.Log(i.String(), i.ToLong())
		a.IsFalse(i.IsIPv4())
		a.IsFalse(i.IsIPv6())
		a.IsFalse(i.IsValid())
		a.IsFalse(iputils.IsIPv4("WRONG IP"))
		a.IsFalse(iputils.IsIPv6("WRONG IP"))
	}
}

func TestIP_Mod(t *testing.T) {
	for _, ip := range []string{
		"127.0.0.1",
		"::1",
		"2001:db8:0:1::1:101",
		"2001:db8:0:1::1:102",
		"WRONG IP",
	} {
		var i = iputils.ParseIP(ip)
		t.Log(ip, "=>", i.ToLong(), "=>", i.Mod(5))
	}
}

func TestIP_Compare(t *testing.T) {
	var a = assert.NewAssertion(t)

	{
		var i1 = iputils.ParseIP("127.0.0.1")
		var i2 = iputils.ParseIP("127.0.0.1")
		a.IsTrue(i1.Compare(i2) == 0)
	}

	{
		var i1 = iputils.ParseIP("127.0.0.1")
		var i2 = iputils.ParseIP("127.0.0.2")
		a.IsTrue(i1.Compare(i2) == -1)
	}

	{
		var i1 = iputils.ParseIP("127.0.0.2")
		var i2 = iputils.ParseIP("127.0.0.1")
		a.IsTrue(i1.Compare(i2) == 1)
	}

	{
		var i1 = iputils.ParseIP("2001:db8:0:1::101")
		var i2 = iputils.ParseIP("127.0.0.1")
		a.IsTrue(i1.Compare(i2) == 1)
	}

	{
		var i1 = iputils.ParseIP("127.0.0.1")
		var i2 = iputils.ParseIP("2001:db8:0:1::101")
		a.IsTrue(i1.Compare(i2) == -1)
	}

	{
		var i1 = iputils.ParseIP("2001:db8:0:1::101")
		var i2 = iputils.ParseIP("2001:db8:0:1::101")
		a.IsTrue(i1.Compare(i2) == 0)
	}

	{
		var i1 = iputils.ParseIP("2001:db8:0:1::101")
		var i2 = iputils.ParseIP("2001:db8:0:1::102")
		a.IsTrue(i1.Compare(i2) == -1)
	}

	{
		var i1 = iputils.ParseIP("2001:db8:0:1::102")
		var i2 = iputils.ParseIP("2001:db8:0:1::101")
		a.IsTrue(i1.Compare(i2) == 1)
	}

	{
		var i1 = iputils.ParseIP("2001:db8:0:1::2:100")
		var i2 = iputils.ParseIP("2001:db8:0:1::1:101")
		a.IsTrue(i1.Compare(i2) == 1)
	}
}

func TestIP_Between(t *testing.T) {
	var a = assert.NewAssertion(t)
	a.IsTrue(iputils.ParseIP("127.0.0.2").Between(iputils.ParseIP("127.0.0.1"), iputils.ParseIP("127.0.0.3")))
	a.IsTrue(iputils.ParseIP("127.0.0.1").Between(iputils.ParseIP("127.0.0.1"), iputils.ParseIP("127.0.0.3")))
	a.IsFalse(iputils.ParseIP("127.0.0.2").Between(iputils.ParseIP("127.0.0.3"), iputils.ParseIP("127.0.0.4")))
	a.IsFalse(iputils.ParseIP("127.0.0.5").Between(iputils.ParseIP("127.0.0.3"), iputils.ParseIP("127.0.0.4")))
	a.IsFalse(iputils.ParseIP("127.0.0.2").Between(iputils.ParseIP("127.0.0.3"), iputils.ParseIP("127.0.0.1")))
}

func TestIP_ToLong(t *testing.T) {
	for _, ip := range []string{
		"127.0.0.1",
		"192.168.1.100",
		"::1",
		"fd00:6868:6868:0:10ac:d056:3bf6:7452",
		"fd00:6868:6868:0:10ac:d056:3bf6:7453",
		"2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		"2001:db8:0:1::101",
		"2001:db8:0:2::101",
		"wrong ip",
	} {
		var goIP = iputils.ParseIP(ip)
		t.Log(ip, "=>", "\n", goIP.String(), "\n", "=>", "\n", "long1:", goIP.ToLong(), "\n", "long2:", iputils.ToLong(ip), "\n", "little long:", iputils.ToLittleLong(ip))
	}
}

func TestIP_CompareLong(t *testing.T) {
	var a = assert.NewAssertion(t)
	a.IsTrue(iputils.CompareLong("1", "2") == -1)
	a.IsTrue(iputils.CompareLong("11", "2") == 1)
	a.IsTrue(iputils.CompareLong("11", "22") == -1)
	a.IsTrue(iputils.CompareLong("22", "101") == -1)
	a.IsTrue(iputils.CompareLong("33", "22") == 1)
	a.IsTrue(iputils.CompareLong("101", "22") == 1)
	a.IsTrue(iputils.CompareLong("22", "22") == 0)
}

func TestIP_Memory(t *testing.T) {
	var list []iputils.IP

	var stat1 = &runtime.MemStats{}
	runtime.ReadMemStats(stat1)

	for i := 0; i < 1_000_000; i++ {
		list = append(list, iputils.ParseIP("fd00:6868:6868:0:10ac:d056:3bf6:7452"))
	}

	//runtime.GC()

	var stat2 = &runtime.MemStats{}
	runtime.ReadMemStats(stat2)

	t.Log((stat2.Alloc-stat1.Alloc)>>10, "KB", (stat2.HeapInuse-stat1.HeapInuse)>>10, "KB")

	// hold the memory
	for _, v := range list {
		_ = v
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iputils.ParseIP("fd00:6868:6868:0:10ac:d056:3bf6:7452")
	}
}

func BenchmarkToLongV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iputils.ToLong("192.168.2.100")
	}
}

func BenchmarkToLongV6(b *testing.B) {
	runtime.GOMAXPROCS(1)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		iputils.ToLong("fd00:6868:6868:0:10ac:d056:3bf6:7452")
	}
}
