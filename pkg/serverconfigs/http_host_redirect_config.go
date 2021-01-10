package serverconfigs

import "net/url"

// 主机名跳转设置
type HTTPHostRedirectConfig struct {
	IsOn   bool `yaml:"isOn" json:"isOn"`     // 是否开启
	Status int  `yaml:"status" json:"status"` // 跳转用的状态码

	BeforeURL string `yaml:"beforeURL" json:"beforeURL"` // 跳转前的地址
	AfterURL  string `yaml:"afterURL" json:"afterURL"`   // 跳转后的地址

	realBeforeURL string
}

func (this *HTTPHostRedirectConfig) Init() error {
	{
		u, err := url.Parse(this.BeforeURL)
		if err != nil {
			return err
		}
		if len(u.Path) == 0 {
			this.realBeforeURL = this.BeforeURL + "/"
		} else {
			this.realBeforeURL = this.BeforeURL
		}
	}

	return nil
}

func (this *HTTPHostRedirectConfig) RealBeforeURL() string {
	return this.realBeforeURL
}
