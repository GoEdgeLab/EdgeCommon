// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dnsconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/types"
)

type DNSResolver struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

func (this *DNSResolver) Addr() string {
	var port = this.Port
	if port <= 0 {
		// 暂时不支持DoH
		// 实际应用中只支持udp
		switch this.Protocol {
		case "tls":
			port = 853
		default:
			port = 53
		}
	}
	return configutils.QuoteIP(this.Host) + ":" + types.String(port)
}
