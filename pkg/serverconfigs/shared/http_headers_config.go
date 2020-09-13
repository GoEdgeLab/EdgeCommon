package shared

// HeaderList定义
type HTTPHeadersConfig struct {
	AddHeaders     []*HTTPHeaderConfig `yaml:"addHeaders" json:"addHeaders"`         // TODO
	AddTrailers    []*HTTPHeaderConfig `yaml:"addTrailers" json:"addTrailers"`       // TODO
	SetHeaders     []*HTTPHeaderConfig `yaml:"setHeaders" json:"setHeaders"`         // TODO
	ReplaceHeaders []*HTTPHeaderConfig `yaml:"replaceHeaders" json:"replaceHeaders"` // TODO

	Expires *HTTPExpireHeaderConfig `yaml:"expires" json:"expires"` // TODO
}

// 获取新对象
func NewHTTPHeaders() *HTTPHeadersConfig {
	return &HTTPHeadersConfig{}
}

// 校验
func (this *HTTPHeadersConfig) Init() error {
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
func (this *HTTPHeadersConfig) IsEmpty() bool {
	return len(this.AddHeaders) == 0 && len(this.AddTrailers) == 0 && len(this.SetHeaders) == 0 && len(this.ReplaceHeaders) == 0 && this.Expires == nil
}
