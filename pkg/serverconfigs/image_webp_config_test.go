// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestWebPImageConfig_MatchAccept(t *testing.T) {
	var a = assert.NewAssertion(t)
	{
		var c = &WebPImageConfig{}
		a.IsFalse(c.MatchAccept(""))
		a.IsTrue(c.MatchAccept("image/webp"))
		a.IsTrue(c.MatchAccept("image/webp,image/png"))
		a.IsTrue(c.MatchAccept("image/jpeg,image/webp,image/png"))
		a.IsFalse(c.MatchAccept("mimage/webp"))
		a.IsFalse(c.MatchAccept("image/webpm"))
	}
}
