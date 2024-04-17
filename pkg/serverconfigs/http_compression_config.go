// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/lists"
	"strings"
)

var DefaultHTTPCompressionTypes = []HTTPCompressionType{HTTPCompressionTypeBrotli, HTTPCompressionTypeGzip, HTTPCompressionTypeDeflate}

type HTTPCompressionRef struct {
	Id   int64 `yaml:"id" json:"id"`
	IsOn bool  `yaml:"isOn" json:"isOn"`
}

// HTTPCompressionConfig 内容压缩配置
type HTTPCompressionConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`
	IsOn    bool `yaml:"isOn" json:"isOn"`

	UseDefaultTypes bool                  `yaml:"useDefaultTypes" json:"useDefaultTypes"` // 是否使用默认的类型
	Types           []HTTPCompressionType `yaml:"types" json:"types"`                     // 支持的类型，如果为空表示默认顺序
	Level           int8                  `yaml:"level" json:"level"`                     // 级别：1-12
	DecompressData  bool                  `yaml:"decompressData" json:"decompressData"`   // 是否解压已压缩内容

	GzipRef    *HTTPCompressionRef `yaml:"gzipRef" json:"gzipRef"`
	DeflateRef *HTTPCompressionRef `yaml:"deflateRef" json:"deflateRef"`
	BrotliRef  *HTTPCompressionRef `yaml:"brotliRef" json:"brotliRef"`

	Gzip    *HTTPGzipCompressionConfig    `yaml:"gzip" json:"gzip"`
	Deflate *HTTPDeflateCompressionConfig `yaml:"deflate" json:"deflate"`
	Brotli  *HTTPBrotliCompressionConfig  `yaml:"brotli" json:"brotli"`

	MinLength            *shared.SizeCapacity           `yaml:"minLength" json:"minLength"`                       // 最小压缩对象比如4m, 24k
	MaxLength            *shared.SizeCapacity           `yaml:"maxLength" json:"maxLength"`                       // 最大压缩对象
	MimeTypes            []string                       `yaml:"mimeTypes" json:"mimeTypes"`                       // 支持的MimeType，支持image/*这样的通配符使用
	Extensions           []string                       `yaml:"extensions" json:"extensions"`                     // 文件扩展名，包含点符号，不区分大小写
	ExceptExtensions     []string                       `yaml:"exceptExtensions" json:"exceptExtensions"`         // 例外扩展名
	Conds                *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"`                               // 匹配条件
	EnablePartialContent bool                           `yaml:"enablePartialContent" json:"enablePartialContent"` // 支持PartialContent压缩

	OnlyURLPatterns   []*shared.URLPattern `yaml:"onlyURLPatterns" json:"onlyURLPatterns"`     // 仅限的URL
	ExceptURLPatterns []*shared.URLPattern `yaml:"exceptURLPatterns" json:"exceptURLPatterns"` // 排除的URL

	minLength        int64
	maxLength        int64
	mimeTypeRules    []*shared.MimeTypeRule
	extensions       []string
	exceptExtensions []string

	types []HTTPCompressionType

	supportGzip    bool
	supportDeflate bool
	supportBrotli  bool
	supportZSTD    bool
}

