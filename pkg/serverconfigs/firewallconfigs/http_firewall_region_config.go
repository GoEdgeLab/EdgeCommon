package firewallconfigs

type HTTPFirewallRegionConfig struct {
	IsOn           bool    `yaml:"isOn" json:"isOn"`
	DenyCountryIds []int64 `yaml:"denyCountryIds" json:"denyCountryIds"`   // 封禁的国家|地区
	DenyProvinces  []int64 `yaml:"denyProvinceIds" json:"denyProvinceIds"` // 封禁的省或自治区
}
