// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

const (
	EmailVerificationDefaultLife  = 86400 * 2 // 2 days
	EmailResetPasswordDefaultLife = 3600      // 1 hour

	MobileVerificationDefaultLife  = 1800 // 30 minutes
	MobileResetPasswordDefaultLife = 1800 // 30 minutes
)

type UserRegisterConfig struct {
	IsOn                bool `yaml:"isOn" json:"isOn"`                               // 是否启用用户注册
	ComplexPassword     bool `yaml:"complexPassword" json:"complexPassword"`         // 必须使用复杂密码
	RequireVerification bool `yaml:"requireVerification" json:"requireVerification"` // 是否需要审核
	RequireIdentity     bool `yaml:"requireIdentity" json:"requireIdentity"`         // 是否需要实名认证
	CheckClientRegion   bool `yaml:"checkClientRegion" json:"checkClientRegion"`     // 在登录状态下检查客户端区域

	// 电子邮箱激活设置
	EmailVerification struct {
		IsOn       bool   `yaml:"isOn" json:"isOn"`             // 是否启用
		ShowNotice bool   `yaml:"showNotice" json:"showNotice"` // 提示用户未绑定
		Subject    string `yaml:"subject" json:"subject"`       // 标题
		Body       string `yaml:"body" json:"body"`             // 内容
		CanLogin   bool   `yaml:"canLogin" json:"canLogin"`     // 是否可以使用激活的邮箱登录
		Life       int32  `yaml:"life" json:"life"`             // 有效期
	} `yaml:"emailVerification" json:"emailVerification"`

	// 通过邮件找回密码设置
	EmailResetPassword struct {
		IsOn    bool   `yaml:"isOn" json:"isOn"`       // 是否启用
		Subject string `yaml:"subject" json:"subject"` // 标题
		Body    string `yaml:"body" json:"body"`       // 内容
		Life    int32  `yaml:"life" json:"life"`       // 有效期
	} `yaml:"emailResetPassword" json:"emailResetPassword"`

	// 手机号码激活设置
	MobileVerification struct {
		IsOn       bool   `yaml:"isOn" json:"isOn"`             // 是否启用
		ShowNotice bool   `yaml:"showNotice" json:"showNotice"` // 提示用户未绑定
		CanLogin   bool   `yaml:"canLogin" json:"canLogin"`     // 是否可以使用激活的邮箱登录
		Body       string `yaml:"body" json:"body"`             // 内容
		Life       int32  `yaml:"life" json:"life"`             // 有效期
	} `yaml:"mobileVerification" json:"mobileVerification"`

	// CDN
	CDNIsOn   bool     `json:"cdnIsOn"`                    // 是否开启CDN服务
	ClusterId int64    `yaml:"clusterId" json:"clusterId"` // 用户创建服务集群
	Features  []string `yaml:"features" json:"features"`   // 默认启用的功能

	// 开通DNS服务
	NSIsOn bool `json:"nsIsOn"` // 是否开启智能DNS服务

	// 开通高防服务
	ADIsOn bool `json:"adIsOn"` // 是否开启高防服务
}

func DefaultUserRegisterConfig() *UserRegisterConfig {
	var config = &UserRegisterConfig{
		IsOn:            false,
		ComplexPassword: true,
		CDNIsOn:         true,
		NSIsOn:          false,
		Features: []string{
			UserFeatureCodeServerAccessLog,
			UserFeatureCodeServerViewAccessLog,
			UserFeatureCodeServerWAF,
			UserFeatureCodePlan,
		},
		RequireVerification: false,
	}

	// 邮箱激活相关
	config.EmailVerification.CanLogin = true
	config.EmailVerification.ShowNotice = true
	config.EmailVerification.Subject = "【${product.name}】Email地址激活"
	config.EmailVerification.Body = `<p>欢迎你使用 ${product.name} 提供的服务，你需要点击以下链接激活你的Email邮箱：</p>
<p><a href="${url.verify}" target="_blank">${url.verify}</a></p>
<p>如果上面内容不是链接形式，请将该地址手工粘贴到浏览器地址栏再访问。</p>
<p></p>
<p>此致</p>
<p>${product.name} 管理团队</p>
<p><a href="${url.home}" target="_blank">${url.home}</a></p>
`

	// 通过邮件重置密码相关
	config.EmailResetPassword.IsOn = true
	config.EmailResetPassword.Subject = "【${product.name}】找回密码"
	config.EmailResetPassword.Body = `<p>你正在使用 ${product.name} 提供的找回密码功能，你需要将以下的数字验证码输入到找回密码页面中：</p>
<p><strong>验证码：${code}</strong></p>
<p></p>
<p>${product.name} 管理团队</p>
<p><a href="${url.home}" target="_blank">${url.home}</a></p>
`
	// 短信验证码
	config.MobileVerification.Body = "你的短信验证码${code}"

	return config
}
