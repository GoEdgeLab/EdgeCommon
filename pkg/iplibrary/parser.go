// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"errors"
)

type Parser struct {
	config *ParserConfig

	data []byte
}

func NewParser(config *ParserConfig) (*Parser, error) {
	if config == nil {
		config = &ParserConfig{}
	}

	if config.Template == nil {
		return nil, errors.New("template must not be nil")
	}

	return &Parser{
		config: config,
	}, nil
}

func (this *Parser) Write(data []byte) {
	this.data = append(this.data, data...)
}

func (this *Parser) Parse() error {
	if len(this.data) == 0 {
		return nil
	}
	for {
		var index = bytes.IndexByte(this.data, '\n')
		if index >= 0 {
			var line = this.data[:index+1]
			values, found := this.config.Template.Extract(string(line), this.config.EmptyValues)
			if found {
				if this.config.Iterator != nil {
					err := this.config.Iterator(values)
					if err != nil {
						return err
					}
				}
			} else {
				// 防止错误信息太长
				if len(line) > 256 {
					line = line[:256]
				}
				return errors.New("invalid line '" + string(line) + "'")
			}

			this.data = this.data[index+1:]
		} else {
			break
		}
	}
	return nil
}
