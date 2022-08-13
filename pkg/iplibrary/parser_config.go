// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

type ParserConfig struct {
	Template    *Template
	EmptyValues []string
	Iterator    func(values map[string]string) error
}
