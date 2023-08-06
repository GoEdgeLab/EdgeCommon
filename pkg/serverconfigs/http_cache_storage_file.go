package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type CacheDir struct {
	Path     string               `yaml:"path" json:"path"`         // 目录
	Capacity *shared.SizeCapacity `yaml:"capacity" json:"capacity"` // 容量限制 TODO 暂时不实现
}

// HTTPFileCacheStorage 文件缓存存储策略
type HTTPFileCacheStorage struct {
	Dir          string               `yaml:"dir" json:"dir"`                   // 目录
	MinFreeSize  *shared.SizeCapacity `yaml:"minFreeSize" json:"minFreeSize"`   // 最小剩余空间
	SubDirs      []*CacheDir          `yaml:"cacheDir" json:"subDirs"`          // 子目录
	MemoryPolicy *HTTPCachePolicy     `yaml:"memoryPolicy" json:"memoryPolicy"` // 内存二级缓存

	OpenFileCache  *OpenFileCacheConfig `yaml:"openFileCache" json:"openFileCache"`   // open file cache配置
	EnableSendfile bool                 `yaml:"enableSendFile" json:"enableSendfile"` // 是否启用Sendfile
}

func (this *HTTPFileCacheStorage) Init() error {
	if this.OpenFileCache != nil {
		err := this.OpenFileCache.Init()
		if err != nil {
			return err
		}
	}

	if this.MemoryPolicy != nil {
		err := this.MemoryPolicy.Init()
		if err != nil {
			return err
		}
	}

	return nil
}
