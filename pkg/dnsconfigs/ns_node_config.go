// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

import "fmt"

type NSNodeConfig struct {
	Id           int64         `json:"id"`
	ClusterId    int64         `json:"clusterId"`
	AccessLogRef *AccessLogRef `json:"accessLogRef"`

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
