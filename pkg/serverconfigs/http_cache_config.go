package serverconfigs

import (
	"encoding/json"
	"github.com/iwind/TeaGo/rands"
)

type HTTPCacheConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`
	IsOn    bool `yaml:"isOn" json:"isOn"`

	AddStatusHeader          bool `yaml:"addStatusHeader" json:"addStatusHeader"`                   // 是否增加命中状态Header（X-Cache）
	AddAgeHeader             bool `yaml:"addAgeHeader" json:"addAgeHeader"`                         // 是否增加Age Header
	EnableCacheControlMaxAge bool `yaml:"enableCacheControlMaxAge" json:"enableCacheControlMaxAge"` // 是否支持Cache-Control: max-age=...

	PurgeIsOn bool   `yaml:"purgeIsOn" json:"purgeIsOn"` // 是否允许使用Purge方法清理
	PurgeKey  string `yaml:"purgeKey" json:"purgeKey"`   // Purge时使用的X-Edge-Purge-Key

	CacheRefs []*HTTPCacheRef `yaml:"cacheRefs" json:"cacheRefs"` // 缓存配置
}

func (this *HTTPCacheConfig) Init() error {
	for _, cacheRef := range this.CacheRefs {
		err := cacheRef.Init()
		if err != nil {
			return err
		}
	}

	if this.PurgeIsOn && len(this.PurgeKey) == 0 {
		this.PurgeKey = rands.HexString(32)
	}

	return nil
}

func (this *HTTPCacheConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
