// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ddosconfigs

func DefaultProtectionConfig() *ProtectionConfig {
	return &ProtectionConfig{}
}

type ProtectionConfig struct {
	TCP *TCPConfig `yaml:"tcp" json:"tcp"`
}

func (this *ProtectionConfig) Init() error {
	// tcp
	if this.TCP != nil {
		err := this.TCP.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *ProtectionConfig) IsPriorEmpty() bool {
	if this.TCP != nil && this.TCP.IsPrior {
		return false
	}

	return true
}

func (this *ProtectionConfig) IsOn() bool {
	// tcp
	if this.TCP != nil && this.TCP.IsOn {
		return true
	}

	return false
}

func (this *ProtectionConfig) Merge(childConfig *ProtectionConfig) {
	if childConfig == nil {
		return
	}

	// tcp
	if childConfig.TCP != nil && childConfig.TCP.IsPrior {
		this.TCP = childConfig.TCP
	}
}
