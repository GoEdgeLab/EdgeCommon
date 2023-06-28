// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package langs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/actions"
	"net/http"
	"strings"
)

// Message 读取消息
// Read message
func Message(langCode LangCode, messageCode MessageCode, args ...any) string {
	return defaultManager.GetMessage(langCode, messageCode, args...)
}

func ParseLangFromRequest(req *http.Request) (langCode string) {
	// parse language from cookie
	const cookieName = "edgelang"
	cookie, _ := req.Cookie(cookieName)
	if cookie != nil && len(cookie.Value) > 0 && defaultManager.HasLang(cookie.Value) {
		return cookie.Value
	}

	// parse language from 'Accept-Language'
	var acceptLanguage = req.Header.Get("Accept-Language")
	if len(acceptLanguage) > 0 {
		var pieces = strings.Split(acceptLanguage, ",")
		for _, lang := range pieces {
			var index = strings.Index(lang, ";")
			if index >= 0 {
				lang = lang[:index]
			}

			var match = defaultManager.MatchLang(lang)
			if len(match) > 0 {
				return match
			}
		}
	}

	return defaultManager.DefaultLang()
}

func ParseLangFromAction(action actions.ActionWrapper) (langCode string) {
	return ParseLangFromRequest(action.Object().Request)
}

// Format 格式化变量
// Format string that contains message variables, such as ${lang.MESSAGE_CODE}
//
// 暂时不支持变量中加参数
func Format(langCode LangCode, varString string) string {
	return configutils.ParseVariables(varString, func(varName string) (value string) {
		if !strings.HasPrefix(varName, varPrefix) {
			return "${" + varName + "}" // keep origin variable
		}
		return Message(langCode, varName[len(varPrefix):])
	})
}

// Load 加载消息定义
// Load message definitions from map
func Load(langCode LangCode, messageMap map[string]string) {
	lang, ok := defaultManager.GetLang(langCode)
	if !ok {
		lang = defaultManager.AddLang(langCode)
	}
	for messageCode, messageText := range messageMap {
		lang.Set(messageCode, messageText)
	}
}
