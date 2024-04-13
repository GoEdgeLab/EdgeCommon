// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iputils

import (
	"net"
)

type CIDR struct {
	rawIPNet *net.IPNet
}

func ParseCIDR(s string) (*CIDR, error) {
	_, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}
	return &CIDR{
		rawIPNet: ipNet,
	}, nil
}

func (this *CIDR) IsIPv4() bool {
	return this.rawIPNet.IP.To4() != nil
}

func (this *CIDR) IsIPv6() bool {
	return this.rawIPNet.IP.To4() == nil
}

func (this *CIDR) From() net.IP {
	return this.rawIPNet.IP
}

func (this *CIDR) To() net.IP {
	var start = this.rawIPNet.IP.To4()
	if start != nil {
		return bitsOr(bitsAnd(start, this.rawIPNet.Mask), bitsXor(this.rawIPNet.Mask[:4], []byte{0xff, 0xff, 0xff, 0xff}))
	}

	start = this.rawIPNet.IP.To16()
	return bitsOr(bitsAnd(start, this.rawIPNet.Mask), bitsXor(this.rawIPNet.Mask[:16], []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}))
}

func (this *CIDR) Contains(ip net.IP) bool {
	return this.rawIPNet.Contains(ip)
}

func (this *CIDR) String() string {
	return this.rawIPNet.String()
}

func bitsAnd(x []byte, y []byte) []byte {
	var l = len(x)
	var z = make([]byte, l)
	for i := 0; i < l; i++ {
		z[i] = x[i] & y[i]
	}
	return z
}

func bitsOr(x []byte, y []byte) []byte {
	var l = len(x)
	var z = make([]byte, l)
	for i := 0; i < l; i++ {
		z[i] = x[i] | y[i]
	}
	return z
}

func bitsXor(x []byte, y []byte) []byte {
	var l = len(x)
	var z = make([]byte, l)
	for i := 0; i < l; i++ {
		z[i] = x[i] ^ y[i]
	}
	return z
}
