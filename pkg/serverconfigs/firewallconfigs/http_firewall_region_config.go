package firewallconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type HTTPFirewallRegionConfig struct {
	IsOn            bool    `yaml:"isOn" json:"isOn"`
	DenyCountryIds  []int64 `yaml:"denyCountryIds" json:"denyCountryIds"`   // 封禁的国家|地区
	DenyProvinceIds []int64 `yaml:"denyProvinceIds" json:"denyProvinceIds"` // 封禁的省或自治区

	CountryOnlyURLPatterns   []*shared.URLPattern `yaml:"countryOnlyURLPatterns" json:"countryOnlyURLPatterns"`     // 仅限的URL
	CountryExceptURLPatterns []*shared.URLPattern `yaml:"countryExceptURLPatterns" json:"countryExceptURLPatterns"` // 排除的URL

	ProvinceOnlyURLPatterns   []*shared.URLPattern `yaml:"provinceOnlyURLPatterns" json:"provinceOnlyURLPatterns"`     // 仅限的URL
	ProvinceExceptURLPatterns []*shared.URLPattern `yaml:"provinceExceptURLPatterns" json:"provinceExceptURLPatterns"` // 排除的URL

	isNotEmpty bool
}

func (this *HTTPFirewallRegionConfig) Init() error {
	this.isNotEmpty = len(this.DenyCountryIds) > 0 || len(this.DenyProvinceIds) > 0

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
