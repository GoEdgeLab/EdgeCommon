package configutils

import (
	"net"
	"strings"
)

// IsIPv4 检查是否为IPv4
func IsIPv4(netIP net.IP) bool {
	if len(netIP) == 0 {
		return false
	}
	return netIP.To4() != nil
}

// IsIPv6 检查是否为IPv6
func IsIPv6(netIP net.IP) bool {
	if len(netIP) == 0 {
		return false
	}
	return netIP.To4() == nil && netIP.To16() != nil
}

// IPVersion 获取IP版本号
func IPVersion(netIP net.IP) int {
	if len(netIP) == 0 {
		return 0
	}

	if netIP.To4() != nil {
		return 4
	}

	if netIP.To16() != nil {
		return 6
	}

	return 0
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
