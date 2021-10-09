// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type RecordTTL struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func FindAllRecordTTL() []*RecordTTL {
	return []*RecordTTL{
		{
			Name:  "5秒",
			Value: 5,
		},
		{
			Name:  "10秒",
			Value: 10,
		},
		{
			Name:  "30秒",
			Value: 30,
		},
		{
			Name:  "1分钟",
			Value: 60,
		},
		{
			Name:  "3分钟",
			Value: 3 * 60,
		},
		{
			Name:  "5分钟",
			Value: 5 * 60,
		},
		{
			Name:  "10分钟",
			Value: 10 * 60,
		},
		{
			Name:  "30分钟",
			Value: 30 * 60,
		},
		{
			Name:  "1小时",
			Value: 3600,
		},
		{
			Name:  "12小时",
			Value: 12 * 3600,
		},
		{
			Name:  "1天",
			Value: 86400,
		},
		{
			Name:  "30天",
			Value: 30 * 86400,
		},
		{
			Name:  "一年",
			Value: 365 * 86400,
		},
	}
}
