// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type HTTPDeflateCompressionConfig struct {
	Id   int64 `yaml:"id" json:"id"` // ID
	IsOn bool  `yaml:"isOn" json:"isOn"`

	Level     int8                           `yaml:"level" json:"level"`         // 级别
	MinLength *shared.SizeCapacity           `yaml:"minLength" json:"minLength"` // 最小压缩对象比如4m, 24k
	MaxLength *shared.SizeCapacity           `yaml:"maxLength" json:"maxLength"` // 最大压缩对象
	Conds     *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"`         // 匹配条件

	minLength int64
	maxLength int64
}

func (this *HTTPDeflateCompressionConfig) Init() error {
	if this.MinLength != nil {
		this.minLength = this.MinLength.Bytes()
	}
	if this.MaxLength != nil {
		this.maxLength = this.MaxLength.Bytes()
	}

	if this.Conds != nil {
		err := this.Conds.Init()
		if err != nil {
			return err
		}
	}

	return nil
}
