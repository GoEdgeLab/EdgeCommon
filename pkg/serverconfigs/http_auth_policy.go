// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"net/http"
)

// HTTPAuthPolicy HTTP认证策略
type HTTPAuthPolicy struct {
	Id     int64                  `json:"id"`
	Name   string                 `json:"name"`
	IsOn   bool                   `json:"isOn"`
	Type   HTTPAuthType           `json:"type"`
	Params map[string]interface{} `json:"params"`

	method HTTPAuthMethodInterface
}

// MatchRequest 检查是否匹配请求
func (this *HTTPAuthPolicy) MatchRequest(req *http.Request) bool {
	if this.method == nil {
		return false
	}
	return this.method.MatchRequest(req)
}

// Filter 过滤
func (this *HTTPAuthPolicy) Filter(req *http.Request, subReqFunc func(subReq *http.Request) (status int, err error), formatter func(string) string) (ok bool, newURI string, uriChanged bool, err error) {
	if this.method == nil {
		// 如果设置正确的方法，我们直接允许请求
		return true, "", false, nil
	}
	return this.method.Filter(req, subReqFunc, formatter)
}

// Method 获取认证实例
func (this *HTTPAuthPolicy) Method() HTTPAuthMethodInterface {
	return this.method
}
