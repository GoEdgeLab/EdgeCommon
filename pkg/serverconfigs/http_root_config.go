package serverconfigs

// Web文档目录配置
type HTTPRootConfig struct {
	IsPrior     bool     `yaml:"isPrior" json:"isPrior"`         // 是否优先
	IsOn        bool     `yaml:"isOn" json:"isOn"`               // 是否启用
	Dir         string   `yaml:"dir" json:"dir"`                 // 目录
	Indexes     []string `yaml:"indexes" json:"indexes"`         // 默认首页文件
	StripPrefix string   `yaml:"stripPrefix" json:"stripPrefix"` // 去除前缀
	DecodePath  bool     `yaml:"decodePath" json:"decodePath"`   // 是否对请求路径进行解码
	IsBreak     bool     `yaml:"isBreak" json:"isBreak"`         // 找不到文件的情况下是否终止
}

// 初始化
func (ths *HTTPRootConfig) Init() error {
	return nil
}
