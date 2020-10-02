package serverconfigs

type HTTPCacheRef struct {
	IsPrior       bool           `yaml:"isPrior" json:"isPrior"`             // 是否覆盖
	IsOn          bool           `yaml:"isOn" json:"isOn"`                   // 是否开启
	CachePolicyId int64          `yaml:"cachePolicyId" json:"cachePolicyId"` // 缓存策略ID
	Cond          *HTTPCacheCond `yaml:"cond" json:"cond"`                   // 条件
}

func (this *HTTPCacheRef) Init() error {
	if this.Cond != nil {
		err := this.Cond.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
