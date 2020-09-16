package serverconfigs

type HTTPWebConfig struct {
	Id        int64             `yaml:"id" json:"id"`               // ID
	IsOn      bool              `yaml:"isOn" json:"isOn"`           // 是否启用
	Locations []*LocationConfig `yaml:"locations" json:"locations"` // 路径规则 TODO
	Gzip      *HTTPGzipConfig   `yaml:"gzip" json:"gzip"`           // Gzip配置

	// 本地静态资源配置
	Root string `yaml:"root" json:"root"` // 资源根目录 TODO
}
