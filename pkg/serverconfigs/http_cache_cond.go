package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/lists"
	"strings"
	"time"
)

var DefaultSkippedResponseCacheControlValues = []string{"private", "no-cache", "no-store"}

type HTTPCacheCond struct {
	Key     string               `yaml:"key" json:"key"`         // 每个缓存的Key规则，里面可以有变量
	Life    *shared.TimeDuration `yaml:"life" json:"life"`       // 时间
	Status  []int                `yaml:"status" json:"status"`   // 缓存的状态码列表
	MaxSize *shared.SizeCapacity `yaml:"maxSize" json:"maxSize"` // 能够请求的最大尺寸

	SkipResponseCacheControlValues []string `yaml:"skipCacheControlValues" json:"skipCacheControlValues"`     // 可以跳过的响应的Cache-Control值
	SkipResponseSetCookie          bool     `yaml:"skipSetCookie" json:"skipSetCookie"`                       // 是否跳过响应的Set-Cookie Header
	EnableRequestCachePragma       bool     `yaml:"enableRequestCachePragma" json:"enableRequestCachePragma"` // 是否支持客户端的Pragma: no-cache

	Conds *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"` // 请求条件

	life                            time.Duration
	maxSize                         int64
	uppercaseSkipCacheControlValues []string
}

func (this *HTTPCacheCond) Init() error {
	if this.MaxSize != nil {
		this.maxSize = this.MaxSize.Bytes()
	}
	if this.Life != nil {
		this.life = this.Life.Duration()
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

	return nil
}

// 最大数据尺寸
func (this *HTTPCacheCond) MaxDataSize() int64 {
	return this.maxSize
}

// 生命周期
func (this *HTTPCacheCond) LifeDuration() time.Duration {
	return this.life
}

// 是否包含某个Cache-Control值
func (this *HTTPCacheCond) ContainsCacheControl(value string) bool {
	return lists.ContainsString(this.uppercaseSkipCacheControlValues, strings.ToUpper(value))
}
