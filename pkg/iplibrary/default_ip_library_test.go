// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary_test

import (
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/iplibrary"
	"net"
	"runtime"
	"runtime/debug"
	"testing"
	"time"
)

func TestIPLibrary_Init(t *testing.T) {
	var lib = iplibrary.NewIPLibrary()

	err := lib.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestIPLibrary_Lookup(t *testing.T) {
	var stat1 = &runtime.MemStats{}
	runtime.ReadMemStats(stat1)

	var lib = iplibrary.NewIPLibrary()

	var before = time.Now()

	err := lib.Init()
	if err != nil {
		t.Fatal(err)
	}

	var costMs = time.Since(before).Seconds() * 1000
	runtime.GC()
	debug.FreeOSMemory()

	var stat2 = &runtime.MemStats{}
	runtime.ReadMemStats(stat2)

	t.Log((stat2.Alloc-stat1.Alloc)/1024/1024, "M", fmt.Sprintf("%.2f", costMs), "ms")

	for _, ip := range []string{
		"127.0.0.1",
		"8.8.8.8",
		"4.4.4.4",
		"202.96.0.20",
		"66.249.66.69",
		"2222",                           // wrong ip
		"2406:8c00:0:3401:133:18:168:70", // ipv6
	} {
		var result = lib.Lookup(net.ParseIP(ip))
		t.Log(ip, "=>", result.IsOk(), "[", result.CountryName(), result.CountryId(), "][", result.ProvinceName(), result.ProvinceId(), "][", result.TownName(), result.TownId(), "][", result.ProviderName(), result.ProviderId(), "]")
	}
}

func TestIPLibrary_LookupIP(t *testing.T) {
	var lib = iplibrary.NewIPLibrary()
	err := lib.Init()
	if err != nil {
		t.Fatal(err)
	}

	for _, ip := range []string{
		"66.249.66.69",
	} {
		var result = lib.LookupIP(ip)
		if result.IsOk() {
			t.Log(ip, "=>", result.IsOk(), "[", result.CountryName(), result.CountryId(), "][", result.ProvinceName(), result.ProvinceId(), "][", result.TownName(), result.TownId(), "][", result.ProviderName(), result.ProviderId(), "]")
		} else {
			t.Log(ip, "=>", result.IsOk())
		}
	}
}

func BenchmarkIPLibrary_Lookup(b *testing.B) {
	var lib = iplibrary.NewIPLibrary()
	err := lib.Init()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = lib.LookupIP("66.249.66.69")
	}
}
