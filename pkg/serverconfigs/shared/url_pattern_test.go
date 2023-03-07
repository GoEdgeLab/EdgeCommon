// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package shared_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestURLPattern_Match(t *testing.T) {
	var a = assert.NewAssertion(t)

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
	} {
		var p = &shared.URLPattern{
			Type:    ut.patternType,
			Pattern: ut.pattern,
		}
		err := p.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(p.Match(ut.url) == ut.result)
	}
}
