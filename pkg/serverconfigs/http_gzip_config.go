package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"strings"
)

// 默认的文件类型
var (
	DefaultGzipMimeTypes = []string{"text/html", "application/json"}
)

// gzip配置
type HTTPGzipConfig struct {
	Id         int64                          `yaml:"id" json:"id"`                 // ID
	IsOn       bool                           `yaml:"isOn" json:"isOn"`             // 是否启用
	Level      int8                           `yaml:"level" json:"level"`           // 1-9
	MinLength  *shared.SizeCapacity           `yaml:"minLength" json:"minLength"`   // 最小压缩对象比如4m, 24k
	MaxLength  *shared.SizeCapacity           `yaml:"maxLength" json:"maxLength"`   // 最大压缩对象 TODO 需要实现
	CondGroups []*shared.HTTPRequestCondGroup `yaml:"condGroups" json:"condGroups"` // 匹配条件

	minLength int64
	mimeTypes []*MimeTypeRule
}

// 校验
func (this *HTTPGzipConfig) Init() error {
	if this.MinLength != nil {
		this.minLength = this.MinLength.Bytes()
	}

	return nil
}

// 可压缩最小尺寸
func (this *HTTPGzipConfig) MinBytes() int64 {
	return this.minLength
}

// 检查是否匹配Content-Type
func (this *HTTPGzipConfig) MatchContentType(contentType string) bool {
	index := strings.Index(contentType, ";")
	if index >= 0 {
		contentType = contentType[:index]
	}
	for _, mimeType := range this.mimeTypes {
		if mimeType.Regexp == nil && contentType == mimeType.Value {
			return true
		} else if mimeType.Regexp != nil && mimeType.Regexp.MatchString(contentType) {
			return true
		}
	}
	return false
}
