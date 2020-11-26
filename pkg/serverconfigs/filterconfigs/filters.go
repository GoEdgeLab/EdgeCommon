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
	"md5":           new(Md5Filter),
	"urlEncode":     new(URLEncodeFilter),
	"urlDecode":     new(URLDecodeFilter),
	"base64Encode":  new(Base64EncodeFilter),
	"base64Decode":  new(Base64DecodeFilter),
	"unicodeEncode": new(UnicodeEncodeFilter),
	"unicodeDecode": new(UnicodeDecodeFilter),
	"htmlEscape":    new(HTMLEscapeFilter),
	"htmlUnescape":  new(HTMLUnescapeFilter),
	"length":        new(LengthFilter),
	"hex2dec":       new(Hex2DecFilter),
	"dec2hex":       new(Dec2HexFilter),
	"sha1":          new(Sha1Filter),
	"sha256":        new(Sha256Filter),
}

// 查找Filter
func FindFilter(code string) FilterInterface {
	return allFilters[code]
}
