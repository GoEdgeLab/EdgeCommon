package firewallconfigs

type HTTPFirewallCaptchaAction struct {
	Life             int32 `yaml:"life" json:"life"`                         // 有效期
	MaxFails         int   `yaml:"maxFails" json:"maxFails"`                 // 最大失败次数
	FailBlockTimeout int   `yaml:"failBlockTimeout" json:"failBlockTimeout"` // 失败拦截时间

	UIIsOn          bool   `yaml:"uiIsOn" json:"uiIsOn"`                   // 是否使用自定义UI TODO
	UITitle         string `yaml:"uiTitle" json:"uiTitle"`                 // 消息标题 TODO
	UIPrompt        string `yaml:"uiPrompt" json:"uiPrompt"`               // 消息提示 TODO
	UIButtonTitle   string `yaml:"uiButtonTitle" json:"uiButtonTitle"`     // 按钮标题 TODO
	UIShowRequestId bool   `yaml:"uiShowRequestId" json:"uiShowRequestId"` // 是否显示请求ID TODO
	UICss           string `yaml:"uiCss" json:"uiCss"`                     // CSS样式 TODO
	UIFooter        string `yaml:"uiFooter" json:"uiFooter"`               // TODO

	CookieId string `yaml:"cookieId" json:"cookieId"` // TODO

	Language string `yaml:"language" json:"language"` // 语言，zh-CN, en-US ... TODO 需要实现，目前是根据浏览器Accept-Language动态获取
}
