package filterconfigs

import (
	"fmt"
	"github.com/iwind/TeaGo/types"
)

type Dec2HexFilter struct {
}

// 初始化
func (this *Dec2HexFilter) Init() error {
	return nil
}

// 执行过滤
func (this *Dec2HexFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	v := types.Int64(input)
	return fmt.Sprintf("%x", v), true, nil
}
