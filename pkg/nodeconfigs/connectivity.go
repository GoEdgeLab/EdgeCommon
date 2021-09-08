package nodeconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/reporterconfigs"

// Connectivity 连通性状态
type Connectivity struct {
	CostMs    float64                     `json:"costMs"`    // 平均耗时
	Level     reporterconfigs.ReportLevel `json:"level"`     // 级别
	Percent   float64                     `json:"percent"`   // 连通的百分比，是一个0到100之间的小数
	UpdatedAt int64                       `json:"updatedAt"` // 更新时间
}
