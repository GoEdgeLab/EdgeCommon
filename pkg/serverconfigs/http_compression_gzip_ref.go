package serverconfigs

type HTTPGzipRef struct {
	IsPrior bool  `yaml:"isPrior" json:"isPrior"` // 是否覆盖
	IsOn    bool  `yaml:"isOn" json:"isOn"`       // 是否开启
	GzipId  int64 `yaml:"gzipId" json:"gzipId"`   // 使用的配置ID
}
