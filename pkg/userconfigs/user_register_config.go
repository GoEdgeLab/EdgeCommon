// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

const (
	EmailVerificationDefaultLife = 86400 * 2 // 2 days
)

type UserRegisterConfig struct {
	IsOn                bool `yaml:"isOn" json:"isOn"`                               // 是否启用用户注册
	ComplexPassword     bool `yaml:"complexPassword" json:"complexPassword"`         // 必须使用复杂密码
	RequireVerification bool `yaml:"requireVerification" json:"requireVerification"` // 是否需要审核
	RequireIdentity     bool `yaml:"requireIdentity" json:"requireIdentity"`         // 是否需要实名认证

	// 电子邮箱激活设置
	EmailVerification struct {
		IsOn       bool   `yaml:"isOn" json:"isOn"`             // 是否启用
		ShowNotice bool   `yaml:"showNotice" json:"showNotice"` // 提示用户未绑定
		Subject    string `yaml:"subject" json:"subject"`       // 标题
		Body       string `yaml:"body" json:"body"`             // 内容
		CanLogin   bool   `yaml:"canLogin" json:"canLogin"`     // 是否可以使用激活的邮箱登录
		Life       int32  `yaml:"life" json:"life"`             // 有效期
	} `yaml:"emailVerification" json:"emailVerification"`

	// CDN
	CDNIsOn   bool     `json:"cdnIsOn"`                    // 是否开启CDN服务
	ClusterId int64    `yaml:"clusterId" json:"clusterId"` // 用户创建服务集群
	Features  []string `yaml:"features" json:"features"`   // 默认启用的功能

	// 开通DNS
	NSIsOn bool `json:"nsIsOn"` // 是否开启智能DNS服务
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

	// 激活相关
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

	return config
}
