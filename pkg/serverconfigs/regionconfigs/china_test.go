// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package regionconfigs_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/regionconfigs"
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestMatchUserRegion(t *testing.T) {
	var a = assert.NewAssertion(t)
	a.IsTrue(regionconfigs.MatchUserRegion(regionconfigs.RegionChinaId, 1, 1))
	a.IsTrue(regionconfigs.MatchUserRegion(regionconfigs.RegionChinaId, 1, regionconfigs.RegionChinaIdMainland))
	a.IsTrue(regionconfigs.MatchUserRegion(regionconfigs.RegionChinaId, regionconfigs.RegionChinaProvinceIdHK, regionconfigs.RegionChinaIdHK))
	a.IsTrue(regionconfigs.MatchUserRegion(regionconfigs.RegionChinaId, regionconfigs.RegionChinaProvinceIdMO, regionconfigs.RegionChinaIdMO))
	a.IsTrue(regionconfigs.MatchUserRegion(regionconfigs.RegionChinaId, regionconfigs.RegionChinaProvinceIdTW, regionconfigs.RegionChinaIdTW))
	a.IsFalse(regionconfigs.MatchUserRegion(regionconfigs.RegionChinaId, 0, regionconfigs.RegionChinaIdHK))
	a.IsFalse(regionconfigs.MatchUserRegion(regionconfigs.RegionChinaId, 1, regionconfigs.RegionChinaIdHK))
}
