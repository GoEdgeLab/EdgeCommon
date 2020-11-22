package filterconfigs

import (
	stringutil "github.com/iwind/TeaGo/utils/string"
	"testing"
)

func TestHex2DecFilter_Do(t *testing.T) {
	filter := &Hex2DecFilter{}
	err := filter.Init()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(filter.Do("0e", nil))
	t.Log(filter.Do("e", nil))

	{
		result, _, _ := filter.Do("123", nil)
		t.Logf("%x", result)
	}

	{
		md5 := stringutil.Md5("123456")
		t.Log("md5:", md5)
		result, _, _ := filter.Do(md5, nil)
		t.Log(result)
		t.Logf("%x", result)
	}
}
