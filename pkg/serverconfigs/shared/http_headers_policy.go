package shared

import "strings"

// HTTPHeaderPolicy HeaderList定义
type HTTPHeaderPolicy struct {
	Id          int64  `yaml:"id" json:"id"`                   // ID
	Name        string `yaml:"name" json:"name"`               // 名称 TODO
	IsOn        bool   `yaml:"isOn" json:"isOn"`               // 是否启用 TODO
	Description string `yaml:"description" json:"description"` // 描述 TODO

	AddHeaderRefs     []*HTTPHeaderRef    `yaml:"addHeaderRefs" json:"addHeaderRefs"`
	AddHeaders        []*HTTPHeaderConfig `yaml:"addHeaders" json:"addHeaders"`
	AddTrailerRefs    []*HTTPHeaderRef    `yaml:"addTrailerRefs" json:"addTrailerRefs"`
	AddTrailers       []*HTTPHeaderConfig `yaml:"addTrailers" json:"addTrailers"` // TODO
	SetHeaderRefs     []*HTTPHeaderRef    `yaml:"setHeaderRefs" json:"setHeaderRefs"`
	SetHeaders        []*HTTPHeaderConfig `yaml:"setHeaders" json:"setHeaders"`
	ReplaceHeaderRefs []*HTTPHeaderRef    `yaml:"replaceHeaderRefs" json:"replaceHeaderRefs"`
	ReplaceHeaders    []*HTTPHeaderConfig `yaml:"replaceHeaders" json:"replaceHeaders"` // 替换Header内容 TODO
	DeleteHeaders     []string            `yaml:"deleteHeaders" json:"deleteHeaders"`   // 删除的Header

	Expires *HTTPExpireHeaderConfig `yaml:"expires" json:"expires"` // TODO

	addHeaderNames  []string
	setHeaderNames  []string
	deleteHeaderMap map[string]bool // header => bool
}

// Init 校验
func (this *HTTPHeaderPolicy) Init() error {
	this.addHeaderNames = []string{}
	for _, h := range this.AddHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
		this.addHeaderNames = append(this.addHeaderNames, strings.ToUpper(h.Name))
	}

	for _, h := range this.AddTrailers {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	this.setHeaderNames = []string{}
	for _, h := range this.SetHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
		this.setHeaderNames = append(this.setHeaderNames, strings.ToUpper(h.Name))
	}

	for _, h := range this.ReplaceHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	// delete
	this.deleteHeaderMap = map[string]bool{}
	for _, header := range this.DeleteHeaders {
		this.deleteHeaderMap[strings.ToUpper(header)] = true
	}

	return nil
}

// IsEmpty 判断是否为空
func (this *HTTPHeaderPolicy) IsEmpty() bool {
	return len(this.AddHeaders) == 0 && len(this.AddTrailers) == 0 && len(this.SetHeaders) == 0 && len(this.ReplaceHeaders) == 0 && this.Expires == nil && len(this.DeleteHeaders) == 0
}

// ContainsHeader 判断Add和Set中是否包含某个Header
func (this *HTTPHeaderPolicy) ContainsHeader(name string) bool {
	name = strings.ToUpper(name)

	for _, n := range this.addHeaderNames {
		if n == name {
			return true
		}
	}
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