// Init 初始化
func (this *HTTPCompressionConfig) Init() error {
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

	this.exceptExtensions = []string{}
	for _, ext := range this.ExceptExtensions {
		ext = strings.ToLower(ext)
		if len(ext) > 0 && ext[0] != '.' {
			ext = "." + ext
		}
		this.exceptExtensions = append(this.exceptExtensions, ext)
	}

	if this.Gzip != nil {
		err := this.Gzip.Init()
		if err != nil {
			return err
		}
	}

	if this.Deflate != nil {
		err := this.Deflate.Init()
		if err != nil {
			return err
		}
	}

	if this.Brotli != nil {
		err := this.Brotli.Init()
		if err != nil {
			return err
		}
	}

	var supportedTypes = []HTTPCompressionType{}
	if !this.UseDefaultTypes {
		supportedTypes = append(supportedTypes, this.Types...)
	} else {
		supportedTypes = append(supportedTypes, DefaultHTTPCompressionTypes...)
	}
	this.types = supportedTypes

	this.supportGzip = false
	this.supportDeflate = false
	this.supportDeflate = false
	for _, supportType := range supportedTypes {
		switch supportType {
		case HTTPCompressionTypeGzip:
			if this.GzipRef == nil || (this.GzipRef != nil && this.GzipRef.IsOn && this.Gzip != nil && this.Gzip.IsOn) {
				this.supportGzip = true
			}
		case HTTPCompressionTypeDeflate:
			if this.DeflateRef == nil || (this.DeflateRef != nil && this.DeflateRef.IsOn && this.Deflate != nil && this.Deflate.IsOn) {
				this.supportDeflate = true
			}
		case HTTPCompressionTypeBrotli:
			if this.BrotliRef == nil || (this.BrotliRef != nil && this.BrotliRef.IsOn && this.Brotli != nil && this.Brotli.IsOn) {
				this.supportBrotli = true
			}
		case HTTPCompressionTypeZSTD:
			this.supportZSTD = true
		}
	}

	// url patterns
	for _, pattern := range this.ExceptURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	for _, pattern := range this.OnlyURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// MinBytes 可压缩最小尺寸
func (this *HTTPCompressionConfig) MinBytes() int64 {
	return this.minLength
}

// MaxBytes 可压缩最大尺寸
func (this *HTTPCompressionConfig) MaxBytes() int64 {
	return this.maxLength
}

// MatchResponse 是否匹配响应
func (this *HTTPCompressionConfig) MatchResponse(mimeType string, contentLength int64, requestExt string, formatter shared.Formatter) bool {
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

	// except extensions
	if len(this.exceptExtensions) > 0 {
		if len(requestExt) > 0 {
			for _, ext := range this.exceptExtensions {
				if ext == requestExt {
					return false
				}
			}
		}
	}

	// extensions
	if len(this.extensions) > 0 {
		if len(requestExt) > 0 {
			for _, ext := range this.extensions {
				if ext == requestExt {
					return true
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

// MatchAcceptEncoding 根据Accept-Encoding选择优先的压缩方式
func (this *HTTPCompressionConfig) MatchAcceptEncoding(acceptEncodings string) (compressionType HTTPCompressionType, compressionEncoding string, ok bool) {
	if len(acceptEncodings) == 0 {
		return
	}

	if len(this.types) == 0 {
		return
	}

	var pieces = strings.Split(acceptEncodings, ",")
	var encodings = []string{}
	for _, piece := range pieces {
		var qualityIndex = strings.Index(piece, ";")
		if qualityIndex >= 0 {
			// TODO 实现优先级
			piece = piece[:qualityIndex]
		}

		encodings = append(encodings, strings.TrimSpace(piece))
	}

	if len(encodings) == 0 {
		return
	}

	for _, supportType := range this.types {
		switch supportType {
		case HTTPCompressionTypeGzip:
			if this.supportGzip && lists.ContainsString(encodings, "gzip") {
				return HTTPCompressionTypeGzip, "gzip", true
			}
		case HTTPCompressionTypeDeflate:
			if this.supportDeflate && lists.ContainsString(encodings, "deflate") {
				return HTTPCompressionTypeDeflate, "deflate", true
			}
		case HTTPCompressionTypeBrotli:
			if this.supportBrotli && lists.ContainsString(encodings, "br") {
				return HTTPCompressionTypeBrotli, "br", true
			}
		case HTTPCompressionTypeZSTD:
			if this.supportZSTD && lists.ContainsString(encodings, "zstd") {
				return HTTPCompressionTypeZSTD, "zstd", true
			}
		}
	}

	return "", "", false
}

func (this *HTTPCompressionConfig) MatchURL(url string) bool {
	// except
	if len(this.ExceptURLPatterns) > 0 {
		for _, pattern := range this.ExceptURLPatterns {
			if pattern.Match(url) {
				return false
			}
		}
	}

	// only
	if len(this.OnlyURLPatterns) > 0 {
		for _, pattern := range this.OnlyURLPatterns {
			if pattern.Match(url) {
				return true
			}
		}
		return false
	}

	return true
}
