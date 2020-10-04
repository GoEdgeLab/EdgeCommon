package serverconfigs

type HTTPCacheConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`
	IsOn    bool `yaml:"isOn" json:"isOn"`

	CacheRefs []*HTTPCacheRef `yaml:"cacheRefs" json:"cacheRefs"` // 缓存配置
}

func (this *HTTPCacheConfig) Init() error {
	for _, cacheRef := range this.CacheRefs {
		err := cacheRef.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
