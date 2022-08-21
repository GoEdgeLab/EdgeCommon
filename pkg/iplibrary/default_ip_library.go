// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"github.com/iwind/TeaGo/logs"
	"net"
)

//go:embed internal-ip-library.db
var ipLibraryData []byte

var library = NewIPLibrary()

func init() {
	err := library.Init()
	if err != nil {
		logs.Println("IP_LIBRARY", "initialized failed: "+err.Error())
	}
}

func Lookup(ip net.IP) *QueryResult {
	return library.Lookup(ip)
}

func LookupIP(ip string) *QueryResult {
	return library.LookupIP(ip)
}

type IPLibrary struct {
	reader *Reader
}

func NewIPLibrary() *IPLibrary {
	return &IPLibrary{}
}

func (this *IPLibrary) Init() error {
	var reader = bytes.NewReader(ipLibraryData)
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
