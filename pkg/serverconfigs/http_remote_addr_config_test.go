// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHTTPRemoteAddrConfig_IsEmpty(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		var config = &HTTPRemoteAddrConfig{}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsEmpty())
	}

	{
		var config = &HTTPRemoteAddrConfig{
			Value: "${remoteAddr}",
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsEmpty())
	}

	{
		var config = &HTTPRemoteAddrConfig{
			Value: "${ remoteAddr }",
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsEmpty())
	}

	{
		var config = &HTTPRemoteAddrConfig{
			Value: "[${remoteAddr}]",
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(config.IsEmpty())
	}
}
