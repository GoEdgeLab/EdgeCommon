package filterconfigs

import "testing"

func TestUnicodeDecodeFilter_Do(t *testing.T) {
	filter := &UnicodeDecodeFilter{}
	t.Log(filter.Do(`"\u5947\u5c4f`, nil))
	t.Log(filter.Do(`"Hello`, nil))
	t.Log(filter.Do(`真实的存在`, nil))
}
