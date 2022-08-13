// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary_test

import (
	"bytes"
	"github.com/TeaOSLab/EdgeCommon/pkg/iplibrary"
	"github.com/iwind/TeaGo/rands"
	"github.com/iwind/TeaGo/types"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"net"
	"runtime"
	"testing"
	"time"
)

func TestNewReader(t *testing.T) {
	var buf = &bytes.Buffer{}
	var writer = iplibrary.NewWriter(buf, &iplibrary.Meta{
		Author: "GoEdge <https://goedge.cn>",
	})

	err := writer.WriteMeta()
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.1.100", "192.168.1.100", 1, 200, 300, 400, 500)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.2.100", "192.168.3.100", 2, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.3.101", "192.168.3.101", 3, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.0.101", "192.168.0.200", 4, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("::1", "::5", 5, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	/**var n = func() string {
		return types.String(rands.Int(0, 255))
	}

	for i := 0; i < 1_000_000; i++ {
		err = writer.Write(n()+"."+n()+"."+n()+"."+n(), n()+"."+n()+"."+n()+"."+n(), int64(i)+100, 201, 301, 401, 501)
		if err != nil {
			t.Fatal(err)
		}
	}**/

	var stat = &runtime.MemStats{}
	runtime.ReadMemStats(stat)
	reader, err := iplibrary.NewReader(buf)

	var stat2 = &runtime.MemStats{}
	runtime.ReadMemStats(stat2)
	t.Log((stat2.Alloc-stat.Alloc)/1024/1024, "M")

	if err != nil {
		t.Fatal(err)
	}
	t.Log("version:", reader.Meta().Version, "author:", reader.Meta().Author, "createdTime:", timeutil.FormatTime("Y-m-d H:i:s", reader.Meta().CreatedAt))

	if len(reader.IPv4Items()) < 10 {
		t.Log("===")
		for _, item := range reader.IPv4Items() {
			t.Logf("%+v", item)
		}
		t.Log("===")
	}
	if len(reader.IPv6Items()) < 10 {
		t.Log("===")
		for _, item := range reader.IPv6Items() {
			t.Logf("%+v", item)
		}
		t.Log("===")
	}

	var before = time.Now()
	for _, ip := range []string{
		"192.168.0.1",
		"192.168.0.150",
		"192.168.1.100",
		"192.168.2.100",
		"192.168.3.50",
		"192.168.0.150",
		"192.168.4.80",
		"::3",
		"::8",
	} {
		var result = reader.Lookup(net.ParseIP(ip))
		if result.IsOk() {
			t.Log(ip+":", "countryId:", result.CountryId())
		} else {
			t.Log(ip+":", "not found")
		}
	}
	t.Log(time.Since(before).Seconds()*1000, "ms")
}

func BenchmarkNewReader(b *testing.B) {
	runtime.GOMAXPROCS(1)

	var buf = &bytes.Buffer{}
	var writer = iplibrary.NewWriter(buf, &iplibrary.Meta{
		Author: "GoEdge <https://goedge.cn>",
	})

	err := writer.WriteMeta()
	if err != nil {
		b.Fatal(err)
	}

	var n = func() string {
		return types.String(rands.Int(0, 255))
	}

	for i := 0; i < 1_000_000; i++ {
		err = writer.Write(n()+"."+n()+"."+n()+"."+n(), n()+"."+n()+"."+n()+"."+n(), int64(i)+100, 201, 301, 401, 501)
		if err != nil {
			b.Fatal(err)
		}
	}

	reader, err := iplibrary.NewReader(buf)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ip = "192.168.1.100"
		reader.Lookup(net.ParseIP(ip))
	}
}
