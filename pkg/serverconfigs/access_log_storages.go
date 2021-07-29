package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
)

// AccessLogStorageType 访问日志存储类型
type AccessLogStorageType = string

const (
	AccessLogStorageTypeFile    AccessLogStorageType = "file"
	AccessLogStorageTypeES      AccessLogStorageType = "es"
	AccessLogStorageTypeTCP     AccessLogStorageType = "tcp"
	AccessLogStorageTypeSyslog  AccessLogStorageType = "syslog"
	AccessLogStorageTypeCommand AccessLogStorageType = "command"
)

// FindAllAccessLogStorageTypes 所有存储引擎列表
func FindAllAccessLogStorageTypes() []*shared.Definition {
	return []*shared.Definition{
		{
			Name:        "文件",
			Code:        AccessLogStorageTypeFile,
			Description: "将日志存储在磁盘文件中",
		},
		{
			Name:        "ElasticSearch",
			Code:        AccessLogStorageTypeES,
			Description: "将日志存储在ElasticSearch中",
		},
		{
			Name:        "TCP Socket",
			Code:        AccessLogStorageTypeTCP,
			Description: "将日志通过TCP套接字输出",
		},
		{
			Name:        "Syslog",
			Code:        AccessLogStorageTypeSyslog,
			Description: "将日志通过syslog输出，仅支持Linux",
		},
		{
			Name:        "命令行输入流",
			Code:        AccessLogStorageTypeCommand,
			Description: "启动一个命令通过读取stdin接收日志信息",
		},
	}
}

// FindAccessLogStorageTypeName 根据类型查找名称
func FindAccessLogStorageTypeName(storageType string) string {
	for _, m := range FindAllAccessLogStorageTypes() {
		if m.Code == storageType {
			return m.Name
		}
	}
	return ""
}
