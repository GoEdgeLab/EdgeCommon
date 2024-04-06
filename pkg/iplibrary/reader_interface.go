// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import "net"

type ReaderVersion = int

const (
	ReaderVersionV1 ReaderVersion = 0
	ReaderVersionV2 ReaderVersion = 2
)

type ReaderInterface interface {
	Meta() *Meta
	Lookup(ip net.IP) *QueryResult
	Destroy()
}
