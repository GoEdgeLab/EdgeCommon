// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package sslconfigs_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestSSLCertConfig_MatchDomain(t *testing.T) {
	var a = assert.NewAssertion(t)

	var cert = &sslconfigs.SSLCertConfig{
		DNSNames: []string{"a.com", "b.com"},
	}
	a.IsTrue(cert.MatchDomain("a.com"))
	a.IsFalse(cert.MatchDomain("z.com"))
}

/**func TestSSLCertConfig_DNSNames(t *testing.T) {
	var config = sslconfigs.SSLCertConfig{}
	config.CertData = []byte(`YOUR CERT DATA`)
	config.KeyData = []byte(`YOUR KEY DATA`)

	err := config.Init(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(config.DNSNames)
}**/
