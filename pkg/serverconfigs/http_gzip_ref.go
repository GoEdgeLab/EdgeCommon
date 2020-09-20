package serverconfigs

type HTTPGzipRef struct {
	IsOn   bool  `yaml:"isOn" json:"isOn"`
	GzipId int64 `yaml:"gzipId" json:"gzipId"`
}
