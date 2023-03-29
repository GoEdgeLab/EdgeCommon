// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"github.com/iwind/TeaGo/lists"
	"regexp"
)

type Template struct {
	templateString string
	reg            *regexp.Regexp
}

func NewTemplate(templateString string) (*Template, error) {
	var t = &Template{
		templateString: templateString,
	}
	err := t.init()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (this *Template) init() error {
	var template = regexp.QuoteMeta(this.templateString)
	var keywordReg = regexp.MustCompile(`\\\$\\{(\w+)\\}`)
	template = keywordReg.ReplaceAllStringFunc(template, func(keyword string) string {
		var matches = keywordReg.FindStringSubmatch(keyword)
		if len(matches) > 1 {
			switch matches[1] {
			case "ipFrom", "ipTo", "country", "province", "city", "town", "provider":
				return "(?P<" + matches[1] + ">.*)"
			}
			return ".*"
		}

		return keyword
	})
	reg, err := regexp.Compile("^(?U)" + template)
	if err != nil {
		return err
	}
	this.reg = reg
	return nil
}

func (this *Template) Extract(text string, emptyValues []string) (values map[string]string, ok bool) {
	var matches = this.reg.FindStringSubmatch(text)
	if len(matches) == 0 {
		return
	}
	values = map[string]string{}
	for index, name := range this.reg.SubexpNames() {
		if len(name) == 0 {
			continue
		}
		var v = matches[index]
		if name != "ipFrom" && name != "ipTo" && (v == "0" || v == "无" || v == "空" || lists.ContainsString(emptyValues, v)) {
			v = ""
		}
		values[name] = v
	}

	for _, keyword := range []string{"ipFrom", "ipTo", "country", "province", "city", "town", "provider"} {
		_, hasKeyword := values[keyword]
		if !hasKeyword {
			values[keyword] = ""
		}
	}

	ok = true
	return
}
