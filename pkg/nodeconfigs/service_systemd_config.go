package nodeconfigs

type (
	SystemServiceType = string
)

const (
	// TODO 需要支持supervisor等常用daemon管理工具
	SystemServiceTypeSystemd SystemServiceType = "systemd"
)

// Systemd配置
type SystemdServiceConfig struct {
	IsOn             bool   `yaml:"isOn" json:"isOn"`                         // 是否启用
	Provides         string `yaml:"provides" json:"provides"`                 // 提供者，可以是服务名
	ShortDescription string `yaml:"shortDescription" json:"shortDescription"` // 短描述
	Description      string `yaml:"description" json:"description"`           // 长描述
	ExecPath         string `yaml:"execPath" json:"execPath"`                 // 可执行文件的路径
}
