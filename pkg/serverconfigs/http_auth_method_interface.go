// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "net/http"

// HTTPAuthMethodInterface HTTP认证接口定义
type HTTPAuthMethodInterface interface {
	// Init 初始化
	Init(params map[string]any) error

	// MatchRequest 是否匹配请求
	MatchRequest(req *http.Request) bool

	// Filter 过滤
	Filter(req *http.Request, subReqFunc func(subReq *http.Request) (status int, err error), formatter func(string) string) (ok bool, newURI string, uriChanged bool, err error)

	// SetExts 设置扩展名
	SetExts(exts []string)

	// SetDomains 设置域名
	SetDomains(domains []string)
}
