// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"net"
	"os"
)

type FileReader struct {
	rawReader *Reader
	password  string
}

func NewFileReader(path string, password string) (*FileReader, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fp.Close()
	}()

	return NewFileDataReader(fp, password)
}

func NewFileDataReader(dataReader io.Reader, password string) (*FileReader, error) {
	if len(password) > 0 {
		data, err := io.ReadAll(dataReader)
		if err != nil {
			return nil, err
		}

		sourceData, err := NewEncrypt().Decode(data, password)
		if err != nil {
			return nil, err
		}

		dataReader = bytes.NewReader(sourceData)
	}

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
