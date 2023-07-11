// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

import "github.com/tdewolff/minify/v2/js"

type HTTPJavascriptOptimizationConfig struct {
	IsOn bool `yaml:"isOn" json:"isOn"`

	Precision    int  `yaml:"precision" json:"precision"`
	Version      int  `yaml:"version" json:"version"`
	KeepVarNames bool `yaml:"keepVarNames" json:"keepVarNames"`
}

func NewHTTPJavascriptOptimizationConfig() *HTTPJavascriptOptimizationConfig {
	return &HTTPJavascriptOptimizationConfig{}
}

func (this *HTTPJavascriptOptimizationConfig) Init() error {
	return nil
}

func (this *HTTPJavascriptOptimizationConfig) AsMinifier() *js.Minifier {
	return &js.Minifier{
		Precision:    this.Precision,
		KeepVarNames: this.KeepVarNames,
		Version:      this.Version,
	}
}
