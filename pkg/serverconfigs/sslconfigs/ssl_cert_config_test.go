// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package sslconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestSSLCertConfig_MatchDomain(t *testing.T) {
	var a = assert.NewAssertion(t)

	var cert = &SSLCertConfig{
		DNSNames: []string{"a.com", "b.com"},
	}
	a.IsTrue(cert.MatchDomain("a.com"))
	a.IsFalse(cert.MatchDomain("z.com"))
}
