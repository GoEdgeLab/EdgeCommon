package filterconfigs

import (
	"github.com/iwind/TeaGo/types"
	"math/big"
)

type Hex2DecFilter struct {
}

// 初始化
func (this *Hex2DecFilter) Init() error {
	return nil
}

// 执行过滤
func (this *Hex2DecFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	n := new(big.Int)
	n.SetString(types.String(input), 16)
	return n.Uint64(), true, nil
}
