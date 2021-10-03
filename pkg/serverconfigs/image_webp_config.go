// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"strings"
)

// WebPImageConfig WebP配置
type WebPImageConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`
	IsOn    bool `yaml:"isOn" json:"isOn"`

	Quality int `yaml:"quality" json:"quality"` // 0-100

	MinLength  *shared.SizeCapacity           `yaml:"minLength" json:"minLength"`   // 最小压缩对象比如4m, 24k
	MaxLength  *shared.SizeCapacity           `yaml:"maxLength" json:"maxLength"`   // 最大压缩对象
	MimeTypes  []string                       `yaml:"mimeTypes" json:"mimeTypes"`   // 支持的MimeType，支持image/*这样的通配符使用
	Extensions []string                       `yaml:"extensions" json:"extensions"` // 文件扩展名，包含点符号，不区分大小写
	Conds      *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"`           // 匹配条件

	minLength     int64
	maxLength     int64
	mimeTypeRules []*shared.MimeTypeRule
	extensions    []string
}

func (this *WebPImageConfig) Init() error {
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

	// mime types
	this.mimeTypeRules = []*shared.MimeTypeRule{}
	for _, mimeType := range this.MimeTypes {
		rule, err := shared.NewMimeTypeRule(mimeType)
		if err != nil {
			return err
		}
		this.mimeTypeRules = append(this.mimeTypeRules, rule)
	}

	// extensions
	this.extensions = []string{}
	for _, ext := range this.Extensions {
		ext = strings.ToLower(ext)
		if len(ext) > 0 && ext[0] != '.' {
			ext = "." + ext
		}
		this.extensions = append(this.extensions, ext)
	}

	return nil
}

// MatchResponse 是否匹配响应
func (this *WebPImageConfig) MatchResponse(mimeType string, contentLength int64, requestExt string, formatter shared.Formatter) bool {
	if this.Conds != nil && formatter != nil {
		if !this.Conds.MatchRequest(formatter) {
			return false
		}
		if !this.Conds.MatchResponse(formatter) {
			return false
		}
	}

	// min length
	if this.minLength > 0 && contentLength < this.minLength {
		return false
	}

	// max length
	if this.maxLength > 0 && contentLength > this.maxLength {
		return false
	}

	// extensions
	if len(this.extensions) > 0 {
		if len(requestExt) > 0 {
			for _, ext := range this.extensions {
				if ext == requestExt {
					if strings.Contains(mimeType, "image/") {
						return true
					}
				}
			}
		}
	}

	// mime types
	if len(this.mimeTypeRules) > 0 {
		if len(mimeType) > 0 {
			var index = strings.Index(mimeType, ";")
			if index >= 0 {
				mimeType = mimeType[:index]
			}
			for _, rule := range this.mimeTypeRules {
				if rule.Match(mimeType) {
					return true
				}
			}
		}
	}

	// 如果没有指定条件，则所有的都能压缩
	if len(this.extensions) == 0 && len(this.mimeTypeRules) == 0 {
		return true
	}

	return false
}

// MatchRequest 是否匹配请求
func (this *WebPImageConfig) MatchRequest(requestExt string, formatter shared.Formatter) bool {
	if this.Conds != nil && formatter != nil {
		if !this.Conds.MatchRequest(formatter) {
			return false
		}
	}

	// extensions
	if len(this.mimeTypeRules) == 0 && len(this.extensions) > 0 && len(requestExt) > 0 {
		for _, ext := range this.extensions {
			if ext == requestExt {
				return true
			}
		}
		return false
	}

	return true
}

// MatchAccept 检查客户端是否能接受WebP
func (this *WebPImageConfig) MatchAccept(acceptContentTypes string) bool {
	var t = "image/webp"
	var index = strings.Index(acceptContentTypes, t)
	if index < 0 {
		return false
	}
	var l = len(acceptContentTypes)
	if index > 0 && acceptContentTypes[index-1] != ',' {
		return false
	}

	if index+len(t) < l && acceptContentTypes[index+len(t)] != ',' {
		return false
	}
	return true
}
