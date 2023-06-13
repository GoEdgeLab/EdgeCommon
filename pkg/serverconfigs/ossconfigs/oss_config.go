// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package ossconfigs

type OSSConfig struct {
}

func NewOSSConfig() *OSSConfig {
	return &OSSConfig{}
}

func (this *OSSConfig) Init() error {
	return nil
}

func (this *OSSConfig) Summary() string {
	return ""
}
