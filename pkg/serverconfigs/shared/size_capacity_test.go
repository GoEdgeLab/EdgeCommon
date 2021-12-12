// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package shared

import "testing"

func TestSizeCapacity_Bytes(t *testing.T) {
	for _, unit := range []string{
		SizeCapacityUnitByte,
		SizeCapacityUnitKB,
		SizeCapacityUnitMB,
		SizeCapacityUnitGB,
		SizeCapacityUnitTB,
		SizeCapacityUnitPB,
		SizeCapacityUnitEB,
	} {
		var capacity = &SizeCapacity{
			Count: 1,
			Unit:  unit,
		}
		t.Log(unit, capacity.Bytes())
	}
}
