package filterconfigs

import "testing"

func TestLengthFilter_Do(t *testing.T) {
	filter := &LengthFilter{}
	t.Log(filter.Do("hello", nil))
	t.Log(filter.Do([]byte("hello"), nil))
}
