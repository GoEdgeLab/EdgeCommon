package configutils

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestMatchDomain(t *testing.T) {
	var a = assert.NewAssertion(t)
	{
		var ok = MatchDomains([]string{}, "example.com")
		a.IsFalse(ok)
	}

	{
		var ok = MatchDomains([]string{"example.com"}, "example.com")
		a.IsTrue(ok)
	}

	{
		var ok = MatchDomains([]string{"www.example.com"}, "example.com")
		a.IsFalse(ok)
	}

	{
		var ok = MatchDomains([]string{".example.com"}, "www.example.com")
		a.IsTrue(ok)
	}

	{
		var ok = MatchDomains([]string{".example.com"}, "a.www.example.com")
		a.IsTrue(ok)
	}

	{
		var ok = MatchDomains([]string{".example.com"}, "a.www.example123.com")
		a.IsFalse(ok)
	}

	{
		var ok = MatchDomains([]string{"*.example.com"}, "www.example.com")
		a.IsTrue(ok)
	}

	{
		var ok = MatchDomains([]string{"*.*.com"}, "www.example.com")
		a.IsTrue(ok)
	}

	{
		var ok = MatchDomains([]string{"www.*.com"}, "www.example.com")
		a.IsTrue(ok)
	}

	{
		var ok = MatchDomains([]string{"gallery.*.com"}, "www.example.com")
		a.IsFalse(ok)
	}

	{
		var ok = MatchDomains([]string{"~\\w+.example.com"}, "www.example.com")
		a.IsTrue(ok)
	}

	{
		var ok = MatchDomains([]string{"~\\w+.example.com"}, "a.www.example.com")
		a.IsTrue(ok)
	}

	{
		var ok = MatchDomains([]string{"~^\\d+.example.com$"}, "www.example.com")
		a.IsFalse(ok)
	}

	{
		var ok = MatchDomains([]string{"~^\\d+.example.com$"}, "123.example.com")
		a.IsTrue(ok)
	}
	{
		var ok = MatchDomains([]string{"*"}, "example.com")
		a.IsTrue(ok)
	}

	// port
	{
		var ok = MatchDomains([]string{"example.com:8001"}, "example.com:8001")
		a.IsTrue(ok)
	}
	{
		var ok = MatchDomains([]string{"example.com:8002"}, "example.com:8001")
		a.IsFalse(ok)
	}
	{
		var ok = MatchDomains([]string{"*.example.com:8001"}, "a.example.com:8001")
		a.IsTrue(ok)
	}
	{
		var ok = MatchDomains([]string{"a.example.com:*"}, "a.example.com:8001")
		a.IsTrue(ok)
	}
	{
		var ok = MatchDomains([]string{"a.example.com:*"}, "a.example.com")
		a.IsTrue(ok)
	}
	{
		var ok = MatchDomains([]string{"*.example.com:*"}, "a.example.com:8001")
		a.IsTrue(ok)
	}
	{
		var ok = MatchDomains([]string{"*.example.com:8002"}, "a.example.com:8001")
		a.IsFalse(ok)
	}
}

func TestIsSpecialDomain(t *testing.T) {
	var a = assert.NewAssertion(t)

	a.IsTrue(IsFuzzyDomain(""))
	a.IsTrue(IsFuzzyDomain(".hello.com"))
	a.IsTrue(IsFuzzyDomain("*.hello.com"))
	a.IsTrue(IsFuzzyDomain("hello.*.com"))
	a.IsTrue(IsFuzzyDomain("~^hello\\.com"))
	a.IsFalse(IsFuzzyDomain("hello.com"))
}
