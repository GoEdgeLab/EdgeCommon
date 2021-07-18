package firewallconfigs

type HTTPFirewallCaptchaAction struct {
	Life     int32  `yaml:"life" json:"life"`
	Language string `yaml:"language" json:"language"` // 语言，zh-CN, en-US ... TODO 需要实现，目前是根据浏览器Accept-Language动态获取
}
