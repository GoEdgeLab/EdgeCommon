package serverconfigs

type HTTPCacheRef struct {
	IsOn          bool  `yaml:"isOn" json:"isOn"`                   // 是否开启
	CachePolicyId int64 `yaml:"cachePolicyId" json:"cachePolicyId"` // 缓存策略ID
}
