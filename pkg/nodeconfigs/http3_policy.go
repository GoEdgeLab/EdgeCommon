// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodeconfigs

const DefaultHTTP3Port = 443

type HTTP3Policy struct {
	IsOn bool `yaml:"isOn" json:"isOn"`
	Port int  `yaml:"port" json:"port"`
}

func NewHTTP3Policy() *HTTP3Policy {
	return &HTTP3Policy{
		Port: DefaultHTTP3Port,
	}
}

func (this *HTTP3Policy) Init() error {
	if this.Port <= 0 {
		this.Port = DefaultHTTP3Port
	}
	return nil
}
