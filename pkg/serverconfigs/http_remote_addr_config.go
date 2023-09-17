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

	isEmpty   bool
	values    []string
	hasValues bool
}

// Init 初始化
func (this *HTTPRemoteAddrConfig) Init() error {
	this.Value = strings.TrimSpace(this.Value)
	this.isEmpty = false
	if len(this.Value) == 0 {
		this.isEmpty = true
	} else if regexp.MustCompile(`\s+`).ReplaceAllString(this.Value, "") == "${remoteAddr}" {
		this.isEmpty = true
	}

	// values
	this.values = []string{}
	var headerVarReg = regexp.MustCompile(`(\$\{header\.)([\w-,]+)(})`)
	if headerVarReg.MatchString(this.Value) {
		var subMatches = headerVarReg.FindStringSubmatch(this.Value)
		if len(subMatches) > 3 {
			var prefix = subMatches[1]
			var headerNamesString = subMatches[2]
			var suffix = subMatches[3]
			for _, headerName := range strings.Split(headerNamesString, ",") {
				headerName = strings.TrimSpace(headerName)
				if len(headerName) > 0 {
					this.values = append(this.values, prefix+headerName+suffix)
				}
			}
		}
	}
	this.hasValues = len(this.values) > 1 // MUST be 1, not 0

	return nil
}

// IsEmpty 是否为空
func (this *HTTPRemoteAddrConfig) IsEmpty() bool {
	return this.isEmpty
}

// Values 可能的值变量
func (this *HTTPRemoteAddrConfig) Values() []string {
	return this.values
}

// HasValues 检查是否有一组值
func (this *HTTPRemoteAddrConfig) HasValues() bool {
	return this.hasValues
}
