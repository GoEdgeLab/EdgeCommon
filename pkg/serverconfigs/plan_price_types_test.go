// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"testing"
)

func TestPlanBandwidthPriceConfig_Lookup(t *testing.T) {
	{
		var config = &serverconfigs.PlanBandwidthPriceConfig{}
		t.Log(config.LookupRange(1))
	}

	{
		var config = &serverconfigs.PlanBandwidthPriceConfig{
			Ranges: []*serverconfigs.PlanBandwidthPriceRangeConfig{
				{
					MinMB: 1,
					MaxMB: 1.5,
				},
				{
					MinMB: 1.5,
					MaxMB: 2,
				},
				{
					MinMB: 5,
					MaxMB: 10,
				},
				{
					MinMB: 20,
					MaxMB: 100,
				},
			},
		}
		for _, mb := range []float32{0.5, 1, 3, 5, 7, 20, 50, 1000} {
			t.Log(mb, config.LookupRange(mb))
		}
	}
}
