// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package firewallconfigs_test

import (
	"github.com/iwind/TeaGo/assert"
	"regexp"
	"testing"
)

func TestUserAgentMatch_PHP(t *testing.T) {
	var a = assert.NewAssertion(t)

	var expr = `python|pycurl|http-client|httpclient|apachebench|nethttp|http_request|java|perl|ruby|scrapy|php\b|rust`
	var reg = regexp.MustCompile(expr)
	a.IsTrue(reg.MatchString("php"))
	a.IsTrue(reg.MatchString("php/5.0"))
	a.IsFalse(reg.MatchString("php110"))
}
