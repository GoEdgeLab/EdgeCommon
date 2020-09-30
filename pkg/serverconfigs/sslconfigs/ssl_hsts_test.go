package sslconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHSTSConfig(t *testing.T) {
	h := &HSTSConfig{}
	err := h.Init()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(h.HeaderValue())

	h.IncludeSubDomains = true
	err = h.Init()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(h.HeaderValue())

	h.Preload = true
	err = h.Init()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(h.HeaderValue())

	h.IncludeSubDomains = false
	err = h.Init()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(h.HeaderValue())

	h.MaxAge = 86400
	err = h.Init()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(h.HeaderValue())

	a := assert.NewAssertion(t)
	a.IsTrue(h.Match("abc.com"))

	h.Domains = []string{"abc.com"}
	err = h.Init()
	if err != nil {
		t.Fatal(err)
	}
	a.IsTrue(h.Match("abc.com"))

	h.Domains = []string{"1.abc.com"}
	err = h.Init()
	if err != nil {
		t.Fatal(err)
	}
	a.IsFalse(h.Match("abc.com"))
}
