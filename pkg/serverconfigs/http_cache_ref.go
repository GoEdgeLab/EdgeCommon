package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/lists"
	"net/http"
	"strings"
)

var DefaultSkippedResponseCacheControlValues = []string{"private", "no-cache", "no-store"}

type HTTPCacheRef struct {
	IsOn          bool  `yaml:"isOn" json:"isOn"`
	CachePolicyId int64 `yaml:"cachePolicyId" json:"cachePolicyId"`

	Key         string                 `yaml:"key" json:"key"`                 // 每个缓存的Key规则，里面可以有变量
	Life        *shared.TimeDuration   `yaml:"life" json:"life"`               // 时间
	ExpiresTime *HTTPExpiresTimeConfig `yaml:"expiresTime" json:"expiresTime"` // 客户端过期时间
	Status      []int                  `yaml:"status" json:"status"`           // 缓存的状态码列表
	MinSize     *shared.SizeCapacity   `yaml:"minSize" json:"minSize"`         // 能够缓存的最小尺寸
	MaxSize     *shared.SizeCapacity   `yaml:"maxSize" json:"maxSize"`         // 能够缓存的最大尺寸
	Methods     []string               `yaml:"methods" json:"methods"`         // 支持的请求方法

	SkipResponseCacheControlValues []string `yaml:"skipCacheControlValues" json:"skipCacheControlValues"`     // 可以跳过的响应的Cache-Control值
	SkipResponseSetCookie          bool     `yaml:"skipSetCookie" json:"skipSetCookie"`                       // 是否跳过响应的Set-Cookie Header
	EnableRequestCachePragma       bool     `yaml:"enableRequestCachePragma" json:"enableRequestCachePragma"` // 是否支持客户端的Pragma: no-cache
	AllowChunkedEncoding           bool     `yaml:"allowChunkedEncoding" json:"allowChunkedEncoding"`         // 是否允许分片内容
	AllowPartialContent            bool     `yaml:"allowPartialContent" json:"allowPartialContent"`           // 支持分段内容缓存

	Conds *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"` // 请求条件

	CachePolicy *HTTPCachePolicy `yaml:"cachePolicy" json:"cachePolicy"`

	IsReverse bool `yaml:"isReverse" json:"isReverse"` // 是否为反向条件，反向条件的不缓存

	lifeSeconds                     int64
	minSize                         int64
	maxSize                         int64
	uppercaseSkipCacheControlValues []string

	methodMap map[string]bool
}

func (this *HTTPCacheRef) Init() error {
	if this.MinSize != nil {
		this.minSize = this.MinSize.Bytes()
	}
	if this.MaxSize != nil {
		this.maxSize = this.MaxSize.Bytes()
	}
	if this.Life != nil {
		this.lifeSeconds = int64(this.Life.Duration().Seconds())
	}

	// control-values
	this.uppercaseSkipCacheControlValues = []string{}
	for _, value := range this.SkipResponseCacheControlValues {
		this.uppercaseSkipCacheControlValues = append(this.uppercaseSkipCacheControlValues, strings.ToUpper(value))
	}

	// conds
	if this.Conds != nil {
		err := this.Conds.Init()
		if err != nil {
			return err
		}
	}

	// policy
	if this.CachePolicy != nil {
		err := this.CachePolicy.Init()
		if err != nil {
			return err
		}
	}

	// methods
	this.methodMap = map[string]bool{}
	if len(this.Methods) > 0 {
		for _, method := range this.Methods {
			this.methodMap[strings.ToUpper(method)] = true
		}
	}

	return nil
}

// MaxSizeBytes 最大数据尺寸
func (this *HTTPCacheRef) MaxSizeBytes() int64 {
	return this.maxSize
}

// MinSizeBytes 最小数据尺寸
func (this *HTTPCacheRef) MinSizeBytes() int64 {
	return this.minSize
}

// LifeSeconds 生命周期
func (this *HTTPCacheRef) LifeSeconds() int64 {
	return this.lifeSeconds
}

// ContainsCacheControl 是否包含某个Cache-Control值
func (this *HTTPCacheRef) ContainsCacheControl(value string) bool {
	return lists.ContainsString(this.uppercaseSkipCacheControlValues, strings.ToUpper(value))
}

// MatchRequest 匹配请求
func (this *HTTPCacheRef) MatchRequest(req *http.Request) bool {
	// 请求方法
	if len(this.methodMap) > 0 {
		_, ok := this.methodMap[req.Method]
		if !ok {
			return false
		}
	}

	return true
}
