package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"regexp"
	"strconv"
	"strings"
)

type HTTPLocationConfig struct {
	Id              int64                          `yaml:"id" json:"id"`                           // ID
	IsOn            bool                           `yaml:"isOn" json:"isOn"`                       // 是否启用
	Pattern         string                         `yaml:"pattern" json:"pattern"`                 // 匹配规则 TODO 未来支持更多样的匹配规则
	Name            string                         `yaml:"name" json:"name"`                       // 名称
	Web             *HTTPWebConfig                 `yaml:"web" json:"web"`                         // Web配置
	URLPrefix       string                         `yaml:"urlPrefix" json:"urlPrefix"`             // 实际的URL前缀，TODO 未来支持变量
	Description     string                         `yaml:"description" json:"description"`         // 描述
	ReverseProxyRef *ReverseProxyRef               `yaml:"reverseProxyRef" json:"reverseProxyRef"` // 反向代理引用
	ReverseProxy    *ReverseProxyConfig            `yaml:"reverseProxy" json:"reverseProxy"`       // 反向代理设置
	IsBreak         bool                           `yaml:"isBreak" json:"isBreak"`                 // 终止向下解析
	Children        []*HTTPLocationConfig          `yaml:"children" json:"children"`               // 子规则
	Conds           *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"`                     // 匹配条件 TODO

	patternType HTTPLocationPatternType // 规则类型：LocationPattern*
	prefix      string                  // 前缀
	path        string                  // 精确的路径

	reg             *regexp.Regexp // 匹配规则
	caseInsensitive bool           // 大小写不敏感
	reverse         bool           // 是否翻转规则，比如非前缀，非路径
}

