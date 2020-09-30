package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
)

// gzip配置
type HTTPGzipConfig struct {
	Id        int64                          `yaml:"id" json:"id"`               // ID
	IsOn      bool                           `yaml:"isOn" json:"isOn"`           // 是否启用
	Level     int8                           `yaml:"level" json:"level"`         // 1-9
	MinLength *shared.SizeCapacity           `yaml:"minLength" json:"minLength"` // 最小压缩对象比如4m, 24k
	MaxLength *shared.SizeCapacity           `yaml:"maxLength" json:"maxLength"` // 最大压缩对象
	Conds     *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"`         // 匹配条件

	minLength int64
	maxLength int64
	mimeTypes []*MimeTypeRule
}

// 校验
func (this *HTTPGzipConfig) Init() error {
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

// 可压缩最小尺寸
func (this *HTTPGzipConfig) MinBytes() int64 {
	return this.minLength
}

// 可压缩最大尺寸
func (this *HTTPGzipConfig) MaxBytes() int64 {
	return this.maxLength
}
