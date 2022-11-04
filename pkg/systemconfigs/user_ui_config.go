package systemconfigs

import (
	"github.com/iwind/TeaGo/types"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"time"
)

// UserUIConfig 用户界面相关配置
type UserUIConfig struct {
	ProductName        string `json:"productName"`        // 产品名
	UserSystemName     string `json:"userSystemName"`     // 管理员系统名称
	ShowOpenSourceInfo bool   `json:"showOpenSourceInfo"` // 是否显示开源信息
	ShowVersion        bool   `json:"showVersion"`        // 是否显示版本号
	Version            string `json:"version"`            // 显示的版本号
	ShowFinance        bool   `json:"showFinance"`        // 是否显示财务相关信息
	FaviconFileId      int64  `json:"faviconFileId"`      // Favicon文件ID
	LogoFileId         int64  `json:"logoFileId"`         // Logo文件ID
	TimeZone           string `json:"timeZone"`           // 时区

	BandwidthUnit       BandwidthUnit `json:"bandwidthUnit"`       // 带宽单位
	ShowTrafficCharts   bool          `json:"showTrafficCharts"`   // 是否显示流量相关图表和数据
	ShowBandwidthCharts bool          `json:"showBandwidthCharts"` // 是否显示带宽相关图表和数据

	TrafficStats struct {
		BandwidthPercentile       int32  `json:"bandwidthPercentile"`       // 带宽百分位
		DefaultBandwidthDateRange string `json:"defaultBandwidthDateRange"` // 默认带宽周期
	} `json:"trafficStats"` // 流量统计相关设置
}

func DefaultUserUIConfig() *UserUIConfig {
	var config = &UserUIConfig{
		ProductName:         "GoEdge",
		UserSystemName:      "GoEdge用户系统",
		ShowOpenSourceInfo:  true,
		ShowVersion:         true,
		ShowFinance:         true,
		BandwidthUnit:       BandwidthUnitBit,
		ShowBandwidthCharts: true,
		ShowTrafficCharts:   true,
	}
	config.TrafficStats.BandwidthPercentile = 95
	config.TrafficStats.DefaultBandwidthDateRange = "latest30days"
	return config
}

type BandwidthDateRange struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	DayFrom string `json:"dayFrom"`
	DayTo   string `json:"dayTo"`
}

const DefaultBandwidthDateRangeCode = "latest30days"
const DefaultBandwidthPercentile int32 = 95

func FindAllBandwidthDateRanges() []*BandwidthDateRange {
	var dayInWeek = types.Int(timeutil.Format("w"))
	if dayInWeek == 0 {
		dayInWeek = 7
	}

	return []*BandwidthDateRange{
		{
			Name:    "今天",
			Code:    "today",
			DayFrom: timeutil.Format("Y-m-d"),
			DayTo:   timeutil.Format("Y-m-d"),
		},
		{
			Name:    "昨天",
			Code:    "yesterday",
			DayFrom: timeutil.Format("Y-m-d", time.Now().AddDate(0, 0, -1)),
			DayTo:   timeutil.Format("Y-m-d", time.Now().AddDate(0, 0, -1)),
		},
		{
			Name:    "近7天",
			Code:    "latest7days",
			DayFrom: timeutil.Format("Y-m-d", time.Now().AddDate(0, 0, -6)),
			DayTo:   timeutil.Format("Y-m-d"),
		},
		{
			Name:    "上周",
			Code:    "lastWeek",
			DayFrom: timeutil.Format("Y-m-d", time.Now().AddDate(0, 0, -dayInWeek-6)),
			DayTo:   timeutil.Format("Y-m-d", time.Now().AddDate(0, 0, -dayInWeek)),
		},
		{
			Name:    "近30天",
			Code:    "latest30days",
			DayFrom: timeutil.Format("Y-m-d", time.Now().AddDate(0, 0, -29)),
			DayTo:   timeutil.Format("Y-m-d"),
		},
		{
			Name:    "当月",
			Code:    "currentMonth",
			DayFrom: timeutil.Format("Y-m-01"),
			DayTo:   timeutil.Format("Y-m-d"),
		},
		{
			Name:    "上月",
			Code:    "lastMonth",
			DayFrom: timeutil.Format("Y-m-01", time.Now().AddDate(0, -1, 0)),
			DayTo:   timeutil.Format("Y-m-t", time.Now().AddDate(0, -1, 0)),
		},
	}
}

func FindBandwidthDateRange(code string) *BandwidthDateRange {
	for _, r := range FindAllBandwidthDateRanges() {
		if r.Code == code {
			return r
		}
	}
	return nil
}
