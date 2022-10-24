// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestReferersConfig_MatchDomain(t *testing.T) {
	var a = assert.NewAssertion(t)

	{
		var config = &serverconfigs.ReferersConfig{
			IsOn:         true,
			AllowDomains: []string{},
			DenyDomains:  []string{"a.com", "b.com"},
		}
		a.IsTrue(config.MatchDomain("example.com", "c.com"))
		a.IsTrue(config.MatchDomain("example.com", "d.com"))
		a.IsFalse(config.MatchDomain("example.com", "a.com"))
	}

	{
		var config = &serverconfigs.ReferersConfig{
			IsOn:         true,
			AllowDomains: []string{"c.com", "e.com"},
			DenyDomains:  []string{"a.com", "b.com", "e.com"},
		}
		a.IsTrue(config.MatchDomain("example.com", "c.com"))
		a.IsFalse(config.MatchDomain("example.com", "d.com"))
		a.IsFalse(config.MatchDomain("example.com", "e.com"))
		a.IsFalse(config.MatchDomain("example.com", "a.com"))
	}

	{
		var config = &serverconfigs.ReferersConfig{
			IsOn:         true,
			AllowDomains: []string{"c.com", "e.com"},
			DenyDomains:  []string{},
		}
		a.IsTrue(config.MatchDomain("example.com", "c.com"))
		a.IsFalse(config.MatchDomain("example.com", "d.com"))
		a.IsTrue(config.MatchDomain("example.com", "e.com"))
		a.IsFalse(config.MatchDomain("example.com", "a.com"))
	}
}
