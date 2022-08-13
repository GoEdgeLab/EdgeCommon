// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/iplibrary"
	"net"
	"testing"
)

func TestNewFileReader(t *testing.T) {
	reader, err := iplibrary.NewFileReader("./ip.db")
	if err != nil {
		t.Fatal(err)
	}

	for _, ip := range []string{
		"127.0.0.1",
		"192.168.0.1",
		"192.168.0.150",
		"192.168.1.100",
		"192.168.2.100",
		"192.168.3.50",
		"192.168.0.150",
		"192.168.4.80",
		"8.8.8.8",
		"111.197.165.199",
		"175.178.206.125",
	} {
		var result = reader.Lookup(net.ParseIP(ip))
		if result.IsOk() {
			t.Log(ip+":", "countryId:", result.CountryId(), "provinceId:", result.ProvinceId(), "cityId:", result.CityId(), "provider:", result.ProviderId())
		} else {
			t.Log(ip+":", "not found")
		}
	}
}
