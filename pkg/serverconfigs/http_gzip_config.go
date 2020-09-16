package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"regexp"
	"strings"
)

// 默认的文件类型
var (
	DefaultGzipMimeTypes = []string{"text/html", "application/json"}
)

// gzip配置
type HTTPGzipConfig struct {
	Id        int64                `yaml:"id" json:"id"`               // ID
	IsOn      bool                 `yaml:"isOn" json:"isOn"`           // 是否启用
	Level     int8                 `yaml:"level" json:"level"`         // 1-9
	MinLength *shared.SizeCapacity `yaml:"minLength" json:"minLength"` // 最小压缩对象比如4m, 24k
	MaxLength *shared.SizeCapacity `yaml:"minLength" json:"maxLength"` // 最大压缩对象 TODO 需要实现
	MimeTypes []string             `yaml:"mimeTypes" json:"mimeTypes"` // 比如text/html, text/* // TODO 需要实现，可能需要用RequestConds替代

	minLength int64
	mimeTypes []*MimeTypeRule
}

// 校验
func (this *HTTPGzipConfig) Init() error {
	if this.MinLength != nil {
		this.minLength = this.MinLength.Bytes()
	}

	if len(this.MimeTypes) == 0 {
		this.MimeTypes = DefaultGzipMimeTypes
	}

	this.mimeTypes = []*MimeTypeRule{}
	for _, mimeType := range this.MimeTypes {
		if strings.Contains(mimeType, "*") {
			mimeType = regexp.QuoteMeta(mimeType)
			mimeType = strings.Replace(mimeType, "\\*", ".*", -1)
			reg, err := regexp.Compile("^" + mimeType + "$")
			if err != nil {
				return err
			}
			this.mimeTypes = append(this.mimeTypes, &MimeTypeRule{
				Value:  mimeType,
				Regexp: reg,
			})
		} else {
			this.mimeTypes = append(this.mimeTypes, &MimeTypeRule{
				Value:  mimeType,
				Regexp: nil,
			})
		}
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
