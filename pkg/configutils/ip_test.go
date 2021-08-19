// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package configutils

import "testing"

func TestParseCIDR(t *testing.T) {
	t.Log(ParseCIDR("192.168.1.1/32"))
	t.Log(ParseCIDR("192.168.1.1/24"))
	t.Log(ParseCIDR("192.168.1.1/16"))
}
