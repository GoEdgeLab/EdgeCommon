// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package firewallconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type CaptchaType = string

const (
	CaptchaTypeDefault  CaptchaType = "default"
	CaptchaTypeOneClick CaptchaType = "oneClick"
	CaptchaTypeSlide    CaptchaType = "slide"
	CaptchaTypeGeeTest  CaptchaType = "geetest"
)

// FindAllCaptchaTypes Find all captcha types
func FindAllCaptchaTypes() []*shared.Definition {
	return []*shared.Definition{
		{
			Code:        CaptchaTypeDefault,
			Name:        "验证码",
			Description: "通过输入验证码来验证人机。",
		},
		{
			Code:        CaptchaTypeOneClick,
			Name:        "点击验证",
			Description: "通过点击界面元素来验证人机。",
		},
		{
			Code:        CaptchaTypeSlide,
			Name:        "滑动解锁",
			Description: "通过滑动方块解锁来验证人机。",
		},
		{
			Code:        CaptchaTypeGeeTest,
			Name:        "极验-行为验",
			Description: "使用极验-行为验提供的人机验证方式。",
		},
	}
}

func DefaultCaptchaType() *shared.Definition {
	var captchaTypes = FindAllCaptchaTypes()
	if len(captchaTypes) > 0 {
		return captchaTypes[0]
	}
	return &shared.Definition{
		Code: CaptchaTypeDefault,
		Name: "验证码",
	}
}

func FindCaptchaType(code CaptchaType) *shared.Definition {
	if len(code) == 0 {
		code = CaptchaTypeDefault
	}

	for _, t := range FindAllCaptchaTypes() {
		if t.Code == code {
			return t
		}
	}

	return DefaultCaptchaType()
}
