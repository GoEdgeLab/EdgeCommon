package serverconfigs

import (
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
)

type ServerType = string

const (
	ServerTypeHTTPProxy ServerType = "httpProxy"
	ServerTypeHTTPWeb   ServerType = "httpWeb"
	ServerTypeTCPProxy  ServerType = "tcpProxy"
	ServerTypeUnixProxy ServerType = "unixProxy"
	ServerTypeUDPProxy  ServerType = "udpProxy"
)

// AllServerTypes 获取所有的服务类型
func AllServerTypes() []maps.Map {
	return []maps.Map{
		{
			"name":        "CDN加速",
			"code":        ServerTypeHTTPProxy,
			"description": "可以通过CDN边缘节点分发源站内容。",
		},
		{
			"name":        "TCP负载均衡",
			"code":        ServerTypeTCPProxy,
			"description": "通过反向代理访问源站TCP服务",
		},
		/**{
			"name": "UNIX协议反向代理",
			"code": ServerTypeUnixProxy,
		},**/
		{
			"name":        "UDP负载均衡",
			"code":        ServerTypeUDPProxy,
			"description": "通过反向代理访问源站UDP服务",
		},
		{
			"name":        "HTTP Web服务",
			"code":        ServerTypeHTTPWeb,
			"description": "普通的HTTP Web服务，可以用来访问边缘节点上的静态文件内容。",
		},
	}
}

// FindServerType 查找服务类型
func FindServerType(code string) maps.Map {
	for _, m := range AllServerTypes() {
		if m.GetString("code") == code {
			return m
		}
	}
	return nil
}

// FindAllServerProtocols 查找所有协议
func FindAllServerProtocols() []maps.Map {
	return []maps.Map{
		{
			"name":        "HTTP",
			"code":        "http",
			"serverTypes": []ServerType{ServerTypeHTTPProxy, ServerTypeHTTPWeb},
		},
		{
			"name":        "HTTPS",
			"code":        "https",
			"serverTypes": []ServerType{ServerTypeHTTPProxy, ServerTypeHTTPWeb},
		},
		{
			"name":        "TCP",
			"code":        "tcp",
			"serverTypes": []ServerType{ServerTypeTCPProxy},
		},
		{
			"name":        "TLS",
			"code":        "tls",
			"serverTypes": []ServerType{ServerTypeTCPProxy},
		},
		{
			"name":        "Unix",
			"code":        "unix",
			"serverTypes": []ServerType{ServerTypeUnixProxy},
		},
		{
			"name":        "UDP",
			"code":        "udp",
			"serverTypes": []ServerType{ServerTypeUDPProxy},
		},
	}
}

// FindAllServerProtocolsForType 获取所有协议
func FindAllServerProtocolsForType(serverType ServerType) []maps.Map {
	var result = []maps.Map{}
	for _, p := range FindAllServerProtocols() {
		var serverTypes = p.GetSlice("serverTypes")
		if lists.Contains(serverTypes, serverType) {
			result = append(result, p)
		}
	}
	return result
}

// IsHTTPServerType 判断某个服务类型是否属于HTTP簇
func IsHTTPServerType(serverType ServerType) bool {
	return serverType == ServerTypeHTTPProxy || serverType == ServerTypeHTTPWeb
}
