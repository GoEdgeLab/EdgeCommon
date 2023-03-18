package serverconfigs

import (
	"context"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"strconv"
	"strings"
	"time"
)

// OriginConfig 源站服务配置
type OriginConfig struct {
	Id          int64                 `yaml:"id" json:"id"`                   // ID
	IsOn        bool                  `yaml:"isOn" json:"isOn"`               // 是否启用
	Version     int                   `yaml:"version" json:"version"`         // 版本
	Name        string                `yaml:"name" json:"name"`               // 名称
	Addr        *NetworkAddressConfig `yaml:"addr" json:"addr"`               // 地址
	Description string                `yaml:"description" json:"description"` // 描述 TODO
	Code        string                `yaml:"code" json:"code"`               // 代号 TODO

	Weight       uint                 `yaml:"weight" json:"weight"`           // 权重
	ConnTimeout  *shared.TimeDuration `yaml:"connTimeout" json:"connTimeout"` // 连接失败超时
	ReadTimeout  *shared.TimeDuration `yaml:"readTimeout" json:"readTimeout"` // 读取超时时间
	IdleTimeout  *shared.TimeDuration `yaml:"idleTimeout" json:"idleTimeout"` // 空闲连接超时时间
	MaxFails     int                  `yaml:"maxFails" json:"maxFails"`       // 最多失败次数 TODO
	MaxConns     int                  `yaml:"maxConns" json:"maxConns"`       // 最大并发连接数
	MaxIdleConns int                  `yaml:"idleConns" json:"idleConns"`     // 最大空闲连接数

	Domains []string `yaml:"domains" json:"domains"` // 所属域名

	StripPrefix    string                `yaml:"stripPrefix" json:"stripPrefix"`       // 去除URL前缀
	RequestURI     string                `yaml:"requestURI" json:"requestURI"`         // 转发后的请求URI TODO
	RequestHost    string                `yaml:"requestHost" json:"requestHost"`       // 自定义主机名
	FollowPort     bool                  `yaml:"followPort" json:"followPort"`         // 端口跟随
	FollowProtocol *FollowProtocolConfig `yaml:"followProtocol" json:"followProtocol"` // 协议跟随 TODO

	RequestHeaderPolicyRef  *shared.HTTPHeaderPolicyRef `yaml:"requestHeaderPolicyRef" json:"requestHeaderPolicyRef"`   // 请求Header
	RequestHeaderPolicy     *shared.HTTPHeaderPolicy    `yaml:"requestHeaderPolicy" json:"requestHeaderPolicy"`         // 请求Header策略
	ResponseHeaderPolicyRef *shared.HTTPHeaderPolicyRef `yaml:"responseHeaderPolicyRef" json:"responseHeaderPolicyRef"` // 响应Header`
	ResponseHeaderPolicy    *shared.HTTPHeaderPolicy    `yaml:"responseHeaderPolicy" json:"responseHeaderPolicy"`       // 响应Header策略

	// 健康检查URL，目前支持：
	// - http|https 返回2xx-3xx认为成功
	HealthCheck *HealthCheckConfig `yaml:"healthCheck" json:"healthCheck"` // TODO

	CertRef *sslconfigs.SSLCertRef    `yaml:"certRef" json:"certRef"` // 证书的引用
	Cert    *sslconfigs.SSLCertConfig `yaml:"cert" json:"cert"`       // 请求源服务器用的证书

	// ftp
	FTPServerRef *FTPServerRef    `yaml:"ftpServerRef" json:"ftpServerRef"` // TODO
	FTPServer    *FTPServerConfig `yaml:"ftpServer" json:"ftpServer"`       // TODO

	IsOk bool `yaml:"isOk" json:"isOk"` // 是否可以正常访问，仅仅用于运行时检查

	connTimeoutDuration time.Duration
	readTimeoutDuration time.Duration
	idleTimeoutDuration time.Duration

	hasRequestURI bool
	requestPath   string
	requestArgs   string

	hasRequestHeaders  bool
	hasResponseHeaders bool

	uniqueKey string

	requestHostHasVariables bool
	requestURIHasVariables  bool
}

// Init 校验
func (this *OriginConfig) Init(ctx context.Context) error {
	this.IsOk = true

	// URL
	this.requestHostHasVariables = configutils.HasVariables(this.RequestHost)
	this.requestURIHasVariables = configutils.HasVariables(this.RequestURI)

	// unique key
	this.uniqueKey = strconv.FormatInt(this.Id, 10) + "@" + strconv.Itoa(this.Version) + "@" + fmt.Sprintf("%p", this)

	// addr
	if this.Addr != nil {
		err := this.Addr.Init()
		if err != nil {
			return err
		}
	}

	// 证书
	if this.Cert != nil {
		err := this.Cert.Init(ctx)
		if err != nil {
			return err
		}
	}

	// connTimeout
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

	// health check
	if this.HealthCheck != nil {
		err := this.HealthCheck.Init()
		if err != nil {
			return err
		}
	}

	// follow protocol
	if this.FollowProtocol != nil {
		err := this.FollowProtocol.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// CandidateCodes 候选对象代号
func (this *OriginConfig) CandidateCodes() []string {
	codes := []string{strconv.FormatInt(this.Id, 10)}
	if len(this.Code) > 0 {
		codes = append(codes, this.Code)
	}
	return codes
}

// CandidateWeight 候选对象权重
func (this *OriginConfig) CandidateWeight() uint {
	return this.Weight
}

// ConnTimeoutDuration 连接超时时间
func (this *OriginConfig) ConnTimeoutDuration() time.Duration {
	return this.connTimeoutDuration
}

// ReadTimeoutDuration 读取超时时间
func (this *OriginConfig) ReadTimeoutDuration() time.Duration {
	return this.readTimeoutDuration
}

// IdleTimeoutDuration 休眠超时时间
func (this *OriginConfig) IdleTimeoutDuration() time.Duration {
	return this.idleTimeoutDuration
}

// RequestHostHasVariables 判断RequestHost是否有变量
func (this *OriginConfig) RequestHostHasVariables() bool {
	return this.requestHostHasVariables
}

// RequestURIHasVariables 判断RequestURI是否有变量
func (this *OriginConfig) RequestURIHasVariables() bool {
	return this.requestURIHasVariables
}

// UniqueKey 唯一Key
func (this *OriginConfig) UniqueKey() string {
	return this.uniqueKey
}
