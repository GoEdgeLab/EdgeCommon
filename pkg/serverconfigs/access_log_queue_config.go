// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

// AccessLogQueueConfig 访问日志队列配置
type AccessLogQueueConfig struct {
	MaxLength      int `yaml:"maxLength" json:"maxLength"`           // 队列最大长度
	CountPerSecond int `yaml:"countPerSecond" json:"countPerSecond"` // 每秒写入数量
	Percent        int `yaml:"percent" json:"percent"`               // 比例，如果为0-100，默认为100

	EnableAutoPartial bool  `yaml:"enableAutoPartial" json:"enableAutoPartial"` // 是否启动自动分表
	RowsPerTable      int64 `yaml:"rowsPerTable" json:"rowsPerTable"`           // 单表最大行数
}
