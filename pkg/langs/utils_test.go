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

func TestMessageVariable(t *testing.T) {
	var lang = langs.DefaultManager().AddLang("en-us")
	lang.Set("user", "Lily")
	lang.Set("user2", "${lang.user}")
	lang.Set("user_name", "Name: ${lang.user}")
	lang.Set("user_name3", "Name: ${lang.user3}, ${user4}")
	//lang.Set("user3", "")

	err := lang.Compile()
	if err != nil {
		t.Log("ERROR(ignore):", err)
	}

	t.Log("user:", lang.Get("user"))
	t.Log("user2:", lang.Get("user2"))
	t.Log("user_name:", lang.Get("user_name"))
	t.Log("user_name3:", lang.Get("user_name3"))
	t.Log("user_name3_2:", lang.Get("name3"))
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

	t.Log(langs.Format("en-US", "this is ${lang:book_name} book"))
	t.Log(langs.Format("zh-CN", "this is ${lang:book_name} book"))
}
