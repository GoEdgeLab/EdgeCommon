package serverconfigs

type HTTPLocationConfig struct {
	Id              int64                 `yaml:"id" json:"id"`                           // ID
	IsOn            bool                  `yaml:"isOn" json:"isOn"`                       // 是否启用
	Pattern         string                `yaml:"pattern" json:"pattern"`                 // 匹配规则 TODO 未来支持更多样的匹配规则
	Name            string                `yaml:"name" json:"name"`                       // 名称
	Web             *HTTPWebConfig        `yaml:"web" json:"web"`                         // Web配置
	URLPrefix       string                `yaml:"urlPrefix" json:"urlPrefix"`             // 实际的URL前缀，TODO 未来支持变量
	Description     string                `yaml:"description" json:"description"`         // 描述
	ReverseProxyRef *ReverseProxyRef      `yaml:"reverseProxyRef" json:"reverseProxyRef"` // 反向代理引用
	ReverseProxy    *ReverseProxyConfig   `yaml:"reverseProxy" json:"reverseProxy"`       // 反向代理设置
	IsBreak         bool                  `yaml:"isBreak" json:"isBreak"`                 // 终止向下解析
	Children        []*HTTPLocationConfig `yaml:"children" json:"children"`               // 子规则
}

func (this *HTTPLocationConfig) Init() error {
	if this.Web != nil {
		err := this.Web.Init()
		if err != nil {
			return err
		}
	}

	if this.ReverseProxy != nil {
		err := this.Web.Init()
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
