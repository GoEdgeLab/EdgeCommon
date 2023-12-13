// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

// HTTPCacheKeyConfig 缓存Key配置
type HTTPCacheKeyConfig struct {
	IsOn   bool   `yaml:"isOn" json:"isOn"`
	Scheme string `yaml:"scheme" json:"scheme"`
	Host   string `yaml:"host" json:"host"`
}

func (this *HTTPCacheKeyConfig) Init() error {
	return nil
}
