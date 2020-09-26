package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"time"
)

// websocket设置
type HTTPWebsocketConfig struct {
	Id   int64 `yaml:"id" json:"id"`     // ID
	IsOn bool  `yaml:"isOn" json:"isOn"` // 是否开启

	// 握手超时时间
	HandshakeTimeout *shared.TimeDuration `yaml:"handshakeTimeout" json:"handshakeTimeout"`

	// 允许的来源域名，支持 www.example.com, example.com, .example.com, *.example.com
	AllowAllOrigins bool     `yaml:"allowAllOrigins" json:"allowAllOrigins"`
	AllowedOrigins  []string `yaml:"allowedOrigins" json:"allowedOrigins"`

	// 向后传递的来源
	RequestSameOrigin bool   `yaml:"requestSameOrigin" json:"requestSameOrigin"` // 和请求一致
	RequestOrigin     string `yaml:"requestOrigin" json:"requestOrigin"`         // 自行指定Origin，支持变量

	handshakeTimeoutDuration  time.Duration
	requestOriginHasVariables bool
}

// 校验
func (this *HTTPWebsocketConfig) Init() error {
	// duration
	if this.HandshakeTimeout != nil {
		this.handshakeTimeoutDuration = this.HandshakeTimeout.Duration()
	}

	// requestOrigin
	this.requestOriginHasVariables = configutils.HasVariables(this.RequestOrigin)

	return nil
}

// 获取握手超时时间
func (this *HTTPWebsocketConfig) HandshakeTimeoutDuration() time.Duration {
	return this.handshakeTimeoutDuration
}

// 匹配域名
func (this *HTTPWebsocketConfig) MatchOrigin(origin string) bool {
	if this.AllowAllOrigins {
		return true
	}
	return configutils.MatchDomains(this.AllowedOrigins, origin)
}

// 判断请求Origin是否有变量
func (this *HTTPWebsocketConfig) RequestOriginHasVariables() bool {
	return this.requestOriginHasVariables
}
