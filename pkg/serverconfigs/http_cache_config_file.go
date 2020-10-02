package serverconfigs

type CachePolicyType = string

type HTTPFileCacheConfig struct {
	Dir string `yaml:"dir" json:"dir"` // 目录
}

func (this *HTTPFileCacheConfig) Init() error {
	return nil
}
