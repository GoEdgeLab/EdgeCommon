package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/configutils"

type ServerNameType = string

const (
	ServerNameTypeFull   ServerNameType = "full"   // 完整的域名，包含通配符等
	ServerNameTypePrefix ServerNameType = "prefix" // 前缀
	ServerNameTypeSuffix ServerNameType = "suffix" // 后缀
	ServerNameTypeMatch  ServerNameType = "match"  // 正则匹配
)

// 主机名(域名)配置
type ServerNameConfig struct {
	Name     string   `yaml:"name" json:"name"`         // 名称
	Type     string   `yaml:"type" json:"type"`         // 类型
	SubNames []string `yaml:"subNames" json:"subNames"` // 子名称，用来支持大量的域名批量管理
}

// 判断主机名是否匹配
func (this *ServerNameConfig) Match(name string) bool {
	if len(this.SubNames) > 0 {
		return configutils.MatchDomains(this.SubNames, name)
	}
	return configutils.MatchDomains([]string{this.Name}, name)
}
