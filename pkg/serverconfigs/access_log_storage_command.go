// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

// AccessLogCommandStorageConfig 通过命令行存储
type AccessLogCommandStorageConfig struct {
	Command string   `yaml:"command" json:"command"`
	Args    []string `yaml:"args" json:"args"`
	Dir     string   `yaml:"dir" json:"dir"`
}
