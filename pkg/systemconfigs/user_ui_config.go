package systemconfigs

import (
	"github.com/iwind/TeaGo/types"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"time"
)

// UserUIConfig 用户界面相关配置
type UserUIConfig struct {
	ProductName    string `json:"productName"`    // 产品名
	UserSystemName string `json:"userSystemName"` // 管理员系统名称
	ShowPageFooter bool   `json:"showPageFooter"` // 是否显示页脚
	PageFooterHTML string `json:"pageFooterHTML"` // 页脚HTML
	ShowVersion    bool   `json:"showVersion"`    // 是否显示版本号
	Version        string `json:"version"`        // 显示的版本号
	ShowFinance    bool   `json:"showFinance"`    // 是否显示财务相关信息
	FaviconFileId  int64  `json:"faviconFileId"`  // Favicon文件ID
	LogoFileId     int64  `json:"logoFileId"`     // 控制面板Logo文件ID

	TimeZone string `json:"timeZone"` // 时区

	ClientIPHeaderNames string `json:"clientIPHeaderNames"` // 客户端IP获取报头名称列表

	Server struct {
		CheckCNAME bool `json:"checkCNAME"` // 是否检查CNAME
	} `json:"server"` // 服务相关设置

	BandwidthUnit                BandwidthUnit `json:"bandwidthUnit"`                // 带宽单位
	ShowTrafficCharts            bool          `json:"showTrafficCharts"`            // 是否显示流量相关图表和数据
	ShowCacheInfoInTrafficCharts bool          `json:"showCacheInfoInTrafficCharts"` // 在流量图中显示缓存相关信息
	ShowBandwidthCharts          bool          `json:"showBandwidthCharts"`          // 是否显示带宽相关图表和数据

	TrafficStats struct {
		BandwidthPercentile       int32         `json:"bandwidthPercentile"`       // 带宽百分位
		DefaultBandwidthDateRange string        `json:"defaultBandwidthDateRange"` // 默认带宽周期
		BandwidthAlgo             BandwidthAlgo `json:"bandwidthAlgo"`             // 带宽算法
	} `json:"trafficStats"` // 流量统计相关设置

	Portal struct {
		IsOn       bool  `json:"isOn"`       // 是否启用
		LogoFileId int64 `json:"logoFileId"` // Logo文件ID
	} `json:"portal"` // 门户页面相关设置

	Theme ThemeConfig `yaml:"theme" json:"theme"` // 风格模板
}

func NewUserUIConfig() *UserUIConfig {
	var config = &UserUIConfig{
		ProductName:         "GoEdge",
		UserSystemName:      "GoEdge用户系统",
		ShowPageFooter:      false,
		ShowVersion:         true,
		ShowFinance:         true,
		BandwidthUnit:       BandwidthUnitBit,
		ShowBandwidthCharts: true,
		ShowTrafficCharts:   true,
		TimeZone:            "Asia/Shanghai",
	}

	// 服务相关
	config.Server.CheckCNAME = true

	// 流量相关
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
