package serverconfigs

import (
	"bytes"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
)

// HTTPCachePolicy 缓存策略配置
type HTTPCachePolicy struct {
	Id                   int64                  `yaml:"id" json:"id"`
	IsOn                 bool                   `yaml:"isOn" json:"isOn"`                                 // 是否开启
	Name                 string                 `yaml:"name" json:"name"`                                 // 名称
	Description          string                 `yaml:"description" json:"description"`                   // 描述
	Capacity             *shared.SizeCapacity   `yaml:"capacity" json:"capacity"`                         // 最大内容容量
	MaxKeys              int64                  `yaml:"maxKeys" json:"maxKeys"`                           // 最多Key值
	MaxSize              *shared.SizeCapacity   `yaml:"maxSize" json:"maxSize"`                           // 单个缓存最大尺寸
	Type                 CachePolicyStorageType `yaml:"type" json:"type"`                                 // 类型
	Options              map[string]interface{} `yaml:"options" json:"options"`                           // 选项
	Life                 *shared.TimeDuration   `yaml:"life" json:"life"`                                 // 默认有效期 TODO 需要实现
	MinLife              *shared.TimeDuration   `yaml:"minLife" json:"minLife"`                           // 最小有效期 TODO 需要实现
	MaxLife              *shared.TimeDuration   `yaml:"maxLife" json:"maxLife"`                           // 最大有效期 TODO 需要实现
	SyncCompressionCache bool                   `yaml:"syncCompressionCache" json:"syncCompressionCache"` // 同步写入压缩缓存

	CacheRefs []*HTTPCacheRef `yaml:"cacheRefs" json:"cacheRefs"` // 缓存配置

	PersistenceAutoPurgeCount    int     `yaml:"persistenceAutoPurgeCount" json:"persistenceAutoPurgeCount"`       // 每次自动清理的条数 TODO 需要实现
	PersistenceAutoPurgeInterval int     `yaml:"persistenceAutoPurgeInterval" json:"persistenceAutoPurgeInterval"` // 自动清理的时间间隔（秒） TODO 需要实现
	PersistenceLFUFreePercent    float32 `yaml:"persistenceLFUFreePercent" json:"persistenceLFUFreePercent"`       // LFU算法执行阈值，剩余空间比例，使用百分比，比如20 TODO 需要实现
	PersistenceHitSampleRate     int     `yaml:"persistenceHitSampleRate" json:"persistenceHitSampleRate"`         // 热点数据采样比例 TODO 需要实现

	MemoryAutoPurgeCount     int     `yaml:"memoryAutoPurgeCount" json:"memoryAutoPurgeCount"`         // 每次自动清理的条数 TODO 需要实现
	MemoryAutoPurgeInterval  int     `yaml:"memoryAutoPurgeInterval" json:"memoryAutoPurgeInterval"`   // 自动清理的时间间隔（秒） TODO 需要实现
	MemoryLFUFreePercent     float32 `yaml:"memoryLFUFreePercent" json:"memoryLFUFreePercent"`         // LFU算法执行阈值，剩余空间比例，使用百分比，比如20 TODO 需要实现
	MemoryAutoFlushQueueSize int     `yaml:"memoryAutoFlushQueueSize" json:"memoryAutoFlushQueueSize"` // 自动刷新到持久层队列尺寸 TODO 需要实现

	capacity int64
	maxSize  int64
}

// Init 校验
func (this *HTTPCachePolicy) Init() error {
	if this.Capacity != nil {
		this.capacity = this.Capacity.Bytes()
	}

	if this.MaxSize != nil {
		this.maxSize = this.MaxSize.Bytes()
	}

	for _, cacheRef := range this.CacheRefs {
		err := cacheRef.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// CapacityBytes 容量
func (this *HTTPCachePolicy) CapacityBytes() int64 {
	return this.capacity
}

// MaxSizeBytes 单个缓存最大尺寸
func (this *HTTPCachePolicy) MaxSizeBytes() int64 {
	return this.maxSize
}

// IsSame 对比Policy是否有变化
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
