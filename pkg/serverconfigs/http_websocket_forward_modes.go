package serverconfigs

import "github.com/iwind/TeaGo/maps"

// Websocket转发类型
type HTTPWebsocketForwardMode = string

const (
	HTTPWebsocketForwardModeWebsocket = "websocket"
	HTTPWebsocketForwardModeHttp      = "http"
)

// 所有的转发方式
func AllWebsocketForwardModes() []maps.Map {
	return []maps.Map{
		{
			"name":        "Websocket连接",
			"mode":        HTTPWebsocketForwardModeWebsocket,
			"description": "通过Websocket连接后端服务器并发送数据",
		},
		{
			"name":        "HTTP连接",
			"mode":        HTTPWebsocketForwardModeHttp,
			"description": "通过HTTP PUT转发服务器到后端服务器",
		},
	}
}
