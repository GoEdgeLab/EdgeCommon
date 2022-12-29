// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package shared

// HTTPCORSHeaderConfig 参考 https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
type HTTPCORSHeaderConfig struct {
	IsOn             bool     `yaml:"isOn" json:"isOn"`
	AllowMethods     []string `yaml:"allowMethods" json:"allowMethods"`         // TODO
	AllowOrigin      string   `yaml:"allowOrigin" json:"allowOrigin"`           // TODO
	AllowCredentials bool     `yaml:"allowCredentials" json:"allowCredentials"` // TODO
	ExposeHeaders    []string `yaml:"exposeHeaders" json:"exposeHeaders"`       // TODO
	MaxAge           int32    `yaml:"maxAge" json:"maxAge"`                     // TODO
	RequestHeaders   []string `yaml:"requestHeaders" json:"requestHeaders"`     // TODO
	RequestMethod    string   `yaml:"requestMethod" json:"requestMethod"`       // TODO
}

func (this *HTTPCORSHeaderConfig) Init() error {
	return nil
}
