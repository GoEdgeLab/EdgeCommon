// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"errors"
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

// Init 初始化
func (this *HTTPAuthPolicy) Init() error {
	switch this.Type {
	case HTTPAuthTypeBasicAuth:
		this.method = NewHTTPAuthBasicMethod()
	case HTTPAuthTypeSubRequest:
		this.method = NewHTTPAuthSubRequestMethod()
	}

	if this.method == nil {
		return errors.New("unknown auth method '" + this.Type + "'")
	}
	err := this.method.Init(this.Params)
	if err != nil {
		return err
	}

	return nil
}

// Filter 过滤
func (this *HTTPAuthPolicy) Filter(req *http.Request, subReqFunc func(subReq *http.Request) (status int, err error), formatter func(string) string) (bool, error) {
	if this.method == nil {
		// 如果设置正确的方法，我们直接允许请求
		return true, nil
	}
	return this.method.Filter(req, subReqFunc, formatter)
}

// Method 获取认证实例
func (this *HTTPAuthPolicy) Method() HTTPAuthMethodInterface {
	return this.method
}
