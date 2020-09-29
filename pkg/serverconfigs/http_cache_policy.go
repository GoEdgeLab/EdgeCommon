package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/files"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/logs"
	"strings"
	"time"
)

var DefaultSkippedResponseCacheControlValues = []string{"private", "no-cache", "no-store"}

// 缓存策略配置
type HTTPCachePolicy struct {
	Id   int64  `yaml:"id" json:"id"`
	IsOn bool   `yaml:"isOn" json:"isOn"` // 是否开启 TODO
	Name string `yaml:"name" json:"name"` // 名称

	Key      string               `yaml:"key" json:"key"`           // 每个缓存的Key规则，里面可以有变量
	Capacity *shared.SizeCapacity `yaml:"capacity" json:"capacity"` // 最大内容容量
	Life     *shared.TimeDuration `yaml:"life" json:"life"`         // 时间
	Status   []int                `yaml:"status" json:"status"`     // 缓存的状态码列表
	MaxSize  *shared.SizeCapacity `yaml:"maxSize" json:"maxSize"`   // 能够请求的最大尺寸

	SkipResponseCacheControlValues []string `yaml:"skipCacheControlValues" json:"skipCacheControlValues"`     // 可以跳过的响应的Cache-Control值
	SkipResponseSetCookie          bool     `yaml:"skipSetCookie" json:"skipSetCookie"`                       // 是否跳过响应的Set-Cookie Header
	EnableRequestCachePragma       bool     `yaml:"enableRequestCachePragma" json:"enableRequestCachePragma"` // 是否支持客户端的Pragma: no-cache

	Conds *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"` // 请求条件

	life     time.Duration
	maxSize  int64
	capacity int64

	uppercaseSkipCacheControlValues []string

	Type    string                 `yaml:"type" json:"type"`       // 类型
	Options map[string]interface{} `yaml:"options" json:"options"` // 选项
}

// 获取新对象
func NewHTTPCachePolicy() *HTTPCachePolicy {
	return &HTTPCachePolicy{
		SkipResponseCacheControlValues: DefaultSkippedResponseCacheControlValues,
		SkipResponseSetCookie:          true,
	}
}

// 从文件中读取缓存策略
func NewCachePolicyFromFile(file string) *HTTPCachePolicy {
	if len(file) == 0 {
		return nil
	}
	reader, err := files.NewReader(Tea.ConfigFile(file))
	if err != nil {
		logs.Error(err)
		return nil
	}
	defer func() {
		_ = reader.Close()
	}()

	p := NewHTTPCachePolicy()
	err = reader.ReadYAML(p)
	if err != nil {
		logs.Error(err)
		return nil
	}

	return p
}

// 校验
func (this *HTTPCachePolicy) Init() error {
	var err error
	this.maxSize = this.MaxSize.Bytes()
	this.life = this.Life.Duration()
	this.capacity = this.Capacity.Bytes()

	this.uppercaseSkipCacheControlValues = []string{}
	for _, value := range this.SkipResponseCacheControlValues {
		this.uppercaseSkipCacheControlValues = append(this.uppercaseSkipCacheControlValues, strings.ToUpper(value))
	}

	// cond
	if this.Conds != nil {
		err := this.Conds.Init()
		if err != nil {
			return err
		}
	}

	return err
}

// 最大数据尺寸
func (this *HTTPCachePolicy) MaxDataSize() int64 {
	return this.maxSize
}

// 容量
func (this *HTTPCachePolicy) CapacitySize() int64 {
	return this.capacity
}

// 生命周期
func (this *HTTPCachePolicy) LifeDuration() time.Duration {
	return this.life
}

// 是否包含某个Cache-Control值
func (this *HTTPCachePolicy) ContainsCacheControl(value string) bool {
	return lists.ContainsString(this.uppercaseSkipCacheControlValues, strings.ToUpper(value))
}

// 检查是否匹配关键词
func (this *HTTPCachePolicy) MatchKeyword(keyword string) (matched bool, name string, tags []string) {
	if configutils.MatchKeyword(this.Name, keyword) || configutils.MatchKeyword(this.Type, keyword) {
		matched = true
		name = this.Name
		if len(this.Type) > 0 {
			tags = []string{"类型：" + this.Type}
		}
	}
	return
}
