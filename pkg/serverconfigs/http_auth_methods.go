// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build !plus

package serverconfigs

type HTTPAuthType = string

const (
	HTTPAuthTypeBasicAuth  HTTPAuthType = "basicAuth"  // BasicAuth
	HTTPAuthTypeSubRequest HTTPAuthType = "subRequest" // 子请求
)

type HTTPAuthTypeDefinition struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func FindAllHTTPAuthTypes() []*HTTPAuthTypeDefinition {
	return []*HTTPAuthTypeDefinition{
		{
			Name:        "基本认证",
			Code:        HTTPAuthTypeBasicAuth,
			Description: "BasicAuth，最简单的HTTP请求认证方式，通过传递<span class=\"ui label tiny basic text\">Authorization: Basic xxx</span> Header认证。",
		},
		{
			Name:        "子请求",
			Code:        HTTPAuthTypeSubRequest,
			Description: "通过自定义的URL子请求来认证请求。",
		},
	}
}
