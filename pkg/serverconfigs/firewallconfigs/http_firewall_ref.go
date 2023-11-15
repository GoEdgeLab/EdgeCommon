package firewallconfigs

type ServerCaptchaType = string

const (
	ServerCaptchaTypeNone     ServerCaptchaType = "none" // 不设置表示策略整体配置
	ServerCaptchaTypeDefault  ServerCaptchaType = CaptchaTypeDefault
	ServerCaptchaTypeOneClick ServerCaptchaType = CaptchaTypeOneClick
	ServerCaptchaTypeSlide    ServerCaptchaType = CaptchaTypeSlide
)

type HTTPFirewallRef struct {
	IsPrior           bool  `yaml:"isPrior" json:"isPrior"`
	IsOn              bool  `yaml:"isOn" json:"isOn"`
	FirewallPolicyId  int64 `yaml:"firewallPolicyId" json:"firewallPolicyId"`
	IgnoreGlobalRules bool  `yaml:"ignoreGlobalRules" json:"ignoreGlobalRules"` // 忽略系统定义的全局规则

	DefaultCaptchaType ServerCaptchaType `yaml:"defaultCaptchaType" json:"defaultCaptchaType"` // 默认人机识别方式
}

func (this *HTTPFirewallRef) Init() error {
	return nil
}
