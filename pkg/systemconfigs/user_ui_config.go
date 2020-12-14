package systemconfigs

// 用户界面相关配置
type UserUIConfig struct {
	ProductName        string `json:"productName"`        // 产品名
	UserSystemName     string `json:"userSystemName"`     // 管理员系统名称
	ShowOpenSourceInfo bool   `json:"showOpenSourceInfo"` // 是否显示开源信息
	ShowVersion        bool   `json:"showVersion"`        // 是否显示版本号
	Version            string `json:"version"`            // 显示的版本号
}
