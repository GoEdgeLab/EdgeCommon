package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
	"time"
)

func TestHTTPRewriteRule(t *testing.T) {
	a := assert.NewAssertion(t).Quiet()

	rule := HTTPRewriteRule{
		Pattern: "/(hello)/(world)",
		Replace: "/${1}/${2}",
	}
	a.IsNil(rule.Init())

	{
		replace, _, b := rule.MatchRequest("/hello/worl", func(source string) string {
			return source
		})
		a.IsFalse(b)
		a.Log("url:", replace)
	}

	{
		replace, _, b := rule.MatchRequest("/hello/world", func(source string) string {
			return source
		})
		a.IsTrue(b)
		a.Log("url:", replace)
	}

	{
		r := HTTPRewriteRule{}
		r.Replace = "http://127.0.0.1${0}"
		r.Pattern = ".*"
		err := r.Init()
		if err != nil {
			t.Fatal(err)
		}
		u, _, b := r.MatchRequest("/hello", func(source string) string {
			return source
		})
		a.Log(b)
		a.Log(u)
	}
}

func TestRewriteRule_NamedMatch(t *testing.T) {
	r := &HTTPRewriteRule{}
	r.Replace = "http://127.0.0.1/${1}/${last}/${ni}"
	r.Pattern = "/(\\w+)/(?P<last>\\w+)/(?P<ni>\\w+)"
	err := r.Init()
	if err != nil {
		t.Fatal(err)
	}

	before := time.Now()
	count := 100
	for i := 0; i < count; i++ {
		s, _, b := r.MatchRequest("/hello/world/ni", func(source string) string {
			return source
		})
		if i == 0 {
			if b {
				t.Log("matched:", s)
			} else {
				t.Log("not matched")
			}
		}
	}
	t.Log(float64(count) / (time.Since(before).Seconds()))
}

func TestRewriteRule_CaseInsensitive(t *testing.T) {
	a := assert.NewAssertion(t)

	r := &HTTPRewriteRule{}
	r.Replace = "http://127.0.0.1${0}"
	r.Pattern = "(?i)/index.php"
	err := r.Init()
	if err != nil {
		t.Fatal(err)
	}
	_, _, ok := r.MatchRequest("/index.php", func(source string) string {
		return source
	})
	a.IsTrue(ok)

	_, _, ok = r.MatchRequest("/INDEX.php", func(source string) string {
		return source
	})
	a.IsTrue(ok)
}

func TestRewriteRule_Slashes(t *testing.T) {
	a := assert.NewAssertion(t)

	r := &HTTPRewriteRule{}
	r.Replace = "http://127.0.0.1/${0}"
	r.Pattern = "(?i)/index.php"
	err := r.Init()
	if err != nil {
		t.Fatal(err)
	}
	replace, _, ok := r.MatchRequest("/index.php", func(source string) string {
		return source
	})
	a.IsTrue(ok)
	t.Log(replace)
}

func TestRewriteRuleProxy(t *testing.T) {
	a := assert.NewAssertion(t).Quiet()

	rule := &HTTPRewriteRule{
		Pattern: "/(hello)/(world)",
		Replace: "/${1}/${2}",
	}
	a.IsNil(rule.Init())

	replace, _, b := rule.MatchRequest("/hello/world", func(source string) string {
		return source
	})
	a.IsTrue(b)
	a.IsTrue(replace == "/hello/world")
}
