// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

// HLSConfig HTTP Living Streaming相关配置
type HLSConfig struct {
	IsPrior    bool                 `yaml:"isPrior" json:"isPrior"`
	Encrypting *HLSEncryptingConfig `yaml:"encrypting" json:"encrypting"` // 加密设置
}

func (this *HLSConfig) Init() error {
	// encrypting
	if this.Encrypting != nil {
		err := this.Encrypting.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *HLSConfig) IsEmpty() bool {
	if this.Encrypting != nil && this.Encrypting.IsOn {
		return false
	}

	return true
}
