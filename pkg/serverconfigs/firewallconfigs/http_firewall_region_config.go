package firewallconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/regionconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"strings"
)

type HTTPFirewallRegionConfig struct {
	IsOn bool `yaml:"isOn" json:"isOn"`

	AllowCountryIds  []int64 `yaml:"allowCountryIds" json:"allowCountryIds"`   // 允许的国家/地区
	DenyCountryIds   []int64 `yaml:"denyCountryIds" json:"denyCountryIds"`     // 封禁的国家/地区
	AllowProvinceIds []int64 `yaml:"allowProvinceIds" json:"allowProvinceIds"` // 允许的省或自治区
	DenyProvinceIds  []int64 `yaml:"denyProvinceIds" json:"denyProvinceIds"`   // 封禁的省或自治区

	CountryOnlyURLPatterns   []*shared.URLPattern `yaml:"countryOnlyURLPatterns" json:"countryOnlyURLPatterns"`     // 仅限的URL
	CountryExceptURLPatterns []*shared.URLPattern `yaml:"countryExceptURLPatterns" json:"countryExceptURLPatterns"` // 排除的URL
	CountryHTML              string               `yaml:"countryHTML" json:"countryHTML"`                           // 提示HTML

	ProvinceOnlyURLPatterns   []*shared.URLPattern `yaml:"provinceOnlyURLPatterns" json:"provinceOnlyURLPatterns"`     // 仅限的URL
	ProvinceExceptURLPatterns []*shared.URLPattern `yaml:"provinceExceptURLPatterns" json:"provinceExceptURLPatterns"` // 排除的URL
	ProvinceHTML              string               `yaml:"provinceHTML" json:"provinceHTML"`                           // 提示HTML

	isNotEmpty bool

	allowCountryIdMap  map[int64]bool
	denyCountryIdMap   map[int64]bool
	allowProvinceIdMap map[int64]bool
	denyProvinceIdMap  map[int64]bool
}

func (this *HTTPFirewallRegionConfig) Init() error {
	// countries and provinces
	this.isNotEmpty = len(this.AllowCountryIds) > 0 || len(this.AllowProvinceIds) > 0 || len(this.DenyCountryIds) > 0 || len(this.DenyProvinceIds) > 0
	this.allowCountryIdMap = map[int64]bool{}
	for _, countryId := range this.AllowCountryIds {
		this.allowCountryIdMap[countryId] = true
	}

	this.denyCountryIdMap = map[int64]bool{}
	for _, countryId := range this.DenyCountryIds {
		this.denyCountryIdMap[countryId] = true
	}

	this.CountryHTML = strings.TrimSpace(this.CountryHTML)

	this.allowProvinceIdMap = map[int64]bool{}
	for _, provinceId := range this.AllowProvinceIds {
		this.allowProvinceIdMap[provinceId] = true
	}

	this.denyProvinceIdMap = map[int64]bool{}
	for _, provinceId := range this.DenyProvinceIds {
		this.denyProvinceIdMap[provinceId] = true
	}

	this.ProvinceHTML = strings.TrimSpace(this.ProvinceHTML)

	// url patterns
	for _, pattern := range this.CountryExceptURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	for _, pattern := range this.CountryOnlyURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	for _, pattern := range this.ProvinceExceptURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	for _, pattern := range this.ProvinceOnlyURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *HTTPFirewallRegionConfig) IsNotEmpty() bool {
	return this.isNotEmpty
}

func (this *HTTPFirewallRegionConfig) IsAllowedCountry(countryId int64, provinceId int64) bool {
	if len(this.allowCountryIdMap) > 0 {
		if this.allowCountryIdMap[countryId] {
			return true
		}

		// china sub regions
		if countryId == regionconfigs.RegionChinaId && provinceId > 0 {
			if this.allowCountryIdMap[regionconfigs.RegionChinaIdHK] && provinceId == regionconfigs.RegionChinaProvinceIdHK {
				return true
			}
			if this.allowCountryIdMap[regionconfigs.RegionChinaIdMO] && provinceId == regionconfigs.RegionChinaProvinceIdMO {
				return true
			}
			if this.allowCountryIdMap[regionconfigs.RegionChinaIdTW] && provinceId == regionconfigs.RegionChinaProvinceIdTW {
				return true
			}
			if this.allowCountryIdMap[regionconfigs.RegionChinaIdMainland] && regionconfigs.CheckRegionProvinceIsInChinaMainland(provinceId) {
				return true
			}
		}

		return false
	}
	if len(this.denyCountryIdMap) > 0 {
		if !this.denyCountryIdMap[countryId] {
			// china sub regions
			if countryId == regionconfigs.RegionChinaId && provinceId > 0 {
				if this.denyCountryIdMap[regionconfigs.RegionChinaIdHK] && provinceId == regionconfigs.RegionChinaProvinceIdHK {
					return false
				}
				if this.denyCountryIdMap[regionconfigs.RegionChinaIdMO] && provinceId == regionconfigs.RegionChinaProvinceIdMO {
					return false
				}
				if this.denyCountryIdMap[regionconfigs.RegionChinaIdTW] && provinceId == regionconfigs.RegionChinaProvinceIdTW {
					return false
				}
				if this.denyCountryIdMap[regionconfigs.RegionChinaIdMainland] && regionconfigs.CheckRegionProvinceIsInChinaMainland(provinceId) {
					return false
				}
			}

			return true
		}

		return false
	}

	return true
}

func (this *HTTPFirewallRegionConfig) IsAllowedProvince(countryId int64, provinceId int64) bool {
	if countryId != regionconfigs.RegionChinaId {
		return true
	}
	if len(this.allowProvinceIdMap) > 0 {
		return this.allowProvinceIdMap[provinceId]
	}
	if len(this.denyProvinceIdMap) > 0 {
		return !this.denyProvinceIdMap[provinceId]
	}
	return true
}

func (this *HTTPFirewallRegionConfig) MatchCountryURL(url string) bool {
	// except
	if len(this.CountryExceptURLPatterns) > 0 {
		for _, pattern := range this.CountryExceptURLPatterns {
			if pattern.Match(url) {
				return false
			}
		}
	}

	if len(this.CountryOnlyURLPatterns) > 0 {
		for _, pattern := range this.CountryOnlyURLPatterns {
			if pattern.Match(url) {
				return true
			}
		}
		return false
	}

	return true
}

func (this *HTTPFirewallRegionConfig) MatchProvinceURL(url string) bool {
	// except
	if len(this.ProvinceExceptURLPatterns) > 0 {
		for _, pattern := range this.ProvinceExceptURLPatterns {
			if pattern.Match(url) {
				return false
			}
		}
	}

	if len(this.ProvinceOnlyURLPatterns) > 0 {
		for _, pattern := range this.ProvinceOnlyURLPatterns {
			if pattern.Match(url) {
				return true
			}
		}
		return false
	}

	return true
}
