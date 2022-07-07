package systemconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// SecurityConfig 安全相关配置
type SecurityConfig struct {
	Frame              string   `json:"frame"`              // Frame嵌套
	AllowCountryIds    []int64  `json:"allowCountryIds"`    // 允许的国家/地区
	AllowProvinceIds   []int64  `json:"allowProvinceIds"`   // 允许的省份
	AllowLocal         bool     `json:"allowLocal"`         // 允许本地+局域网IP访问
	AllowIPs           []string `json:"allowIPs"`           // 允许访问的IP
	AllowRememberLogin bool     `json:"allowRememberLogin"` // 是否允许在设备上记住登录
	DenySearchEngines  bool     `json:"denySearchEngines"`  // 禁止常见的搜索引擎访问
	DenySpiders        bool     `json:"denySpiders"`        // 禁止常见的爬虫
	AllowDomains       []string `json:"allowDomains"`       // 允许访问的域名

	allowIPRanges []*shared.IPRangeConfig
}

// Init 初始化
func (this *SecurityConfig) Init() error {
	this.allowIPRanges = []*shared.IPRangeConfig{}
	for _, allowIP := range this.AllowIPs {
		r, err := shared.ParseIPRange(allowIP)
		if err != nil {
			return err
		}
		this.allowIPRanges = append(this.allowIPRanges, r)
	}
	return nil
}

// AllowIPRanges 查询允许的IP区域
func (this *SecurityConfig) AllowIPRanges() []*shared.IPRangeConfig {
	return this.allowIPRanges
}
