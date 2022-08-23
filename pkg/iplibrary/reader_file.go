// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"compress/gzip"
	"errors"
	"io"
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

	return NewFileDataReader(fp)
}

func NewFileDataReader(dataReader io.Reader) (*FileReader, error) {
	gzReader, err := gzip.NewReader(dataReader)
	if err != nil {
		return nil, errors.New("create gzip reader failed: " + err.Error())
	}

	reader, err := NewReader(gzReader)
	if err != nil {
		return nil, err
	}

	return &FileReader{
		rawReader: reader,
	}, nil
}

func (this *FileReader) Meta() *Meta {
	return this.rawReader.meta
}

func (this *FileReader) Lookup(ip net.IP) *QueryResult {
	return this.rawReader.Lookup(ip)
}

func (this *FileReader) RawReader() *Reader {
	return this.rawReader
}
