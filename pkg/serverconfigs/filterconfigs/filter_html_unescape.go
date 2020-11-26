package filterconfigs

import (
	"github.com/iwind/TeaGo/types"
	"html"
)

type HTMLUnescapeFilter struct {
}

// 初始化
func (this *HTMLUnescapeFilter) Init() error {
	return nil
}

// 执行过滤
func (this *HTMLUnescapeFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	s := types.String(input)
	return html.UnescapeString(s), true, nil
}
