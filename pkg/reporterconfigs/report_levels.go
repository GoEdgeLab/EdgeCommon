// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package reporterconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
)

type ReportLevel = string

const (
	ReportLevelGood   ReportLevel = "good"
	ReportLevelNormal ReportLevel = "normal"
	ReportLevelBad    ReportLevel = "bad"
	ReportLevelBroken ReportLevel = "broken"
)

func FindAllReportLevels() []*shared.Definition {
	return []*shared.Definition{
		{
			Name: "良好",
			Code: ReportLevelGood,
		},
		{
			Name: "正常",
			Code: ReportLevelNormal,
		},
		{
			Name: "不良",
			Code: ReportLevelBad,
		},
		{
			Name: "错误",
			Code: ReportLevelBroken,
		},
	}
}

func FindReportLevelName(level ReportLevel) string {
	for _, def := range FindAllReportLevels() {
		if def.Code == level {
			return def.Name
		}
	}
	return ""
}
