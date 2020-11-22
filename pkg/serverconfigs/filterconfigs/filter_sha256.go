package filterconfigs

import (
	"crypto/sha256"
	"fmt"
	"github.com/iwind/TeaGo/types"
)

type Sha256Filter struct {
}

// 初始化
func (this *Sha256Filter) Init() error {
	return nil
}

// 执行过滤
func (this *Sha256Filter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(types.String(input)))), true, nil
}
