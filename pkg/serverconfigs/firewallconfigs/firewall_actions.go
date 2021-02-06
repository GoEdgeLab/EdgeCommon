package firewallconfigs

type FirewallActionType = string

const (
	FirewallActionTypeIPSet     FirewallActionType = "ipset"
	FirewallActionTypeFirewalld FirewallActionType = "firewalld"
	FirewallActionTypeIPTables  FirewallActionType = "iptables"
	FirewallActionTypeScript    FirewallActionType = "script"
	FirewallActionTypeHTTPAPI   FirewallActionType = "httpAPI"
)

type FirewallActionTypeDefinition struct {
	Name        string             `json:"name"`
	Code        FirewallActionType `json:"code"`
	Description string             `json:"description"`
}

func FindAllFirewallActionTypes() []*FirewallActionTypeDefinition {
	return []*FirewallActionTypeDefinition{
		{
			Name:        "ipset",
			Code:        FirewallActionTypeIPSet,
			Description: "使用特定的ipset管理IP，可以结合iptables和firewalld等工具一起工作。",
		},
		{
			Name:        "firewalld",
			Code:        FirewallActionTypeFirewalld,
			Description: "使用Firewalld管理IP，非持久保存，reload之后重置规则。",
		},
		{
			Name:        "iptables",
			Code:        FirewallActionTypeIPTables,
			Description: "使用IPTables管理IP，不支持超时时间设定，非持久保存，reload之后重置规则。",
		},
		{
			Name:        "自定义脚本",
			Code:        FirewallActionTypeScript,
			Description: "使用自定义的脚本执行IP操作。",
		},
		{
			Name:        "自定义HTTP API",
			Code:        FirewallActionTypeHTTPAPI,
			Description: "使用自定义的HTTP API执行IP操作。",
		},
	}
}

func FindFirewallActionTypeName(actionType FirewallActionType) string {
	for _, a := range FindAllFirewallActionTypes() {
		if a.Code == actionType {
			return a.Name
		}
	}
	return ""
}

type FirewallActionIPSetConfig struct {
	Path               string `json:"path"`               // 命令路径 TODO 暂时不实现
	WhiteName          string `json:"whiteName"`          // IPSet白名单名称
	BlackName          string `json:"blackName"`          // IPSet黑名单名称
	MaxElements        int    `json:"maxElements"`        // 最多IP数量 TODO 暂时不实现
	AutoAddToIPTables  bool   `json:"autoAddToIPTables"`  // 是否自动创建IPTables规则
	AutoAddToFirewalld bool   `json:"autoAddToFirewalld"` // 是否自动加入到Firewalld

	// TODO 添加需要阻止的端口列表
}

type FirewallActionFirewalldConfig struct {
	Path string `json:"path"` // 命令路径 TODO 暂时不实现

	// TODO 添加需要阻止的端口列表
}

type FirewallActionIPTablesConfig struct {
	Path string `json:"path"` // 命令路径 TODO 暂时不实现

	// TODO 添加需要阻止的端口列表
}

type FirewallActionScriptConfig struct {
	Path string   `json:"path"` // 脚本路径
	Cwd  string   `json:"cwd"`  // 工作目录 TODO 暂时不实现
	Args []string `json:"args"` // 附加参数 TODO 暂时不实现

	// TODO 添加需要阻止的端口列表
}

type FirewallActionHTTPAPIConfig struct {
	URL            string `json:"url"`            // URL路径
	TimeoutSeconds int    `json:"timeoutSeconds"` // 超时时间 TODO 暂时不实现
	Secret         string `json:"secret"`         // 认证密钥 TODO 暂时不实现

	// TODO 添加需要阻止的端口列表
}
