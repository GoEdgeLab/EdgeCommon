package systemconfigs

// 管理员界面相关配置
type AdminUIConfig struct {
	ProductName        string `json:"productName"`        // 产品名
	AdminSystemName    string `json:"adminSystemName"`    // 管理员系统名称
	ShowOpenSourceInfo bool   `json:"showOpenSourceInfo"` // 是否显示开源信息
	ShowVersion        bool   `json:"showVersion"`        // 是否显示版本号
	Version            string `json:"version"`            // 显示的版本号
}
