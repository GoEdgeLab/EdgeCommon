package shared

// HeaderList定义
type HTTPHeaderPolicy struct {
	Id   int64 `yaml:"id" json:"id"`
	IsOn bool  `yaml:"isOn" json:"isOn"` // TODO

	AddHeaders     []*HTTPHeaderConfig `yaml:"addHeaders" json:"addHeaders"`         // TODO
	AddTrailers    []*HTTPHeaderConfig `yaml:"addTrailers" json:"addTrailers"`       // TODO
	SetHeaders     []*HTTPHeaderConfig `yaml:"setHeaders" json:"setHeaders"`         // TODO
	ReplaceHeaders []*HTTPHeaderConfig `yaml:"replaceHeaders" json:"replaceHeaders"` // 替换Header内容 TODO
	DeletedHeaders []string            `yaml:"deleteHeaders" json:"deleteHeaders"`   // 删除的Header TODO

	Expires *HTTPExpireHeaderConfig `yaml:"expires" json:"expires"` // TODO
}

// 获取新对象
func NewHTTPHeaderPolicy() *HTTPHeaderPolicy {
	return &HTTPHeaderPolicy{}
}

// 校验
func (this *HTTPHeaderPolicy) Init() error {
	for _, h := range this.AddHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	for _, h := range this.AddTrailers {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	for _, h := range this.SetHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	for _, h := range this.ReplaceHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 判断是否为空
func (this *HTTPHeaderPolicy) IsEmpty() bool {
	return len(this.AddHeaders) == 0 && len(this.AddTrailers) == 0 && len(this.SetHeaders) == 0 && len(this.ReplaceHeaders) == 0 && this.Expires == nil
}
