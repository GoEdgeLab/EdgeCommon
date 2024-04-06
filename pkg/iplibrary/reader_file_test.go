// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary_test

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/iplibrary"
	"github.com/iwind/TeaGo/maps"
	stringutil "github.com/iwind/TeaGo/utils/string"
	"net"
	"testing"
)

func TestNewFileReader(t *testing.T) {
	reader, err := iplibrary.NewFileReader("./default_ip_library_plus_test.go", stringutil.Md5("123456"))
	if err != nil {
		t.Fatal(err)
	}

	for _, ip := range []string{
		"127.0.0.1",
		"192.168.0.1",
		"192.168.0.150",
		"8.8.8.8",
		"111.197.165.199",
		"175.178.206.125",
	} {
		var result = reader.Lookup(net.ParseIP(ip))
		if result.IsOk() {
			var data = maps.Map{
				"countryId":    result.CountryId(),
				"countryName":  result.CountryName(),
				"provinceId":   result.ProvinceId(),
				"provinceName": result.ProvinceName(),
				"cityId":       result.CityId(),
				"cityName":     result.CityName(),
				"townId":       result.TownId(),
				"townName":     result.TownName(),
				"providerId":   result.ProviderId(),
				"providerName": result.ProviderName(),
				"summary":      result.Summary(),
			}
			dataJSON, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			t.Log(ip, "=>", string(dataJSON))
		} else {
			t.Log(ip+":", "not found")
		}
	}
}
