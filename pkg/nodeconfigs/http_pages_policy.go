// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodeconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"

// HTTPPagesPolicy 全局的HTTP自定义页面设置
type HTTPPagesPolicy struct {
	IsOn  bool                            `json:"isOn" yaml:"isOn"`   // 是否启用
	Pages []*serverconfigs.HTTPPageConfig `json:"pages" yaml:"pages"` // 自定义页面
}

func NewHTTPPagesPolicy() *HTTPPagesPolicy {
	return &HTTPPagesPolicy{}
}

func (this *HTTPPagesPolicy) Init() error {
	if len(this.Pages) > 0 {
		for _, page := range this.Pages {
			err := page.Init()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
