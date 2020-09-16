package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestPageConfig_Match(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		page := NewHTTPPageConfig()
		page.Status = []string{"200"}
		err := page.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(page.Match(200))
		a.IsFalse(page.Match(201))
	}

	{
		page := NewHTTPPageConfig()
		page.Status = []string{"4xx", "5xx"}
		err := page.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(page.Match(200))
		a.IsTrue(page.Match(401))
		a.IsTrue(page.Match(404))
		a.IsTrue(page.Match(500))
		a.IsTrue(page.Match(505))
	}
}