func (this *HTTPLocationConfig) Init() error {
	err := this.parsePattern()
	if err != nil {
		return err
	}

	if this.Web != nil {
		err := this.Web.Init()
		if err != nil {
			return err
		}
	}

	if this.ReverseProxyRef != nil {
		err := this.ReverseProxyRef.Init()
		if err != nil {
			return err
		}
	}

	if this.ReverseProxy != nil {
		err := this.ReverseProxy.Init()
		if err != nil {
			return err
		}
	}

	// Children
	for _, child := range this.Children {
		err := child.Init()
		if err != nil {
			return err
		}
	}

	// conds
	if this.Conds != nil {
		err := this.Conds.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 组合参数为一个字符串
func (this *HTTPLocationConfig) SetPattern(pattern string, patternType int, caseInsensitive bool, reverse bool) {
	op := ""
	if patternType == HTTPLocationPatternTypePrefix {
		if caseInsensitive {
			op = "*"
			if reverse {
				op = "!*"
			}
		} else {
			if reverse {
				op = "!"
			}
		}
	} else if patternType == HTTPLocationPatternTypeExact {
		op = "="
		if caseInsensitive {
			op += "*"
		}
		if reverse {
			op = "!" + op
		}
	} else if patternType == HTTPLocationPatternTypeRegexp {
		op = "~"
		if caseInsensitive {
			op += "*"
		}
		if reverse {
			op = "!" + op
		}
	}
	if len(op) > 0 {
		pattern = op + " " + pattern
	}
	this.Pattern = pattern
}

// 模式类型
func (this *HTTPLocationConfig) PatternType() int {
	return this.patternType
}

// 模式字符串
// 去掉了模式字符
func (this *HTTPLocationConfig) PatternString() string {
	if this.patternType == HTTPLocationPatternTypePrefix {
		return this.prefix
	}
	return this.path
}

// 是否翻转
func (this *HTTPLocationConfig) IsReverse() bool {
	return this.reverse
}

// 是否大小写非敏感
func (this *HTTPLocationConfig) IsCaseInsensitive() bool {
	return this.caseInsensitive
}

// 分析匹配条件
func (this *HTTPLocationConfig) parsePattern() error {
	// 分析pattern
	this.reverse = false
	this.caseInsensitive = false
	if len(this.Pattern) > 0 {
		spaceIndex := strings.Index(this.Pattern, " ")
		if spaceIndex < 0 {
			this.patternType = HTTPLocationPatternTypePrefix
			this.prefix = this.Pattern
		} else {
			cmd := this.Pattern[:spaceIndex]
			pattern := strings.TrimSpace(this.Pattern[spaceIndex+1:])
			if cmd == "*" { // 大小写非敏感
				this.patternType = HTTPLocationPatternTypePrefix
				this.prefix = pattern
				this.caseInsensitive = true
			} else if cmd == "!*" { // 大小写非敏感，翻转
				this.patternType = HTTPLocationPatternTypePrefix
				this.prefix = pattern
				this.caseInsensitive = true
				this.reverse = true
			} else if cmd == "!" {
				this.patternType = HTTPLocationPatternTypePrefix
				this.prefix = pattern
				this.reverse = true
			} else if cmd == "=" {
				this.patternType = HTTPLocationPatternTypeExact
				this.path = pattern
			} else if cmd == "=*" {
				this.patternType = HTTPLocationPatternTypeExact
				this.path = pattern
				this.caseInsensitive = true
			} else if cmd == "!=" {
				this.patternType = HTTPLocationPatternTypeExact
				this.path = pattern
				this.reverse = true
			} else if cmd == "!=*" {
				this.patternType = HTTPLocationPatternTypeExact
				this.path = pattern
				this.reverse = true
				this.caseInsensitive = true
			} else if cmd == "~" { // 正则
				this.patternType = HTTPLocationPatternTypeRegexp
				reg, err := regexp.Compile(pattern)
				if err != nil {
					return err
				}
				this.reg = reg
				this.path = pattern
			} else if cmd == "!~" {
				this.patternType = HTTPLocationPatternTypeRegexp
				reg, err := regexp.Compile(pattern)
				if err != nil {
					return err
				}
				this.reg = reg
				this.reverse = true
				this.path = pattern
			} else if cmd == "~*" { // 大小写非敏感小写
				this.patternType = HTTPLocationPatternTypeRegexp
				reg, err := regexp.Compile("(?i)" + pattern)
				if err != nil {
					return err
				}
				this.reg = reg
				this.caseInsensitive = true
				this.path = pattern
			} else if cmd == "!~*" {
				this.patternType = HTTPLocationPatternTypeRegexp
				reg, err := regexp.Compile("(?i)" + pattern)
				if err != nil {
					return err
				}
				this.reg = reg
				this.reverse = true
				this.caseInsensitive = true
				this.path = pattern
			} else {
				this.patternType = HTTPLocationPatternTypePrefix
				this.prefix = pattern
			}
		}
	} else {
		this.patternType = HTTPLocationPatternTypePrefix
		this.prefix = this.Pattern
	}

	return nil
}

// 判断是否匹配路径
// TODO 支持子Location
func (this *HTTPLocationConfig) Match(path string, formatter func(source string) string) (vars map[string]string, isMatched bool) {
	// 判断条件
	if this.Conds != nil && !this.Conds.MatchRequest(formatter) {
		return
	}

	if this.patternType == HTTPLocationPatternTypePrefix {
		if this.reverse {
			if this.caseInsensitive {
				return nil, !strings.HasPrefix(strings.ToLower(path), strings.ToLower(this.prefix))
			} else {
				return nil, !strings.HasPrefix(path, this.prefix)
			}
		} else {
			if this.caseInsensitive {
				return nil, strings.HasPrefix(strings.ToLower(path), strings.ToLower(this.prefix))
			} else {
				return nil, strings.HasPrefix(path, this.prefix)
			}
		}
	}

	if this.patternType == HTTPLocationPatternTypeExact {
		if this.reverse {
			if this.caseInsensitive {
				return nil, strings.ToLower(path) != strings.ToLower(this.path)
			} else {
				return nil, path != this.path
			}
		} else {
			if this.caseInsensitive {
				return nil, strings.ToLower(path) == strings.ToLower(this.path)
			} else {
				return nil, path == this.path
			}
		}
	}

	// TODO 正则表达式匹配会让请求延迟0.01-0.02ms，可以使用缓存加速正则匹配，因为大部分路径都是不变的
	if this.patternType == HTTPLocationPatternTypeRegexp {
		if this.reg != nil {
			if this.reverse {
				return nil, !this.reg.MatchString(path)
			} else {
				b := this.reg.MatchString(path)
				if b {
					result := map[string]string{}
					matches := this.reg.FindStringSubmatch(path)
					subNames := this.reg.SubexpNames()
					for index, value := range matches {
						result[strconv.Itoa(index)] = value
						subName := subNames[index]
						if len(subName) > 0 {
							result[subName] = value
						}
					}
					return result, true
				}
				return nil, b
			}
		}

		return nil, this.reverse
	}

	return nil, false
}
