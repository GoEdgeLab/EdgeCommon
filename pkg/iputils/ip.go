// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iputils

import (
	"encoding/binary"
	"encoding/hex"
	"math"
	"math/big"
	"net"
	"strconv"
	"sync"
)

type IP struct {
	rawIP  net.IP
	bigInt *big.Int
}

var uint32BigInt = big.NewInt(int64(math.MaxUint32))

func ParseIP(ipString string) IP {
	return NewIP(net.ParseIP(ipString))
}

func NewIP(rawIP net.IP) IP {
	if rawIP == nil {
		return IP{}
	}

	if rawIP.To4() == nil {
		var bigInt = big.NewInt(0)
		bigInt.SetBytes(rawIP.To16())
		bigInt.Add(bigInt, uint32BigInt)
		return IP{
			rawIP:  rawIP,
			bigInt: bigInt,
		}
	}

	return IP{
		rawIP: rawIP,
	}
}

func IsIPv4(ipString string) bool {
	var rawIP = net.ParseIP(ipString)
	return rawIP != nil && rawIP.To4() != nil
}

func IsIPv6(ipString string) bool {
	var rawIP = net.ParseIP(ipString)
	return rawIP != nil && rawIP.To4() == nil && rawIP.To16() != nil
}

func CompareLong(i1 string, i2 string) int {
	if i1 == "" {
		i1 = "0"
	}
	if i2 == "" {
		i2 = "0"
	}

	var l = len(i1) - len(i2)
	if l > 0 {
		return 1
	}
	if l < 0 {
		return -1
	}

	if i1 > i2 {
		return 1
	}
	if i1 < i2 {
		return -1
	}

	return 0
}

var bigIntPool = &sync.Pool{
	New: func() any {
		return big.NewInt(0)
	},
}

func ToLong(ip string) string {
	var rawIP = net.ParseIP(ip)
	if rawIP == nil {
		return "0"
	}

	var i4 = rawIP.To4()
	if i4 != nil {
		return strconv.FormatUint(uint64(binary.BigEndian.Uint32(i4)), 10)
	}

	var bigInt = bigIntPool.Get().(*big.Int)
	bigInt.SetBytes(rawIP.To16())
	bigInt.Add(bigInt, uint32BigInt)
	var s = bigInt.String()
	bigIntPool.Put(bigInt)
	return s
}

func ToHex(ip string) string {
	var rawIP = net.ParseIP(ip)
	if rawIP == nil {
		return ""
	}

	if rawIP.To4() != nil {
		return hex.EncodeToString(rawIP.To4())
	}

	return hex.EncodeToString(rawIP.To16())
}

func ToLittleLong(ip string) string {
	var rawIP = net.ParseIP(ip)
	if rawIP == nil {
		return "0"
	}

	var i4 = rawIP.To4()
	if i4 != nil {
		return strconv.FormatUint(uint64(binary.BigEndian.Uint32(i4)), 10)
	}

	var bigInt = bigIntPool.Get().(*big.Int)
	bigInt.SetBytes(rawIP.To16())
	var s = bigInt.String()
	bigIntPool.Put(bigInt)
	return s
}

func (this IP) ToLong() string {
	if this.rawIP == nil {
		return "0"
	}
	if this.bigInt != nil {
		return this.bigInt.String()
	}
	return strconv.FormatUint(uint64(binary.BigEndian.Uint32(this.rawIP.To4())), 10)
}

func (this IP) Mod(d int) int {
	if this.rawIP == nil {
		return 0
	}
	if this.bigInt != nil {
		return int(this.bigInt.Mod(this.bigInt, big.NewInt(int64(d))).Int64())
	}
	return int(binary.BigEndian.Uint32(this.rawIP.To4()) % uint32(d))
}

func (this IP) Compare(anotherIP IP) int {
	if this.rawIP == nil {
		if anotherIP.rawIP == nil {
			return 0
		}
		return -1
	} else if anotherIP.rawIP == nil {
		return 1
	}

	if this.bigInt != nil {
		if anotherIP.bigInt == nil {
			return 1 // IPv6 always greater than IPv4
		}
		return this.bigInt.Cmp(anotherIP.bigInt)
	}

	if anotherIP.bigInt == nil {
		var i1 = binary.BigEndian.Uint32(this.rawIP.To4())
		var i2 = binary.BigEndian.Uint32(anotherIP.rawIP.To4())

		if i1 > i2 {
			return 1
		}
		if i1 < i2 {
			return -1
		}
		return 0
	}

	return -1
}

func (this IP) Between(ipFrom IP, ipTo IP) bool {
	return ipFrom.Compare(this) <= 0 && ipTo.Compare(this) >= 0
}

func (this IP) IsIPv4() bool {
	return this.rawIP != nil && this.bigInt == nil
}

func (this IP) IsIPv6() bool {
	return this.bigInt != nil
}

func (this IP) IsValid() bool {
	return this.rawIP != nil
}

func (this IP) Raw() net.IP {
	return this.rawIP
}

func (this IP) String() string {
	if this.rawIP == nil {
		return ""
	}
	return this.rawIP.String()
}
