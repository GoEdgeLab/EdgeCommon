// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package shared

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type URLPatternType = string

const (
	URLPatternTypeWildcard URLPatternType = "wildcard" // 通配符
	URLPatternTypeRegexp   URLPatternType = "regexp"   // 正则表达式
	URLPatternTypeImages   URLPatternType = "images"   // 常见图片
	URLPatternTypeAudios   URLPatternType = "audios"   // 常见音频
	URLPatternTypeVideos   URLPatternType = "videos"   // 常见视频
)

var commonImageExtensions = []string{".apng", ".avif", ".gif", ".jpg", ".jpeg", ".jfif", ".pjpeg", ".pjp", ".png", ".svg", ".webp", ".bmp", ".ico", ".cur", ".tif", ".tiff"}
var commonAudioExtensions = []string{".mp3", ".flac", ".wav", ".aac", ".ogg", ".m4a", ".wma", ".m3u8"} // m3u8 is special
var commonVideoExtensions = []string{".mp4", ".avi", ".mkv", ".mov", ".wmv", ".mpeg", ".3gp", ".webm", ".ts", ".m3u8"}

type URLPattern struct {
	Type    URLPatternType `yaml:"type" json:"type"`
	Pattern string         `yaml:"pattern" json:"pattern"`

	reg *regexp.Regexp
}

func (this *URLPattern) Init() error {
	switch this.Type {
	case URLPatternTypeWildcard:
		if len(this.Pattern) > 0 {
			// 只支持星号
			var pieces = strings.Split(this.Pattern, "*")
			for index, piece := range pieces {
				pieces[index] = regexp.QuoteMeta(piece)
			}
			var pattern = strings.Join(pieces, "(.*)")
			if len(pattern) > 0 && pattern[0] == '/' {
				pattern = "(http|https)://[\\w.-]+" + pattern
			}
			reg, err := regexp.Compile("(?i)" /** 大小写不敏感 **/ + "^" + pattern + "$")
			if err != nil {
				return err
			}
			this.reg = reg
		}
	case URLPatternTypeRegexp:
		if len(this.Pattern) > 0 {
			var pattern = this.Pattern
			if !strings.HasPrefix(pattern, "(?i)") { // 大小写不敏感
				pattern = "(?i)" + pattern
			}
			reg, err := regexp.Compile(pattern)
			if err != nil {
				return fmt.Errorf("compile '%s' failed: %w", pattern, err)
			}
			this.reg = reg
		}
	}

	return nil
}

func (this *URLPattern) Match(url string) bool {
	if len(this.Pattern) == 0 && len(url) == 0 {
		return true
	}

	switch this.Type {
	case URLPatternTypeImages:
		var urlExt = strings.ToLower(filepath.Ext(url))
		if len(urlExt) > 0 {
			for _, ext := range commonImageExtensions {
				if ext == urlExt {
					return true
				}
			}
		}
	case URLPatternTypeAudios:
		var urlExt = strings.ToLower(filepath.Ext(url))
		if len(urlExt) > 0 {
			for _, ext := range commonAudioExtensions {
				if ext == urlExt {
					return true
				}
			}
		}
	case URLPatternTypeVideos:
		var urlExt = strings.ToLower(filepath.Ext(url))
		if len(urlExt) > 0 {
			for _, ext := range commonVideoExtensions {
				if ext == urlExt {
					return true
				}
			}
		}
	default:
		if this.reg != nil {
			return this.reg.MatchString(url)
		}
	}

	return false
}
