// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"net"
	"sync"
)

//go:embed internal-ip-library.db
var ipLibraryData []byte

var defaultLibrary = NewIPLibrary()
var commonLibrary *IPLibrary

var libraryLocker = &sync.Mutex{} // 为了保持加载顺序性

func DefaultIPLibraryData() []byte {
	return ipLibraryData
}

// InitDefault 加载默认的IP库
func InitDefault() error {
	libraryLocker.Lock()
	defer libraryLocker.Unlock()

	if commonLibrary != nil {
		defaultLibrary = commonLibrary
		return nil
	}

	var library = NewIPLibrary()
	err := library.InitFromData(ipLibraryData, "")
	if err != nil {
		return err
	}

	commonLibrary = library
	defaultLibrary = commonLibrary
	return nil
}

// Lookup 查询IP信息
func Lookup(ip net.IP) *QueryResult {
	return defaultLibrary.Lookup(ip)
}

// LookupIP 查询IP信息
func LookupIP(ip string) *QueryResult {
	return defaultLibrary.LookupIP(ip)
}

// LookupIPSummaries 查询一组IP对应的区域描述
func LookupIPSummaries(ipList []string) map[string]string /** ip => summary **/ {
	var result = map[string]string{}
	for _, ip := range ipList {
		var region = LookupIP(ip)
		if region != nil && region.IsOk() {
			result[ip] = region.Summary()
		}
	}
	return result
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

func (this *IPLibrary) InitFromData(data []byte, password string) error {
	if len(data) == 0 || this.reader != nil {
		return nil
	}

	if len(password) > 0 {
		srcData, err := NewEncrypt().Decode(data, password)
		if err != nil {
			return err
		}
		data = srcData
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

func (this *IPLibrary) Destroy() {
	if this.reader != nil {
		this.reader.Destroy()
		this.reader = nil
	}
}
