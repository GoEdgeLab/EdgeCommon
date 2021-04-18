package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/lists"
	"strings"
)

var DefaultSkippedResponseCacheControlValues = []string{"private", "no-cache", "no-store"}

type HTTPCacheRef struct {
	IsOn          bool  `yaml:"isOn" json:"isOn"`
	CachePolicyId int64 `yaml:"cachePolicyId" json:"cachePolicyId"`

	Key     string               `yaml:"key" json:"key"`         // 每个缓存的Key规则，里面可以有变量
	Life    *shared.TimeDuration `yaml:"life" json:"life"`       // 时间
	Status  []int                `yaml:"status" json:"status"`   // 缓存的状态码列表
	MaxSize *shared.SizeCapacity `yaml:"maxSize" json:"maxSize"` // 能够请求的最大尺寸

	SkipResponseCacheControlValues []string `yaml:"skipCacheControlValues" json:"skipCacheControlValues"`     // 可以跳过的响应的Cache-Control值
	SkipResponseSetCookie          bool     `yaml:"skipSetCookie" json:"skipSetCookie"`                       // 是否跳过响应的Set-Cookie Header
	EnableRequestCachePragma       bool     `yaml:"enableRequestCachePragma" json:"enableRequestCachePragma"` // 是否支持客户端的Pragma: no-cache
	AllowChunkedEncoding           bool     `yaml:"allowChunkedEncoding" json:"allowChunkedEncoding"`         // 是否允许分片内容

	Conds *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"` // 请求条件

	CachePolicy *HTTPCachePolicy `yaml:"cachePolicy" json:"cachePolicy"`

	lifeSeconds                     int64
	maxSize                         int64
	uppercaseSkipCacheControlValues []string
}

func (this *HTTPCacheRef) Init() error {
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

	return nil
}

// 最大数据尺寸
func (this *HTTPCacheRef) MaxSizeBytes() int64 {
	return this.maxSize
}

// 生命周期
func (this *HTTPCacheRef) LifeSeconds() int64 {
	return this.lifeSeconds
}

// 是否包含某个Cache-Control值
func (this *HTTPCacheRef) ContainsCacheControl(value string) bool {
	return lists.ContainsString(this.uppercaseSkipCacheControlValues, strings.ToUpper(value))
}
