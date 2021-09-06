// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package reporterconfigs

type TaskType = string

const (
	TaskTypeIPAddr TaskType = "ipAddr"
)

type IPTask struct {
	AddrId int64  `json:"addrId"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
}

func FindTaskTypeName(taskType TaskType) string {
	switch taskType {
	case TaskTypeIPAddr:
		return "IP地址"
	}
	return ""
}
