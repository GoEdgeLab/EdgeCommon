package shared

import (
	"testing"
)

func TestHeaderList_FormatHeaders(t *testing.T) {
	list := NewHTTPHeaderPolicy()
	err := list.Init()
	if err != nil {
		t.Fatal(err)
	}
}
