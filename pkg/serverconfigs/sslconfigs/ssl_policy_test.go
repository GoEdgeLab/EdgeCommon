// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package sslconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestSSLPolicy_MatchDomain(t *testing.T) {
	var a = assert.NewAssertion(t)

	var policy = &SSLPolicy{}
	policy.Certs = []*SSLCertConfig{
		{
			Id:       1,
			DNSNames: []string{"a.com", "b.com"},
		},
		{
			Id:       2,
			DNSNames: []string{"c.com", "d.com"},
		},
		{
			Id:       3,
			DNSNames: []string{"e.com", "f.com"},
		},
	}

	{
		_, ok := policy.MatchDomain("c.com")
		a.IsTrue(ok)
	}
}
