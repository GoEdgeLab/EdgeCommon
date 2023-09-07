// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"regexp"
	"strings"
)

type HTTPRemoteAddrType = string

const (
	HTTPRemoteAddrTypeDefault       HTTPRemoteAddrType = "default"       // 默认（直连）
	HTTPRemoteAddrTypeProxy         HTTPRemoteAddrType = "proxy"         // 代理
	HTTPRemoteAddrTypeRequestHeader HTTPRemoteAddrType = "requestHeader" // 请求报头
	HTTPRemoteAddrTypeVariable      HTTPRemoteAddrType = "variable"      // 变量
)

// HTTPRemoteAddrConfig HTTP获取客户端IP地址方式
type HTTPRemoteAddrConfig struct {
	IsPrior bool               `yaml:"isPrior" json:"isPrior"`
	IsOn    bool               `yaml:"isOn" json:"isOn"`
	Value   string             `yaml:"value" json:"value"` // 值变量
	Type    HTTPRemoteAddrType `yaml:"type" json:"type"`   // 类型

	RequestHeaderName string `yaml:"requestHeaderName" json:"requestHeaderName"` // 请求报头名称（type = requestHeader时生效）

	isEmpty bool
}

// Init 初始化
func (this *HTTPRemoteAddrConfig) Init() error {
	this.Value = strings.TrimSpace(this.Value)
	if len(this.Value) == 0 {
		this.isEmpty = true
	} else if regexp.MustCompile(`\s+`).ReplaceAllString(this.Value, "") == "${remoteAddr}" {
		this.isEmpty = true
	}

	return nil
}

// IsEmpty 是否为空
func (this *HTTPRemoteAddrConfig) IsEmpty() bool {
	return this.isEmpty
}
