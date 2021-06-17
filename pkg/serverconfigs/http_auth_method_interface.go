// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "net/http"

// HTTPAuthMethodInterface HTTP认证接口定义
type HTTPAuthMethodInterface interface {
	// Init 初始化
	Init(params map[string]interface{}) error

	// Filter 过滤
	Filter(req *http.Request, subReqFunc func(subReq *http.Request) (status int, err error), formatter func(string) string) (bool, error)
}
