package systemconfigs

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

	BandwidthUnit       BandwidthUnit `json:"bandwidthUnit"`                                  // 带宽单位
	ShowTrafficCharts   bool          `json:"showTrafficCharts" json:"showTrafficCharts"`     // 是否显示流量相关图表和数据
	ShowBandwidthCharts bool          `json:"showBandwidthCharts" json:"showBandwidthCharts"` // 是否显示带宽相关图表和数据
}
