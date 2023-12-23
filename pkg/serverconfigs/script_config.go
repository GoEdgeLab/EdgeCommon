// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	stringutil "github.com/iwind/TeaGo/utils/string"
	"strings"
)

type ScriptConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`
	IsOn    bool `yaml:"isOn" json:"isOn"`

	Code            string `yaml:"code" json:"code"`                       // 当前运行的代码
	AuditingCode    string `yaml:"auditingCode" json:"auditingCode"`       // 审核中的代码
	AuditingCodeMD5 string `yaml:"auditingCodeMD5" json:"auditingCodeMD5"` // 审核中的代码MD5

	realCode string
}

func (this *ScriptConfig) Init() error {
	this.realCode = this.TrimCode()

	return nil
}

func (this *ScriptConfig) TrimCode() string {
	return strings.TrimSpace(this.Code)
}

func (this *ScriptConfig) RealCode() string {
	return this.realCode
}

func (this *ScriptConfig) CodeMD5() string {
	return stringutil.Md5(this.TrimCode())
}
