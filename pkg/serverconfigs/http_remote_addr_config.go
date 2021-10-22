// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"regexp"
	"strings"
)

// HTTPRemoteAddrConfig HTTP获取客户端IP地址方式
type HTTPRemoteAddrConfig struct {
	IsPrior      bool   `yaml:"isPrior" json:"isPrior"`
	IsOn         bool   `yaml:"isOn" json:"isOn"`
	Value        string `yaml:"value" json:"value"`               // 值变量
	IsCustomized bool   `yaml:"isCustomized" json:"isCustomized"` // 是否自定义

	isEmpty bool
}

// Init 初始化
func (this *HTTPRemoteAddrConfig) Init() error {
	if len(this.Value) == 0 {
		this.isEmpty = true
	} else if regexp.MustCompile(`\s+`).ReplaceAllString(this.Value, "") == "${remoteAddr}" {
		this.isEmpty = true
	}

	this.Value = strings.ReplaceAll(this.Value, "${remoteAddr}", "${remoteAddrValue}")

	return nil
}

// IsEmpty 是否为空
func (this *HTTPRemoteAddrConfig) IsEmpty() bool {
	return this.isEmpty
}
