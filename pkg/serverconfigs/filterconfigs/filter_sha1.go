package filterconfigs

import (
	"crypto/sha1"
	"fmt"
	"github.com/iwind/TeaGo/types"
)

type Sha1Filter struct {
}

// 初始化
func (this *Sha1Filter) Init() error {
	return nil
}

// 执行过滤
func (this *Sha1Filter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	return fmt.Sprintf("%x", sha1.Sum([]byte(types.String(input)))), true, nil
}
