package configutils

import (
	"encoding/binary"
	"github.com/cespare/xxhash/v2"
	"math"
	"net"
	"strings"
)

// IP2Long 将IP转换为整型
// 注意IPv6没有顺序
func IP2Long(ip string) uint64 {
	if len(ip) == 0 {
		return 0
	}
	s := net.ParseIP(ip)
	if len(s) == 0 {
		return 0
	}

	if strings.Contains(ip, ":") {
		return math.MaxUint32 + xxhash.Sum64(s)
	}
	return uint64(binary.BigEndian.Uint32(s.To4()))
}

// ParseCIDR 计算CIDR最大值
func ParseCIDR(cidr string) (ipFrom string, ipTo string, err error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", "", err
	}
	ipFrom = ipNet.IP.String()
Loop:
	for i := len(ipNet.Mask) - 1; i >= 0; i-- {
		for j := 7; j >= 0; j-- {
			var m = ipNet.Mask[i] >> (7 - j) & 1 // 读取某位bit
			if m == 0 {
				ipNet.IP[i] |= 1 << j // 将此位置为1
			} else {
				break Loop
			}
		}
	}
	ipTo = ipNet.IP.String()
	return
}

// QuoteIP 为IPv6加上括号
func QuoteIP(ip string) string {
	if len(ip) == 0 {
		return ip
	}
	if !strings.Contains(ip, ":") {
		return ip
	}
	if ip[0] != '[' {
		return "[" + ip + "]"
	}
	return ip
}
