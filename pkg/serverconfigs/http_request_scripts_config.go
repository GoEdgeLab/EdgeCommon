// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type HTTPRequestScriptsConfig struct {
	InitGroup    *ScriptGroupConfig `yaml:"initGroup" json:"initGroup"`
	RequestGroup *ScriptGroupConfig `yaml:"requestGroup" json:"requestGroup"`
}

func (this *HTTPRequestScriptsConfig) Init() error {
	if this.InitGroup != nil {
		err := this.InitGroup.Init()
		if err != nil {
			return err
		}
	}

	if this.RequestGroup != nil {
		err := this.RequestGroup.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *HTTPRequestScriptsConfig) IsEmpty() bool {
	return (this.InitGroup == nil || this.InitGroup.IsEmpty()) &&
		(this.RequestGroup == nil || this.RequestGroup.IsEmpty())
}
