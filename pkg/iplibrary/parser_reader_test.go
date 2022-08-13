// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary_test

import (
	"bytes"
	"github.com/TeaOSLab/EdgeCommon/pkg/iplibrary"
	"testing"
)

func TestNewReaderParser(t *testing.T) {
	template, err := iplibrary.NewTemplate("${ipFrom}|${ipTo}|${country}|${any}|${province}|${city}|${provider}")
	if err != nil {
		t.Fatal(err)
	}

	var buf = &bytes.Buffer{}
	buf.WriteString(`8.45.160.0|8.45.161.255|美国|0|华盛顿|西雅图|Level3
8.45.162.0|8.45.162.255|美国|0|马萨诸塞|0|Level3
8.45.163.0|8.45.164.255|美国|0|俄勒冈|0|Level3
8.45.165.0|8.45.165.255|美国|0|华盛顿|0|Level3
8.45.166.0|8.45.167.255|美国|0|华盛顿|西雅图|Level3
8.45.168.0|8.127.255.255|美国|0|0|0|Level3
8.128.0.0|8.128.3.255|中国|0|上海|上海市|阿里巴巴
8.128.4.0|8.128.255.255|中国|0|0|0|阿里巴巴
8.129.0.0|8.129.255.255|中国|0|广东省|深圳市|阿里云
8.130.0.0|8.130.55.255|中国|0|北京|北京市|阿里云
8.130.56.0|8.131.255.255|中国|0|北京|北京市|阿里巴巴
8.132.0.0|8.133.255.255|中国|0|上海|上海市|阿里巴巴
8.134.0.0|8.134.20.63|中国|0|0|0|阿里云
8.134.20.64|8.134.79.255|中国|0|广东省|深圳市|阿里云
8.134.80.0|8.191.255.255|中国|0|0|0|阿里巴巴
8.192.0.0|8.192.0.255|美国|0|0|0|Level3
8.192.1.0|8.192.1.255|美国|0|马萨诸塞|波士顿|Level3
8.192.2.0|8.207.255.255|美国|0|0|0|Level3
8.208.0.0|8.208.127.255|英国|0|伦敦|伦敦|阿里云
8.208.128.0|8.208.255.255|英国|0|伦敦|伦敦|阿里巴巴
8.209.0.0|8.209.15.255|俄罗斯|0|莫斯科|莫斯科|阿里云
8.209.16.0|8.209.31.255|俄罗斯|0|莫斯科|莫斯科|阿里巴巴
8.209.32.0|8.209.32.15|中国|0|台湾省|台北|阿里云
8.209.32.16|8.209.32.255|中国|0|台湾省|台北|阿里巴巴
8.209.33.0|8.209.34.255|中国|0|台湾省|台北|阿里云
8.209.35.0|8.209.35.255|中国|0|台湾省|台北|阿里巴巴`)

	var count int
	parser, err := iplibrary.NewReaderParser(buf, &iplibrary.ParserConfig{
		Template:    template,
		EmptyValues: []string{"0"},
		Iterator: func(values map[string]string) error {
			count++
			t.Log(count, values)
			return nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	err = parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
}
