// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

// AccessLogESStorageConfig ElasticSearch存储策略
type AccessLogESStorageConfig struct {
	Endpoint     string `yaml:"endpoint" json:"endpoint"`
	Index        string `yaml:"index" json:"index"`
	MappingType  string `yaml:"mappingType" json:"mappingType"`
	Username     string `yaml:"username" json:"username"`
	Password     string `yaml:"password" json:"password"`
	IsDataStream bool   `yaml:"isDataStream" json:"isDataStream"` // 是否为Data Stream模式
}
