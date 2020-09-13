package shared

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHeaderConfig_Match(t *testing.T) {
	a := assert.NewAssertion(t)
	h := NewHeaderConfig()
	err := h.Init()
	if err != nil {
		t.Fatal(err)
	}
	a.IsFalse(h.Match(200))
	a.IsFalse(h.Match(400))

	h.Status = &HTTPStatusConfig{
		Always: false,
		Codes:  []int{200, 301, 302, 400},
	}
	err = h.Init()
	if err != nil {
		t.Fatal(err)
	}
	a.IsTrue(h.Match(400))
	a.IsFalse(h.Match(500))

	h.Status.Always = true
	a.IsTrue(h.Match(500))
}
