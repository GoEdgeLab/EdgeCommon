package serverconfigs

type HTTPFileCacheStorage struct {
	Dir string `yaml:"dir" json:"dir"` // 目录
}

func (this *HTTPFileCacheStorage) Init() error {
	return nil
}
