package filterconfigs

import (
	"github.com/iwind/TeaGo/types"
	"html"
)

type HTMLEscapeFilter struct {
}

// 初始化
func (this *HTMLEscapeFilter) Init() error {
	return nil
}

// 执行过滤
func (this *HTMLEscapeFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	s := types.String(input)
	return html.EscapeString(s), true, nil
}
