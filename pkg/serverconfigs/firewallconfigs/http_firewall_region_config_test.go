// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package firewallconfigs_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/regionconfigs"
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHTTPFirewallRegionConfig_IsAllowed(t *testing.T) {
	var a = assert.NewAssertion(t)

	{
		var config = &firewallconfigs.HTTPFirewallRegionConfig{}
		config.AllowCountryIds = []int64{1, 2, 3}
		config.DenyCountryIds = []int64{4, 5, 6}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsAllowedCountry(1, 1))
		a.IsFalse(config.IsAllowedCountry(0, 0))
		a.IsFalse(config.IsAllowedCountry(4, 0))
	}

	{
		var config = &firewallconfigs.HTTPFirewallRegionConfig{}
		config.AllowCountryIds = []int64{1, 2, 3}
		config.DenyCountryIds = []int64{1, 2, 3}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsAllowedCountry(1, 1))
		a.IsFalse(config.IsAllowedCountry(0, 0))
		a.IsFalse(config.IsAllowedCountry(4, 0))
	}

	{
		var config = &firewallconfigs.HTTPFirewallRegionConfig{}
		config.AllowCountryIds = []int64{}
		config.DenyCountryIds = []int64{1, 2, 3}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(config.IsAllowedCountry(1, 0))
		a.IsTrue(config.IsAllowedCountry(0, 0))
		a.IsTrue(config.IsAllowedCountry(4, 0))
	}

	{
		var config = &firewallconfigs.HTTPFirewallRegionConfig{}
		config.AllowProvinceIds = []int64{1, 2, 3}
		config.DenyProvinceIds = []int64{4, 5, 6}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsAllowedProvince(1, 1))
		a.IsFalse(config.IsAllowedProvince(1, 0))
		a.IsFalse(config.IsAllowedProvince(1, 4))
	}

	{
		var config = &firewallconfigs.HTTPFirewallRegionConfig{}
		config.AllowProvinceIds = []int64{}
		config.DenyProvinceIds = []int64{4, 5, 6}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsAllowedProvince(1, 1))
		a.IsTrue(config.IsAllowedProvince(1, 3))
		a.IsFalse(config.IsAllowedProvince(1, 4))
	}

	{
		var config = &firewallconfigs.HTTPFirewallRegionConfig{}
		config.AllowProvinceIds = []int64{}
		config.DenyProvinceIds = []int64{}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsAllowedProvince(1, 1))
		a.IsTrue(config.IsAllowedProvince(1, 4))
		a.IsTrue(config.IsAllowedProvince(1, 4))
	}

	// the greater China area: Taiwan & Hongkong & Macao
	{
		var config = &firewallconfigs.HTTPFirewallRegionConfig{}
		config.AllowCountryIds = []int64{
			regionconfigs.RegionChinaIdHK,
			regionconfigs.RegionChinaIdMO,
			regionconfigs.RegionChinaIdTW,
			//regionconfigs.RegionChinaIdMainland,
		}
		config.DenyCountryIds = []int64{}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsAllowedCountry(1, regionconfigs.RegionChinaProvinceIdHK))
		a.IsTrue(config.IsAllowedCountry(1, regionconfigs.RegionChinaProvinceIdTW))
		a.IsTrue(config.IsAllowedCountry(1, regionconfigs.RegionChinaProvinceIdMO))
		a.IsFalse(config.IsAllowedCountry(1, 1))
		a.IsFalse(config.IsAllowedCountry(1, 0))
	}

	{
		var config = &firewallconfigs.HTTPFirewallRegionConfig{}
		config.AllowCountryIds = []int64{
			regionconfigs.RegionChinaIdHK,
			regionconfigs.RegionChinaIdMainland,
		}
		config.DenyCountryIds = []int64{}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsAllowedCountry(1, regionconfigs.RegionChinaProvinceIdHK))
		a.IsFalse(config.IsAllowedCountry(1, regionconfigs.RegionChinaProvinceIdTW))
		a.IsFalse(config.IsAllowedCountry(1, regionconfigs.RegionChinaProvinceIdMO))
		a.IsTrue(config.IsAllowedCountry(1, 1))
		a.IsFalse(config.IsAllowedCountry(1, 0))
	}

	{
		var config = &firewallconfigs.HTTPFirewallRegionConfig{}
		config.AllowCountryIds = []int64{}
		config.DenyCountryIds = []int64{
			regionconfigs.RegionChinaIdHK,
			regionconfigs.RegionChinaIdMainland,
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(config.IsAllowedCountry(1, regionconfigs.RegionChinaProvinceIdHK))
		a.IsTrue(config.IsAllowedCountry(1, regionconfigs.RegionChinaProvinceIdTW))
		a.IsTrue(config.IsAllowedCountry(1, regionconfigs.RegionChinaProvinceIdMO))
		a.IsFalse(config.IsAllowedCountry(1, 1))
		a.IsTrue(config.IsAllowedCountry(1, 0))
	}
}

func Benchmark_HTTPFirewallRegionConfig_Map(b *testing.B) {
	var m = map[int64]bool{}
	const max = 50
	for i := 0; i < max; i++ {
		m[int64(i)] = true
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m[int64(i%max)]
	}
}

func Benchmark_HTTPFirewallRegionConfig_Slice(b *testing.B) {
	var m = []int64{}
	const max = 50
	for i := 0; i < max; i++ {
		m = append(m, int64(i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var l = int64(i % max)
		for _, v := range m {
			if v == l {
				break
			}
		}
	}
}
