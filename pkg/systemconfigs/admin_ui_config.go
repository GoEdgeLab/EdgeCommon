package systemconfigs

// AdminUIConfig 管理员界面相关配置
type AdminUIConfig struct {
	ProductName        string   `json:"productName"`        // 产品名
	AdminSystemName    string   `json:"adminSystemName"`    // 管理员系统名称
	ShowOpenSourceInfo bool     `json:"showOpenSourceInfo"` // 是否显示开源信息
	ShowVersion        bool     `json:"showVersion"`        // 是否显示版本号
	Version            string   `json:"version"`            // 显示的版本号
	ShowFinance        bool     `json:"showFinance"`        // 是否显示财务相关信息
	FaviconFileId      int64    `json:"faviconFileId"`      // Favicon文件ID
	LogoFileId         int64    `json:"logoFileId"`         // Logo文件ID
	DefaultPageSize    int      `json:"defaultPageSize"`    // 默认每页显示数
	TimeZone           string   `json:"timeZone"`           // 时区
	Modules            []string `json:"modules"`            // 开通模块
	DNSResolver        struct {
		Type string `json:"type"` // 类型，参考 DNSResolverType*
	} `json:"dnsResolver"` // DNS解析设置
}

func (this *AdminUIConfig) ContainsModule(module string) bool {
	if len(this.Modules) == 0 {
		return true
	}
	for _, m := range this.Modules {
		if m == module {
			return true
		}
	}
	return false
}
