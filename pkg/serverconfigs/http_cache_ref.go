package serverconfigs

type HTTPCacheRef struct {
	IsPrior       bool  `yaml:"isPrior" json:"isPrior"`             // 是否覆盖
	IsOn          bool  `yaml:"isOn" json:"isOn"`                   // 是否开启
	CachePolicyId int64 `yaml:"cachePolicyId" json:"cachePolicyId"` // 缓存策略ID
}

func (this *HTTPCacheRef) Init() error {
	return nil
}
