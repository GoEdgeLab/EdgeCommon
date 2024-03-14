// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

type HTTPJavascriptOptimizationConfig struct {
	HTTPBaseOptimizationConfig

	IsOn bool `yaml:"isOn" json:"isOn"`

	Precision    int  `yaml:"precision" json:"precision"`
	Version      int  `yaml:"version" json:"version"`
	KeepVarNames bool `yaml:"keepVarNames" json:"keepVarNames"`
}

func NewHTTPJavascriptOptimizationConfig() *HTTPJavascriptOptimizationConfig {
	return &HTTPJavascriptOptimizationConfig{}
}

func (this *HTTPJavascriptOptimizationConfig) Init() error {
	err := this.HTTPBaseOptimizationConfig.Init()
	if err != nil {
		return err
	}
	return nil
}
