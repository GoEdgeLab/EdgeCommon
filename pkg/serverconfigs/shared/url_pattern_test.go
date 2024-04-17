// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package shared_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"testing"
)

func TestURLPattern_Match(t *testing.T) {
	type unitTest struct {
		patternType string
		pattern     string
		url         string
		result      bool
	}

	for _, ut := range []*unitTest{
		{
			patternType: "wildcard",
			pattern:     "*",
			url:         "https://example.com",
			result:      true,
		},
		{
			patternType: "wildcard",
			pattern:     "https://example*",
			url:         "https://example.com",
			result:      true,
		},
		{
			patternType: "wildcard",
			pattern:     "*com",
			url:         "https://example.com",
			result:      true,
		},
		{
			patternType: "wildcard",
			pattern:     "*COM",
			url:         "https://example.com",
			result:      true,
		},
		{
			patternType: "wildcard",
			pattern:     "*COM",
			url:         "https://example.com/hello",
			result:      false,
		},
		{
			patternType: "wildcard",
			pattern:     "http://*",
			url:         "https://example.com",
			result:      false,
		},
		{
			patternType: "wildcard",
			pattern:     "https://example.com",
			url:         "https://example.com",
			result:      true,
		},
		{
			patternType: "wildcard",
			pattern:     "/hello/world",
			url:         "https://example-test.com/hello/world",
			result:      true,
		},
		{
			patternType: "wildcard",
			pattern:     "/hello/world",
			url:         "https://example-test.com/123/hello/world",
			result:      false,
		},
		{
			patternType: "wildcard",
			pattern:     "/hidden/*",
			url:         "/hidden/index.html",
			result:      false, // because don't have https://HOST in url
		},
		{
			patternType: "wildcard",
			pattern:     "*.jpg",
			url:         "https://example.com/index.jpg",
			result:      true,
		},
		{
			patternType: "wildcard",
			pattern:     "*.jpg",
			url:         "https://example.com/index.js",
			result:      false,
		},
		{
			patternType: "regexp",
			pattern:     ".*",
			url:         "https://example.com",
			result:      true,
		},
		{
			patternType: "regexp",
			pattern:     "^https://.*",
			url:         "https://example.com",
			result:      true,
		},
		{
			patternType: "regexp",
			pattern:     "^https://.*EXAMPLE.COM",
			url:         "https://example.com",
			result:      true,
		},
		{
			patternType: "regexp",
			pattern:     "(?i)https://.*EXAMPLE.COM/\\d+",
			url:         "https://example.com/123456",
			result:      true,
		},
		{
			patternType: "regexp",
			pattern:     "(?i)https://.*EXAMPLE.COM/\\d+$",
			url:         "https://example.com/123456/789",
			result:      false,
		},
		{
			patternType: "images",
			url:         "https://example.com/images/logo.png",
			result:      true,
		},
		{
			patternType: "images",
			url:         "https://example.com/images/logo.webp",
			result:      true,
		},
		{
			patternType: "images",
			url:         "https://example.com/images/logo.mp3",
			result:      false,
		},
		{
			patternType: "audios",
			url:         "https://example.com/audios/music.mp3",
			result:      true,
		},
		{
			patternType: "audios",
			url:         "https://example.com/audios/music.mm",
			result:      false,
		},
		{
			patternType: "videos",
			url:         "https://example.com/images/movie.mp4",
			result:      true,
		},
		{
			patternType: "videos",
			url:         "https://example.com/images/movie.ts",
			result:      true,
		},
		{
			patternType: "videos",
			url:         "https://example.com/images/movie.mp5",
			result:      false,
		},
	} {
		var p = &shared.URLPattern{
			Type:    ut.patternType,
			Pattern: ut.pattern,
		}
		err := p.Init()
		if err != nil {
			t.Fatal(err)
		}
		var b = p.Match(ut.url) == ut.result
		if !b {
			t.Fatal("not matched pattern:", ut.pattern, "url:", ut.url)
		}
	}
}
