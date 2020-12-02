package systemconfigs

// 安全相关配置
type SecurityConfig struct {
	Frame            string  `json:"frame"`
	AllowCountryIds  []int64 `json:"allowCountryIds"`
	AllowProvinceIds []int64 `json:"allowProvinceIds"`
	AllowLocal       bool    `json:"allowLocal"`
}
