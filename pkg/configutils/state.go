package configutils

import "github.com/iwind/TeaGo/types"

type BoolState = int8

const (
	BoolStateAll BoolState = 0 // 全部
	BoolStateYes BoolState = 1 // 已安装
	BoolStateNo  BoolState = 2 // 未安装
)

func ToBoolState(v interface{}) BoolState {
	return types.Int8(v)
}
