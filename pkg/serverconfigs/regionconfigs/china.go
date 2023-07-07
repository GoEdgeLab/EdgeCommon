// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package regionconfigs

type RegionId = int64
type RegionProvinceId = int64

const (
	RegionChinaId         RegionId = 1
	RegionChinaIdHK       RegionId = 261
	RegionChinaIdTW       RegionId = 262
	RegionChinaIdMO       RegionId = 263
	RegionChinaIdMainland RegionId = 264

	RegionChinaProvinceIdHK RegionProvinceId = 32
	RegionChinaProvinceIdTW RegionProvinceId = 34
	RegionChinaProvinceIdMO RegionProvinceId = 33
)

func CheckRegionIsInGreaterChina(regionId RegionId) bool {
	return regionId == RegionChinaId ||
		regionId == RegionChinaIdHK ||
		regionId == RegionChinaIdTW ||
		regionId == RegionChinaIdMO ||
		regionId == RegionChinaIdMainland
}

func FindAllGreaterChinaSubRegionIds() []RegionId {
	return []RegionId{
		RegionChinaIdMainland, RegionChinaIdHK, RegionChinaIdMO, RegionChinaIdTW,
	}
}

func CheckRegionProvinceIsInChinaMainland(regionProvinceId RegionProvinceId) bool {
	if regionProvinceId <= 0 {
		return false
	}
	return regionProvinceId != RegionChinaProvinceIdHK &&
		regionProvinceId != RegionChinaProvinceIdMO &&
		regionProvinceId != RegionChinaProvinceIdTW
}
