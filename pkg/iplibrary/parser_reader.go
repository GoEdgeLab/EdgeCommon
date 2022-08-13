// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"errors"
	"io"
)

type ReaderParser struct {
	reader    io.Reader
	rawParser *Parser
}

func NewReaderParser(reader io.Reader, config *ParserConfig) (*ReaderParser, error) {
	if config == nil {
		config = &ParserConfig{}
	}

	if config.Template == nil {
		return nil, errors.New("template must not be nil")
	}

	parser, err := NewParser(config)
	if err != nil {
		return nil, err
	}

	return &ReaderParser{
		reader:    reader,
		rawParser: parser,
	}, nil
}

func (this *ReaderParser) Parse() error {
	var buf = make([]byte, 1024)
	for {
		n, err := this.reader.Read(buf)
		if n > 0 {
			this.rawParser.Write(buf[:n])
			parseErr := this.rawParser.Parse()
			if parseErr != nil {
				return parseErr
			}
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}
	return nil
}
