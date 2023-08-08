package filterconfigs

import (
	"github.com/iwind/TeaGo/types"
	"strconv"
	"strings"
)

type UnicodeEncodeFilter struct {
}

// Init 初始化
func (this *UnicodeEncodeFilter) Init() error {
	return nil
}

// Do 执行过滤
func (this *UnicodeEncodeFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	var s = types.String(input)
	var result = strings.Builder{}
	for _, r := range s {
		if r < 128 {
			result.WriteRune(r)
		} else {
			result.WriteString("\\u" + strconv.FormatInt(int64(r), 16))
		}
	}
	return result.String(), true, nil
}
