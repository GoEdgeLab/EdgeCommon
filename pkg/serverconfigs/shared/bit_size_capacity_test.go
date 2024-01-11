// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package shared_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"testing"
)

func TestBitSizeCapacity_Bits(t *testing.T) {
	{
		var capacity = shared.NewBitSizeCapacity(1, shared.BitSizeCapacityUnitB)
		t.Log(capacity.Bits())
	}
	{
		var capacity = shared.NewBitSizeCapacity(2, shared.BitSizeCapacityUnitKB)
		t.Log(capacity.Bits())
	}
	{
		var capacity = shared.NewBitSizeCapacity(3, shared.BitSizeCapacityUnitMB)
		t.Log(capacity.Bits())
	}
	{
		var capacity = shared.NewBitSizeCapacity(4, shared.BitSizeCapacityUnitGB)
		t.Log(capacity.Bits())
	}
}
