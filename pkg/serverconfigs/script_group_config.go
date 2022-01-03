// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type ScriptGroupConfig struct {
	IsOn    bool            `yaml:"isOn" json:"isOn"`
	IsPrior bool            `yaml:"isPrior" json:"isPrior"`
	Scripts []*ScriptConfig `yaml:"scripts" json:"scripts"`

	isEmpty bool
}

func (this *ScriptGroupConfig) Init() error {
	this.isEmpty = true

	for _, script := range this.Scripts {
		err := script.Init()
		if err != nil {
			return err
		}
		if script.IsOn {
			this.isEmpty = false
		}
	}
	return nil
}

func (this *ScriptGroupConfig) IsEmpty() bool {
	return this.isEmpty
}
