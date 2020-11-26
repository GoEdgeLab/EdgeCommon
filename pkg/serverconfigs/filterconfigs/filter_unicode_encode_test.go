package filterconfigs

import "testing"

func TestUnicodeEncodeFilter_Do(t *testing.T) {
	filter := &UnicodeEncodeFilter{}
	t.Log(filter.Do("Hello", nil))
	t.Log(filter.Do("我是中文", nil))
	t.Log(filter.Do("我是中文和英文Mixed", nil))
	t.Log(filter.Do("我有特殊字符|'\"", nil))
}
