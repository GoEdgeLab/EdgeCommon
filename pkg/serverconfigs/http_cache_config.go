package serverconfigs

import "encoding/json"

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

func (this *HTTPCacheConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
