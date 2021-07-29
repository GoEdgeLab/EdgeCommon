// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

// AccessLogFileStorageConfig 文件存储配置
type AccessLogFileStorageConfig struct {
	Path       string `yaml:"path" json:"path"`             // 文件路径，支持变量：${year|month|week|day|hour|minute|second}
	AutoCreate bool   `yaml:"autoCreate" json:"autoCreate"` // 是否自动创建目录
}
