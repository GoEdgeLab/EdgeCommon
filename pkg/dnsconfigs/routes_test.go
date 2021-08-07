// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

import (
	"strings"
	"testing"
)

func TestRoutes(t *testing.T) {
	// 检查代号是否有空，或者代号是否重复

	var codeMap = map[string]bool{} // code => true

	for _, routes := range [][]*Route{
		AllDefaultChinaProvinceRoutes,
		AllDefaultISPRoutes,
		AllDefaultWorldRegionRoutes,
	} {
		for _, route := range routes {
			if len(route.Name) == 0 {
				t.Fatal("route name should not empty:", route)
			}
			if len(route.AliasNames) == 0 {
				t.Fatal("route alias names should not empty:", route)
			}
			if len(route.Code) == 0 || route.Code == "world:" {
				t.Fatal("route code should not empty:", route)
			}

			_, ok := codeMap[route.Code]
			if ok {
				t.Fatal("code duplicated:", route)
			}

			codeMap[route.Code] = true

			if strings.HasPrefix(route.Code, "world:sp:") || (strings.HasPrefix(route.Code, "world:") && route.Code != "world:UAR" && len(route.Code) != 8) {
				t.Log("no shorten code:", route)
			}
		}
	}
}
