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

func MatchUserRegion(userCountryId int64, userProvinceId int64, regionId int64) bool {
	if userCountryId == RegionChinaId {
		switch regionId {
		case RegionChinaIdMainland: // china.mainland
			return CheckRegionProvinceIsInChinaMainland(userProvinceId)
		case RegionChinaIdHK: // china.hk
			return userProvinceId == RegionChinaProvinceIdHK
		case RegionChinaIdMO: // china.mo
			return userProvinceId == RegionChinaProvinceIdMO
		case RegionChinaIdTW: // china.tw
			return userProvinceId == RegionChinaProvinceIdTW
		}
	}

	return userCountryId == regionId
}
