package serverconfigs

import (
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"strconv"
	"strings"
	"time"
)

// 源站服务配置
type OriginConfig struct {
	Id          int64                 `yaml:"id" json:"id"`                   // ID
	IsOn        bool                  `yaml:"isOn" json:"isOn"`               // 是否启用 TODO
	Version     int                   `yaml:"version" json:"version"`         // 版本
	Name        string                `yaml:"name" json:"name"`               // 名称 TODO
	Addr        *NetworkAddressConfig `yaml:"addr" json:"addr"`               // 地址
	Description string                `yaml:"description" json:"description"` // 描述 TODO
	Code        string                `yaml:"code" json:"code"`               // 代号 TODO

	Weight       uint                 `yaml:"weight" json:"weight"`           // 权重 TODO
	ConnTimeout  *shared.TimeDuration `yaml:"failTimeout" json:"failTimeout"` // 连接失败超时 TODO
	ReadTimeout  *shared.TimeDuration `yaml:"readTimeout" json:"readTimeout"` // 读取超时时间 TODO
	IdleTimeout  *shared.TimeDuration `yaml:"idleTimeout" json:"idleTimeout"` // 空闲连接超时时间 TODO
	MaxFails     int                  `yaml:"maxFails" json:"maxFails"`       // 最多失败次数 TODO
	MaxConns     int                  `yaml:"maxConns" json:"maxConns"`       // 最大并发连接数 TODO
	MaxIdleConns int                  `yaml:"idleConns" json:"idleConns"`     // 最大空闲连接数 TODO

	RequestURI string `yaml:"requestURI" json:"requestURI"` // 转发后的请求URI TODO
	Host       string `yaml:"host" json:"host"`             // 自定义主机名 TODO

	RequestHeaderPolicyRef  *shared.HTTPHeaderPolicyRef `yaml:"requestHeaderPolicyRef" json:"requestHeaderPolicyRef"`   // 请求Header
	RequestHeaderPolicy     *shared.HTTPHeaderPolicy    `yaml:"requestHeaderPolicy" json:"requestHeaderPolicy"`         // 请求Header策略
	ResponseHeaderPolicyRef *shared.HTTPHeaderPolicyRef `yaml:"responseHeaderPolicyRef" json:"responseHeaderPolicyRef"` // 响应Header`
	ResponseHeaderPolicy    *shared.HTTPHeaderPolicy    `yaml:"responseHeaderPolicy" json:"responseHeaderPolicy"`       // 响应Header策略

	// 健康检查URL，目前支持：
	// - http|https 返回2xx-3xx认为成功
	HealthCheck *HealthCheckConfig `yaml:"healthCheck" json:"healthCheck"`

	Cert *sslconfigs.SSLCertConfig `yaml:"cert" json:"cert"` // 请求源服务器用的证书

	// ftp
	FTP *OriginFTPConfig `yaml:"ftp" json:"ftp"`

	connTimeoutDuration time.Duration
	readTimeoutDuration time.Duration
	idleTimeoutDuration time.Duration

	hasRequestURI bool
	requestPath   string
	requestArgs   string

	hasRequestHeaders  bool
	hasResponseHeaders bool

	hasHost bool

	uniqueKey string

	hasAddrVariables bool // 地址中是否含有变量

	realAddr string // 最终的Addr TODO
}

// 校验
func (this *OriginConfig) Init() error {
	// 证书
	if this.Cert != nil {
		err := this.Cert.Init()
		if err != nil {
			return err
		}
	}

	// unique key
	this.uniqueKey = strconv.FormatInt(this.Id, 10) + "@" + fmt.Sprintf("%d", this.Version)

	// failTimeout
	if this.ConnTimeout != nil {
		this.connTimeoutDuration = this.ConnTimeout.Duration()
	}

	// readTimeout
	if this.ReadTimeout != nil {
		this.readTimeoutDuration = this.ReadTimeout.Duration()
	}

	// idleTimeout
	if this.IdleTimeout != nil {
		this.idleTimeoutDuration = this.IdleTimeout.Duration()
	}

	// Headers
	if this.RequestHeaderPolicyRef != nil {
		err := this.RequestHeaderPolicyRef.Init()
		if err != nil {
			return err
		}
	}
	if this.RequestHeaderPolicy != nil {
		err := this.RequestHeaderPolicy.Init()
		if err != nil {
			return err
		}
	}
	if this.ResponseHeaderPolicyRef != nil {
		err := this.ResponseHeaderPolicyRef.Init()
		if err != nil {
			return err
		}
	}
	if this.ResponseHeaderPolicy != nil {
		err := this.ResponseHeaderPolicy.Init()
		if err != nil {
			return err
		}
	}

	// request uri
	if len(this.RequestURI) == 0 || this.RequestURI == "${requestURI}" {
		this.hasRequestURI = false
	} else {
		this.hasRequestURI = true

		if strings.Contains(this.RequestURI, "?") {
			pieces := strings.SplitN(this.RequestURI, "?", -1)
			this.requestPath = pieces[0]
			this.requestArgs = pieces[1]
		} else {
			this.requestPath = this.RequestURI
		}
	}

	// TODO init health check

	// host
	this.hasHost = len(this.Host) > 0

	// variables
	// TODO 在host和port中支持变量
	this.hasAddrVariables = false

	return nil
}

// 候选对象代号
func (this *OriginConfig) CandidateCodes() []string {
	codes := []string{strconv.FormatInt(this.Id, 10)}
	if len(this.Code) > 0 {
		codes = append(codes, this.Code)
	}
	return codes
}

// 候选对象权重
func (this *OriginConfig) CandidateWeight() uint {
	return this.Weight
}

// 获取最终请求的地址
func (this *OriginConfig) RealAddr() string {
	return this.realAddr
}

// 设置最终请求的地址 TODO 需要实现
func (this *OriginConfig) SetRealAddr(realAddr string) {
	this.realAddr = realAddr
}

// 连接超时时间
func (this *OriginConfig) ConnTimeoutDuration() time.Duration {
	return this.connTimeoutDuration
}
