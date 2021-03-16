package serverconfigs

// 文件缓存存储策略
type HTTPFileCacheStorage struct {
	Dir          string           `yaml:"dir" json:"dir"`                   // 目录
	MemoryPolicy *HTTPCachePolicy `yaml:"memoryPolicy" json:"memoryPolicy"` // 内存二级缓存
}

func (this *HTTPFileCacheStorage) Init() error {
	return nil
}
