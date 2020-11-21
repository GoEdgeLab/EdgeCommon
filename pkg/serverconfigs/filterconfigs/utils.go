package filterconfigs

import "github.com/iwind/TeaGo/types"

// 将输入内容转换为字节
func ToBytes(input interface{}) []byte {
	if input == nil {
		return []byte{}
	}
	var data []byte
	var ok bool
	data, ok = input.([]byte)
	if ok {
		return data
	}
	return []byte(types.String(input))
}

// 将输入内容转换为字符串
func ToString(input interface{}) string {
	if input == nil {
		return ""
	}
	return types.String(input)
}
