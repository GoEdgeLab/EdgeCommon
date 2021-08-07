// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

import "fmt"

type NSNodeConfig struct {
	Id           int64           `yaml:"id" json:"id"`
	NodeId       string          `yaml:"nodeId" json:"nodeId"`
	ClusterId    int64           `yaml:"clusterId" json:"clusterId"`
	AccessLogRef *NSAccessLogRef `yaml:"accessLogRef" json:"accessLogRef"`

	paddedId string
}

func (this *NSNodeConfig) Init() error {
	this.paddedId = fmt.Sprintf("%08d", this.Id)

	// accessLog
	if this.AccessLogRef != nil {
		err := this.AccessLogRef.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *NSNodeConfig) PaddedId() string {
	return this.paddedId
}
