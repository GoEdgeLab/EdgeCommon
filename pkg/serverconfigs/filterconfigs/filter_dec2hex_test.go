package filterconfigs

import (
	"testing"
)

func TestDec2HexFilter_Do(t *testing.T) {
	filter := &Dec2HexFilter{}
	err := filter.Init()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(filter.Do("123456", nil))
	t.Log(filter.Do("1", nil))
}
