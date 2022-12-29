package shared

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHTTPHeaderPolicy_FormatHeaders(t *testing.T) {
	policy := &HTTPHeaderPolicy{}
	err := policy.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestHTTPHeaderPolicy_ShouldDeleteHeader(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		policy := &HTTPHeaderPolicy{}
		err := policy.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(policy.ContainsDeletedHeader("Origin"))
	}
	{
		policy := &HTTPHeaderPolicy{
			DeleteHeaders: []string{"Hello", "World"},
		}
		err := policy.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(policy.ContainsDeletedHeader("Origin"))
	}
	{
		policy := &HTTPHeaderPolicy{
			DeleteHeaders: []string{"origin"},
		}
		err := policy.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(policy.ContainsDeletedHeader("Origin"))
	}
	{
		policy := &HTTPHeaderPolicy{
			DeleteHeaders: []string{"Origin"},
		}
		err := policy.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(policy.ContainsDeletedHeader("Origin"))
	}
}
