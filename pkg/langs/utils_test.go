// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package langs_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/langs"
	"testing"
)

func TestMessage(t *testing.T) {
	var lang = langs.DefaultManager().AddLang("en-us")
	lang.Set("user_description", "user: %s")
	t.Log(langs.Message("en-us", "user_description", "Lily"))
}

func TestMessageDefault(t *testing.T) {
	var lang = langs.DefaultManager().AddLang("zh-cn")
	lang.Set("user_description", "user: %s")

	t.Log(langs.Message("en-us", "user_description", "Lily"))
}

func TestFormat(t *testing.T) {
	{
		var lang = langs.DefaultManager().AddLang("en-US")
		lang.Set("book_name", "Golang")
	}

	{
		var lang = langs.DefaultManager().AddLang("zh-CN")
		lang.Set("book_name", "Go语言")
	}

	t.Log(langs.Format("en-US", "this is ${lang.book_name} book"))
	t.Log(langs.Format("zh-CN", "this is ${lang.book_name} book"))
}
