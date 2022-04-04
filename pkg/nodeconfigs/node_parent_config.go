// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

// ParentNodeConfig 父级节点配置
type ParentNodeConfig struct {
	Id         int64    `yaml:"id" json:"id"`
	Addrs      []string `yaml:"addrs" json:"addrs"`
	SecretHash string   `yaml:"secretHash" json:"secretHash"`
}
