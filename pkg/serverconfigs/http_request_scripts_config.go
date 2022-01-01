// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type HTTPRequestScriptsConfig struct {
	OnInitScript    *JSScriptConfig `yaml:"onInitScript" json:"onInitScript"`       // 接收到请求之后
	OnRequestScript *JSScriptConfig `yaml:"onRequestScript" json:"onRequestScript"` // 准备转发请求之前
}

func (this *HTTPRequestScriptsConfig) Init() error {
	if this.OnInitScript != nil {
		err := this.OnInitScript.Init()
		if err != nil {
			return err
		}
	}

	if this.OnRequestScript != nil {
		err := this.OnRequestScript.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *HTTPRequestScriptsConfig) IsEmpty() bool {
	if (this.OnInitScript == nil || !this.OnInitScript.IsOn) && (this.OnRequestScript == nil || !this.OnRequestScript.IsOn) {
		return true
	}
	return false
}
