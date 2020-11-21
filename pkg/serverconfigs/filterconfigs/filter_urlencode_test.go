package filterconfigs

import "testing"

func TestURLEncodeFilter_Do(t *testing.T) {
	filter := &URLEncodeFilter{}
	t.Log(filter.Do("hello", nil))
	t.Log(filter.Do("/hello/world?a=b&c=中文&d=<symbol>", nil))
}
