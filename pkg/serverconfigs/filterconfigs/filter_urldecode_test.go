package filterconfigs

import (
	"net/url"
	"testing"
)

func TestURLDecodeFilter_Do(t *testing.T) {
	filter := &URLDecodeFilter{}
	t.Log(filter.Do("hello", nil))
	t.Log(filter.Do(url.QueryEscape("/hello/world/?a=b&c=d"), nil))
	t.Log(filter.Do("/hello/world/?a=b&c=d", nil))
}
