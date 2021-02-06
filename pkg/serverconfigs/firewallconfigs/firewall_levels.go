package firewallconfigs

type FirewallEventLevelDefinition struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func FindAllFirewallEventLevels() []*FirewallEventLevelDefinition {
	return []*FirewallEventLevelDefinition{
		{
			Name:        "调试",
			Code:        "debug",
			Description: "仅作为调试用途",
		},
		{
			Name:        "通知",
			Code:        "notice",
			Description: "需要通知的事件",
		},
		{
			Name:        "警告",
			Code:        "warning",
			Description: "需要警告的事件",
		},
		{
			Name:        "错误",
			Code:        "error",
			Description: "发生系统错误的事件",
		},
		{
			Name:        "严重",
			Code:        "critical",
			Description: "性质较为严重的事件",
		},
		{
			Name:        "致命",
			Code:        "fatal",
			Description: "对系统有重大影响的事件",
		},
	}
}

func FindFirewallEventLevelName(code string) string {
	for _, level := range FindAllFirewallEventLevels() {
		if level.Code == code {
			return level.Name
		}
	}
	return ""
}
