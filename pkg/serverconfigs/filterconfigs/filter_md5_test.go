package filterconfigs

import "testing"

func TestMd5Filter_Do(t *testing.T) {
	filter := &Md5Filter{}
	t.Log(filter.Do("123456", nil))
	t.Log(filter.Do(nil, nil))
	t.Log(filter.Do("", nil))
	t.Log(filter.Do("hello", nil))
}
