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
