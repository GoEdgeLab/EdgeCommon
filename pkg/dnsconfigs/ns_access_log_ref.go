// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type NSAccessLogRef struct {
	IsPrior            bool `yaml:"isPrior" json:"isPrior"`                       // 是否覆盖
	IsOn               bool `yaml:"isOn" json:"isOn"`                             // 是否启用
	LogMissingDomains  bool `yaml:"logMissingDomains" json:"logMissingDomains"`   // 是否记录找不到的域名
	MissingRecordsOnly bool `yaml:"missingRecordsOnly" json:"missingRecordsOnly"` // 只记录找不到解析记录的访问
}

func (this *NSAccessLogRef) Init() error {
	return nil
}
