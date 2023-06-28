// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package langs_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/langs"
	"testing"
)

func TestManager_GetMessage(t *testing.T) {
	var manager = langs.NewManager()
	var lang = manager.AddLang("en-US")
	lang.Set("user_description", "user: %s, age: %d")
	t.Log(manager.GetMessage("en-US", "user_description", "Lily", 23))
	t.Log(manager.GetMessage("zh-CN", "user_description", "Lucy", 23)) // use 'en-US' as fallback language
}

func TestManager_GetMessage2(t *testing.T) {
	var manager = langs.NewManager()
	manager.SetDefaultLang("zh-CN")
	var lang = manager.AddLang("en-US")
	lang.Set("user_description", "user: %s, age: %d")
	t.Log(manager.GetMessage("en-US", "user_description", "Lily", 23))
	t.Log(manager.GetMessage("zh-CN", "user_description", "Lucy", 23)) // should be empty
}

func TestManager_MatchLang(t *testing.T) {
	var manager = langs.NewManager()
	manager.AddLang("en-us")
	manager.AddLang("en")
	manager.AddLang("zh-cn")
	manager.AddLang("zh-hk")
	//manager.AddLang("zh-tw")

	for _, code := range []string{
		"en",
		"en-us",
		"zh-cn",
		"zh-tw",
	} {
		t.Log(code, "=>", manager.MatchLang(code))
	}
}
