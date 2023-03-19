// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package shared

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"sync"
)

var dataMapPrefix = []byte("GOEDGE_DATA_MAP:")

// DataMap 二进制数据共享Map
// 用来减少相同数据占用的空间和内存
type DataMap struct {
	Map map[string][]byte
	locker sync.Mutex
}

// NewDataMap 构建对象
func NewDataMap() *DataMap {
	return &DataMap{Map: map[string][]byte{}}
}

// Put 放入数据
func (this *DataMap) Put(data []byte) (keyData []byte) {
	this.locker.Lock()
	defer this.locker.Unlock()
	var key = string(dataMapPrefix) + fmt.Sprintf("%x", md5.Sum(data))
	this.Map[key] = data
	return []byte(key)
}

// Read 读取数据
func (this *DataMap) Read(key []byte) []byte {
	this.locker.Lock()
	defer this.locker.Unlock()
	if bytes.HasPrefix(key, dataMapPrefix) {
		return this.Map[string(key)]
	}
	return key
}
