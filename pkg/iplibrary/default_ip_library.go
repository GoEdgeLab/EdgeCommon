// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"net"
)

//go:embed internal-ip-library.db
var ipLibraryData []byte

var defaultLibrary = NewIPLibrary()

func DefaultIPLibraryData() []byte {
	return ipLibraryData
}

func InitDefault() error {
	defaultLibrary.reader = nil
	return defaultLibrary.InitFromData(ipLibraryData)
}

func Lookup(ip net.IP) *QueryResult {
	return defaultLibrary.Lookup(ip)
}

func LookupIP(ip string) *QueryResult {
	return defaultLibrary.LookupIP(ip)
}

type IPLibrary struct {
	reader *Reader
}

func NewIPLibrary() *IPLibrary {
	return &IPLibrary{}
}

func NewIPLibraryWithReader(reader *Reader) *IPLibrary {
	return &IPLibrary{reader: reader}
}

func (this *IPLibrary) InitFromData(data []byte) error {
	if len(data) == 0 || this.reader != nil {
		return nil
	}
	var reader = bytes.NewReader(data)
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer func() {
		_ = gzipReader.Close()
	}()

	libReader, err := NewReader(gzipReader)
	if err != nil {
		return err
	}
	this.reader = libReader

	return nil
}

func (this *IPLibrary) Lookup(ip net.IP) *QueryResult {
	if this.reader == nil {
		return &QueryResult{}
	}

	var result = this.reader.Lookup(ip)
	if result == nil {
		result = &QueryResult{}
	}

	return result
}

func (this *IPLibrary) LookupIP(ip string) *QueryResult {
	if this.reader == nil {
		return &QueryResult{}
	}
	return this.Lookup(net.ParseIP(ip))
}
