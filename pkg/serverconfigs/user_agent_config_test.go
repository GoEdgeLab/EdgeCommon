// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/assert"
	"net/http"
	"testing"
)

func TestUserAgentConfig_AllowRequest(t *testing.T) {
	var a = assert.NewAssertion(t)

	var config = serverconfigs.NewUserAgentConfig()
	config.Filters = append(config.Filters, &serverconfigs.UserAgentFilter{
		Keywords: []string{"Chrome", "Google*Bot", "Opera*a*c", "|(*"},
		Action:   serverconfigs.UserAgentActionAllow,
	})
	config.Filters = append(config.Filters, &serverconfigs.UserAgentFilter{
		Keywords: []string{"Google*a*c"},
		Action:   serverconfigs.UserAgentActionDeny,
	})
	config.Filters = append(config.Filters, &serverconfigs.UserAgentFilter{
		Keywords: []string{""},
		Action:   serverconfigs.UserAgentActionDeny,
	})
	config.Filters = append(config.Filters, &serverconfigs.UserAgentFilter{
		Keywords: []string{"mozilla", "firefox"},
		Action:   serverconfigs.UserAgentActionDeny,
	})
	err := config.Init()
	if err != nil {
		t.Fatal(err)
	}
	{
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		{
			req.Header.Set("User-Agent", "")
			a.IsFalse(config.AllowRequest(req))
		}
		{
			req.Header.Set("User-Agent", "chrome")
			a.IsTrue(config.AllowRequest(req))
		}
		{
			req.Header.Set("User-Agent", "mozilla")
			a.IsFalse(config.AllowRequest(req))
		}
		{
			req.Header.Set("User-Agent", "Firefox")
			a.IsFalse(config.AllowRequest(req))
		}
		{
			req.Header.Set("User-Agent", "Google Bot")
			a.IsTrue(config.AllowRequest(req))
		}
		{
			req.Header.Set("User-Agent", "opera abc")
			a.IsTrue(config.AllowRequest(req))
		}
		{
			req.Header.Set("User-Agent", "google abc")
			a.IsFalse(config.AllowRequest(req))
		}
	}
}
