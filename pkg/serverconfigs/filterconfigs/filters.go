package filterconfigs

import "github.com/iwind/TeaGo/logs"

func init() {
	for code, filter := range allFilters {
		err := filter.Init()
		if err != nil {
			logs.Println("[FILTER]init '" + code + "' failed: " + err.Error())
		}
	}
}

// 所有的筛选条件
var allFilters = map[string]FilterInterface{
	"md5":          new(Md5Filter),
	"urlEncode":    new(URLEncodeFilter),
	"urlDecode":    new(URLDecodeFilter),
	"base64Encode": new(Base64EncodeFilter),
	"base64Decode": new(Base64DecodeFilter),
	"length":       new(LengthFilter),
}

// 查找Filter
func FindFilter(code string) FilterInterface {
	return allFilters[code]
}
