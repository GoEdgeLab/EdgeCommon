// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package langs

import (
	"fmt"
	"strings"
)

var defaultManager = NewManager()

type Manager struct {
	langMap         map[string]*Lang // lang code => *Lang, lang code must be in lowercase
	defaultLangCode string
}

func NewManager() *Manager {
	return &Manager{
		langMap:         map[string]*Lang{},
		defaultLangCode: "zh-cn",
	}
}

func DefaultManager() *Manager {
	return defaultManager
}

func (this *Manager) AddLang(code string) *Lang {
	var lang = NewLang(code)
	this.langMap[code] = lang
	return lang
}

func (this *Manager) HasLang(code string) bool {
	_, ok := this.langMap[code]
	return ok
}

func (this *Manager) GetLang(code string) (lang *Lang, ok bool) {
	lang, ok = this.langMap[code]
	return
}

func (this *Manager) MatchLang(code string) (matchedCode string) {
	// lookup exact match
	code = strings.ToLower(code)
	_, ok := this.langMap[code]
	if ok {
		return code
	}

	// lookup language family, such as en-us, en
	if strings.Contains(code, "-") {
		code, _, _ = strings.Cut(code, "-")
	}
	for rawCode := range this.langMap {
		if strings.HasPrefix(rawCode, code+"-") { // en-us vs en
			return rawCode
		}
	}

	return this.DefaultLang()
}

func (this *Manager) SetDefaultLang(code string) {
	this.defaultLangCode = code
}

func (this *Manager) DefaultLang() string {
	if len(this.defaultLangCode) > 0 {
		return this.defaultLangCode
	}
	return "zh-cn"
}

// GetMessage
// message: name: %s, age: %d, salary: %.2f
func (this *Manager) GetMessage(langCode string, messageCode string, args ...any) string {
	var lang = this.langMap[langCode]
	if lang == nil && len(this.defaultLangCode) > 0 {
		lang = this.langMap[this.defaultLangCode]
	}
	if lang == nil {
		return ""
	}

	var message = lang.Get(messageCode)
	if len(message) == 0 {
		// try to get message from default lang
		if lang.code != this.defaultLangCode {
			var defaultLang = this.langMap[this.defaultLangCode]
			if defaultLang != nil {
				return defaultLang.Get(messageCode)
			}
		}

		return ""
	}

	if len(args) == 0 {
		return message
	}

	return fmt.Sprintf(message, args...)
}
