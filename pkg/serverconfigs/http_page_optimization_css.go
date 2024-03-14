// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

type HTTPCSSOptimizationConfig struct {
	HTTPBaseOptimizationConfig

	IsOn bool `yaml:"isOn" json:"isOn"`

	Precision int  `yaml:"precision" json:"precision"`
	KeepCSS2  bool `yaml:"keepCSS2" json:"keepCSS2"`
}

func NewHTTPCSSOptimizationConfig() *HTTPCSSOptimizationConfig {
	return &HTTPCSSOptimizationConfig{
		KeepCSS2: true,
	}
}

func (this *HTTPCSSOptimizationConfig) Init() error {
	err := this.HTTPBaseOptimizationConfig.Init()
	if err != nil {
		return err
	}
	return nil
}
