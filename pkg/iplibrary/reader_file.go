// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type FileReader struct {
	rawReader ReaderInterface
	//password  string
}

func NewFileReader(path string, password string) (*FileReader, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fp.Close()
	}()

	var version = ReaderVersionV1
	if strings.HasSuffix(filepath.Base(path), ".v2.db") {
		version = ReaderVersionV2
	}

	return NewFileDataReader(fp, password, version)
}

func NewFileDataReader(dataReader io.Reader, password string, readerVersion ReaderVersion) (*FileReader, error) {
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
		return nil, fmt.Errorf("create gzip reader failed: %w", err)
	}

	var reader ReaderInterface
	if readerVersion == ReaderVersionV2 {
		reader, err = NewReaderV2(gzReader)
	} else {
		reader, err = NewReaderV1(gzReader)
	}
	if err != nil {
		return nil, err
	}

	return &FileReader{
		rawReader: reader,
	}, nil
}

func (this *FileReader) Meta() *Meta {
	return this.rawReader.Meta()
}

func (this *FileReader) Lookup(ip net.IP) *QueryResult {
	return this.rawReader.Lookup(ip)
}

func (this *FileReader) RawReader() ReaderInterface {
	return this.rawReader
}
