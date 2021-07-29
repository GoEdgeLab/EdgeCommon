// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "github.com/iwind/TeaGo/maps"

type AccessLogSyslogStorageProtocol = string

const (
	AccessLogSyslogStorageProtocolTCP    AccessLogSyslogStorageProtocol = "tcp"
	AccessLogSyslogStorageProtocolUDP    AccessLogSyslogStorageProtocol = "udp"
	AccessLogSyslogStorageProtocolNone   AccessLogSyslogStorageProtocol = "none"
	AccessLogSyslogStorageProtocolSocket AccessLogSyslogStorageProtocol = "socket"
)

type AccessLogSyslogStoragePriority = int

const (
	AccessLogSyslogStoragePriorityEmerg AccessLogSyslogStoragePriority = iota
	AccessLogSyslogStoragePriorityAlert
	AccessLogSyslogStoragePriorityCrit
	AccessLogSyslogStoragePriorityErr
	AccessLogSyslogStoragePriorityWarning
	AccessLogSyslogStoragePriorityNotice
	AccessLogSyslogStoragePriorityInfo
	AccessLogSyslogStoragePriorityDebug
)

var AccessLogSyslogStoragePriorities = []maps.Map{
	{
		"name":  "[无]",
		"value": -1,
	},
	{
		"name":  "EMERG",
		"value": AccessLogSyslogStoragePriorityEmerg,
	},
	{
		"name":  "ALERT",
		"value": AccessLogSyslogStoragePriorityAlert,
	},
	{
		"name":  "CRIT",
		"value": AccessLogSyslogStoragePriorityCrit,
	},
	{
		"name":  "ERR",
		"value": AccessLogSyslogStoragePriorityErr,
	},
	{
		"name":  "WARNING",
		"value": AccessLogSyslogStoragePriorityWarning,
	},
	{
		"name":  "NOTICE",
		"value": AccessLogSyslogStoragePriorityNotice,
	},
	{
		"name":  "INFO",
		"value": AccessLogSyslogStoragePriorityInfo,
	},
	{
		"name":  "DEBUG",
		"value": AccessLogSyslogStoragePriorityDebug,
	},
}

// AccessLogSyslogStorageConfig syslog存储策略
type AccessLogSyslogStorageConfig struct {
	Protocol   string                         `yaml:"protocol" json:"protocol"` // SysLogStorageProtocol*
	ServerAddr string                         `yaml:"serverAddr" json:"serverAddr"`
	ServerPort int                            `yaml:"serverPort" json:"serverPort"`
	Socket     string                         `yaml:"socket" json:"socket"` // sock file
	Tag        string                         `yaml:"tag" json:"tag"`
	Priority   AccessLogSyslogStoragePriority `yaml:"priority" json:"priority"`
}

func FindAccessLogSyslogStoragePriorityName(priority AccessLogSyslogStoragePriority) string {
	for _, p := range AccessLogSyslogStoragePriorities {
		if p.GetInt("value") == priority {
			return p.GetString("name")
		}
	}
	return ""
}
