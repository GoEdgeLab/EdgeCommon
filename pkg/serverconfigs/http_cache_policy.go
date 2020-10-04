package serverconfigs

import (
	"bytes"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
)

// 缓存策略配置
type HTTPCachePolicy struct {
	Id          int64                  `yaml:"id" json:"id"`
	IsOn        bool                   `yaml:"isOn" json:"isOn"`               // 是否开启
	Name        string                 `yaml:"name" json:"name"`               // 名称
	Description string                 `yaml:"description" json:"description"` // 描述
	Capacity    *shared.SizeCapacity   `yaml:"capacity" json:"capacity"`       // 最大内容容量 TODO 需要实现
	MaxKeys     int64                  `yaml:"maxKeys" json:"maxKeys"`         // 最多Key值 TODO 需要实现
	MaxSize     *shared.SizeCapacity   `yaml:"maxSize" json:"maxSize"`         // 单个缓存最大尺寸 TODO 需要实现
	Type        CachePolicyStorageType `yaml:"type" json:"type"`               // 类型
	Options     map[string]interface{} `yaml:"options" json:"options"`         // 选项

	capacity int64
}

// 校验
func (this *HTTPCachePolicy) Init() error {
	var err error

	if this.Capacity != nil {
		this.capacity = this.Capacity.Bytes()
	}

	return err
}

// 容量
func (this *HTTPCachePolicy) CapacitySize() int64 {
	return this.capacity
}

// 对比Policy是否有变化
func (this *HTTPCachePolicy) IsSame(anotherPolicy *HTTPCachePolicy) bool {
	policyJSON1, err := json.Marshal(this)
	if err != nil {
		return false
	}
	policyJSON2, err := json.Marshal(anotherPolicy)
	if err != nil {
		return false
	}
	return bytes.Equal(policyJSON1, policyJSON2)
}
