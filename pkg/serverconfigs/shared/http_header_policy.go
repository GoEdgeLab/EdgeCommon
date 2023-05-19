package shared

import "strings"

// HTTPHeaderPolicy HeaderList定义
type HTTPHeaderPolicy struct {
	Id          int64  `yaml:"id" json:"id"`                   // ID
	Name        string `yaml:"name" json:"name"`               // 名称 TODO
	IsOn        bool   `yaml:"isOn" json:"isOn"`               // 是否启用 TODO
	Description string `yaml:"description" json:"description"` // 描述 TODO

	SetHeaderRefs []*HTTPHeaderRef    `yaml:"setHeaderRefs" json:"setHeaderRefs"`
	SetHeaders    []*HTTPHeaderConfig `yaml:"setHeaders" json:"setHeaders"`
	DeleteHeaders []string            `yaml:"deleteHeaders" json:"deleteHeaders"` // 删除的Header

	Expires            *HTTPExpireHeaderConfig `yaml:"expires" json:"expires"`                       // 内容过期设置 TODO
	CORS               *HTTPCORSHeaderConfig   `yaml:"cors" json:"cors"`                             // CORS跨域设置
	NonStandardHeaders []string                `yaml:"nonStandardHeaders" json:"nonStandardHeaders"` // 非标Header列表

	setHeaderNames  []string
	deleteHeaderMap map[string]bool // header => bool
}

// Init 校验
func (this *HTTPHeaderPolicy) Init() error {
	this.setHeaderNames = []string{}
	for _, h := range this.SetHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
		this.setHeaderNames = append(this.setHeaderNames, strings.ToUpper(h.Name))
	}

	// delete
	this.deleteHeaderMap = map[string]bool{}
	for _, header := range this.DeleteHeaders {
		this.deleteHeaderMap[strings.ToUpper(header)] = true
	}

	// cors
	if this.CORS != nil {
		err := this.CORS.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// IsEmpty 判断是否为空
func (this *HTTPHeaderPolicy) IsEmpty() bool {
	return len(this.SetHeaders) == 0 && this.Expires == nil && len(this.DeleteHeaders) == 0
}

// ContainsHeader 判断Add和Set中是否包含某个Header
func (this *HTTPHeaderPolicy) ContainsHeader(name string) bool {
	name = strings.ToUpper(name)

	for _, n := range this.setHeaderNames {
		if n == name {
			return true
		}
	}
	return false
}

// ContainsDeletedHeader 判断删除列表中是否包含某个Header
func (this *HTTPHeaderPolicy) ContainsDeletedHeader(name string) bool {
	_, ok := this.deleteHeaderMap[strings.ToUpper(name)]
	return ok
}
