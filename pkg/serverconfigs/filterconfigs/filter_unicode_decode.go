package filterconfigs

import (
	"github.com/iwind/TeaGo/types"
	"strconv"
	"strings"
)

type UnicodeDecodeFilter struct {
}

// 初始化
func (this *UnicodeDecodeFilter) Init() error {
	return nil
}

// 执行过滤
func (this *UnicodeDecodeFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	s := types.String(input)
	result, err := strconv.Unquote("\"" + strings.ReplaceAll(s, "\"", "\\\"") + "\"")
	if err != nil {
		return input, true, nil
	}
	return result, true, nil
}
