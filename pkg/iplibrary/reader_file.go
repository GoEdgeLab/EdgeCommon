// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"compress/gzip"
	"net"
	"os"
)

type FileReader struct {
	rawReader *Reader
}

func NewFileReader(path string) (*FileReader, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fp.Close()
	}()

	gzReader, err := gzip.NewReader(fp)
	if err != nil {
		return nil, err
	}

	reader, err := NewReader(gzReader)
	if err != nil {
		return nil, err
	}

	return &FileReader{
		rawReader: reader,
	}, nil
}

func (this *FileReader) Lookup(ip net.IP) *QueryResult {
	return this.rawReader.Lookup(ip)
}
