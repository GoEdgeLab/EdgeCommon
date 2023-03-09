// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package shared

import (
	"errors"
	"regexp"
	"strings"
)

type URLPatternType = string

const (
	URLPatternTypeWildcard URLPatternType = "wildcard" // 通配符
	URLPatternTypeRegexp   URLPatternType = "regexp"   // 正则表达式
)

type URLPattern struct {
	Type    URLPatternType `yaml:"type" json:"type"`
	Pattern string         `yaml:"pattern" json:"pattern"`

	reg *regexp.Regexp
}

func (this *URLPattern) Init() error {
	if len(this.Pattern) == 0 {
		return nil
	}

	switch this.Type {
	case URLPatternTypeWildcard:
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
	case URLPatternTypeRegexp:
		var pattern = this.Pattern
		if !strings.HasPrefix(pattern, "(?i)") { // 大小写不敏感
			pattern = "(?i)" + pattern
		}
		reg, err := regexp.Compile(pattern)
		if err != nil {
			return errors.New("compile '" + pattern + "' failed: " + err.Error())
		}
		this.reg = reg
	}

	return nil
}

func (this *URLPattern) Match(url string) bool {
	if len(this.Pattern) == 0 && len(url) == 0 {
		return true
	}

	if this.reg != nil {
		return this.reg.MatchString(url)
	}
	return false
}
