package firewallconfigs

type HTTPFirewallRegionConfig struct {
	IsOn            bool    `yaml:"isOn" json:"isOn"`
	DenyCountryIds  []int64 `yaml:"denyCountryIds" json:"denyCountryIds"`   // 封禁的国家|地区
	DenyProvinceIds []int64 `yaml:"denyProvinceIds" json:"denyProvinceIds"` // 封禁的省或自治区

	isNotEmpty bool
}

func (this *HTTPFirewallRegionConfig) Init() error {
	this.isNotEmpty = len(this.DenyCountryIds) > 0 || len(this.DenyProvinceIds) > 0
	return nil
}

func (this *HTTPFirewallRegionConfig) IsNotEmpty() bool {
	return this.isNotEmpty
}
