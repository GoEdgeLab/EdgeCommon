package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/lists"
	"strings"
)

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

// 格式化域名
func (this *ServerNameConfig) Normalize() {
	this.Name = strings.ToLower(this.Name)
	for index, subName := range this.SubNames {
		this.SubNames[index] = strings.ToLower(subName)
	}
}

// 判断主机名是否匹配
func (this *ServerNameConfig) Match(name string) bool {
	if len(this.SubNames) > 0 {
		return configutils.MatchDomains(this.SubNames, name)
	}
	return configutils.MatchDomains([]string{this.Name}, name)
}

// 格式化一组域名
func NormalizeServerNames(serverNames []*ServerNameConfig) {
	for _, serverName := range serverNames {
		serverName.Normalize()
	}
}

// 获取所有域名
func PlainServerNames(serverNames []*ServerNameConfig) (result []string) {
	NormalizeServerNames(serverNames)
	for _, serverName := range serverNames {
		if len(serverName.SubNames) == 0 {
			if len(serverName.Name) > 0 && !lists.ContainsString(result, serverName.Name) {
				result = append(result, serverName.Name)
			}
		} else {
			for _, subName := range serverName.SubNames {
				if len(subName) > 0 && !lists.ContainsString(result, subName) {
					result = append(result, subName)
				}
			}
		}
	}
	return result
}
