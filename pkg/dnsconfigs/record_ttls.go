// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type RecordTTL struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func FindAllRecordTTL() []*RecordTTL {
	return []*RecordTTL{
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
	}
}
