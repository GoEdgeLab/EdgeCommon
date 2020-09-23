package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
	"time"
)

// websocket设置
type HTTPWebsocketConfig struct {
	Id   int64 `yaml:"id" json:"id"`     // ID
	IsOn bool  `yaml:"isOn" json:"isOn"` // 是否开启

	// 握手超时时间
	HandshakeTimeout *shared.TimeDuration `yaml:"handshakeTimeout" json:"handshakeTimeout"`

	// 允许的域名，支持 www.example.com, example.com, .example.com, *.example.com
	AllowAllOrigins bool     `yaml:"allowAllOrigins" json:"allowAllOrigins"`
	Origins         []string `yaml:"origins" json:"origins"`

	// 转发方式
	ForwardMode HTTPWebsocketForwardMode `yaml:"forwardMode" json:"forwardMode"`

	handshakeTimeoutDuration time.Duration
}

// 获取新对象
func NewHTTPWebsocketConfig() *HTTPWebsocketConfig {
	return &HTTPWebsocketConfig{
		IsOn: true,
	}
}

// 校验
func (this *HTTPWebsocketConfig) Init() error {
	// duration
	if this.HandshakeTimeout != nil {
		this.handshakeTimeoutDuration = this.HandshakeTimeout.Duration()
	}

	return nil
}

// 获取握手超时时间
func (this *HTTPWebsocketConfig) HandshakeTimeoutDuration() time.Duration {
	return this.handshakeTimeoutDuration
}

// 转发模式名称
func (this *HTTPWebsocketConfig) ForwardModeSummary() maps.Map {
	for _, mode := range AllWebsocketForwardModes() {
		if mode["mode"] == this.ForwardMode {
			return mode
		}
	}
	return nil
}

// 匹配域名
func (this *HTTPWebsocketConfig) MatchOrigin(origin string) bool {
	if this.AllowAllOrigins {
		return true
	}
	return configutils.MatchDomains(this.Origins, origin)
}
