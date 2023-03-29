// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/iplibrary"
	"testing"
)

func TestNewTemplate(t *testing.T) {
	template, err := iplibrary.NewTemplate("${ipFrom}|${ipTo}|${country}|${any}|${province}|${city}|${provider}")
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{
		"42.0.32.0|42.0.63.255|中国|0|广东省|广州市|电信",
		"42.0.32.0|42.0.63.255|中国|0|广东省|广州市|电信\n123",
		"42.0.32.0|42.0.63.255|中国||广东省|广州市|电信",
		"42.0.32.0|42.0.63.255|中国|0||广州市|电信",
		"42.0.32.0|42.0.63.255|中国|0|广东省|广州市",
	} {
		values, _ := template.Extract(s, []string{})
		t.Log(s, "=>\n", values)
	}
}

func TestNewTemplate2(t *testing.T) {
	template, err := iplibrary.NewTemplate("${any},${any},${ipFrom},${ipTo},${country},${province},${city},${town},${provider},${any},${any}")
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{
		"22723584,22723839,1.90.188.0,1.90.188.255,中国,北京,北京,房山,歌华有线,102400,010,城域网",
	} {
		values, _ := template.Extract(s, []string{})
		t.Log(s, "=>\n", values)
	}
}
