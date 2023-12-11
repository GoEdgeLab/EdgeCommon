// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

func init() {
	_ = DefaultWebPImagePolicy.Init()
}

var DefaultWebPImagePolicy = NewWebPImagePolicy()

func NewWebPImagePolicy() *WebPImagePolicy {
	return &WebPImagePolicy{
		IsOn:         true,
		RequireCache: true,
		MinLength:    shared.NewSizeCapacity(0, shared.SizeCapacityUnitKB),
		MaxLength:    shared.NewSizeCapacity(128, shared.SizeCapacityUnitMB),
	}
}

// WebPImagePolicy WebP策略
type WebPImagePolicy struct {
	IsOn         bool                 `yaml:"isOn" json:"isOn"`                 // 是否启用
	RequireCache bool                 `yaml:"requireCache" json:"requireCache"` // 需要在缓存条件下进行
	MinLength    *shared.SizeCapacity `yaml:"minLength" json:"minLength"`       // 最小压缩对象比如4m, 24k
	MaxLength    *shared.SizeCapacity `yaml:"maxLength" json:"maxLength"`       // 最大压缩对象
	Quality      int                  `yaml:"quality" json:"quality"`           // 生成的图片质量：0-100

	minLength int64
	maxLength int64
}

func (this *WebPImagePolicy) Init() error {
	if this.MinLength != nil {
		this.minLength = this.MinLength.Bytes()
	}
	if this.MaxLength != nil {
		this.maxLength = this.MaxLength.Bytes()
	}

	return nil
}

func (this *WebPImagePolicy) MinLengthBytes() int64 {
	return this.minLength
}

func (this *WebPImagePolicy) MaxLengthBytes() int64 {
	return this.maxLength
}
