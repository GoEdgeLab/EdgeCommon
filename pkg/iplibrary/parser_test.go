// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/iplibrary"
	"testing"
)

func TestNewParser(t *testing.T) {
	template, err := iplibrary.NewTemplate("${ipFrom}|${ipTo}|${country}|${any}|${province}|${city}|${provider}")
	if err != nil {
		t.Fatal(err)
	}

	parser, err := iplibrary.NewParser(&iplibrary.ParserConfig{
		Template:    template,
		EmptyValues: []string{"0"},
		Iterator: func(values map[string]string) error {
			t.Log(values)
			return nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	parser.Write([]byte(`0.0.0.0|0.255.255.255|0|0|0|内网IP|内网IP
1.0.0.0|1.0.0.255|澳大利亚|0|0|0|0
1.0.1.0|1.0.3.255|中国|0|福建省|福州市|电信
1.0.4.0|1.0.7.255|澳大利亚|0|维多利亚|墨尔本|0
1.0.8.0|1.0.15.255|中国|0|广东省|广州市|电信
1.0.16.0|1.0.31.255|日本|0|0|0|0
1.0.32.0|1.0.63.255|中国|0|广东省|广州市|电信
1.0.64.0|1.0.79.255|日本|0|广岛县|0|0
1.0.80.0|1.0.127.255|日本|0|冈山县|0|0
1.0.128.0|1.0.128.255|泰国|0|清莱府|0|TOT
1.0.129.0|1.0.132.191|泰国|0|曼谷|曼谷|TOT`))

	err = parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
}
