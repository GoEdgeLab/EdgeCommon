package shared

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHTTPRequestCondGroup_MatchRequest(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		group := &HTTPRequestCondGroup{}
		group.Connector = "or"
		group.IsOn = false
		err := group.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(group.MatchRequest(func(source string) string {
			return source
		}))
		a.IsTrue(group.MatchResponse(func(source string) string {
			return source
		}))
	}

	{
		group := &HTTPRequestCondGroup{}
		group.IsOn = true
		err := group.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(group.MatchRequest(func(source string) string {
			return source
		}))
		a.IsTrue(group.MatchResponse(func(source string) string {
			return source
		}))
	}

	{
		group := &HTTPRequestCondGroup{}
		group.IsOn = true
		group.Connector = "or"
		group.Conds = []*HTTPRequestCond{
			{
				IsRequest: true,
				Param:     "456",
				Operator:  "gt",
				Value:     "123",
			},
			{
				IsRequest: false,
				Param:     "123",
				Operator:  "gt",
				Value:     "456",
			},
		}
		err := group.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(group.MatchRequest(func(source string) string {
			return source
		}))
		a.IsFalse(group.MatchResponse(func(source string) string {
			return source
		}))
	}
	{
		group := &HTTPRequestCondGroup{}
		group.IsOn = true
		group.Connector = "or"
		group.Conds = []*HTTPRequestCond{
			{
				IsRequest: true,
				Param:     "456",
				Operator:  "gt",
				Value:     "1234",
			},
			{
				IsRequest: true,
				Param:     "456",
				Operator:  "gt",
				Value:     "123",
			},
			{
				IsRequest: false,
				Param:     "123",
				Operator:  "gt",
				Value:     "456",
			},
		}
		err := group.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(group.MatchRequest(func(source string) string {
			return source
		}))
		a.IsFalse(group.MatchResponse(func(source string) string {
			return source
		}))
	}
	{
		group := &HTTPRequestCondGroup{}
		group.IsOn = true
		group.Connector = "and"
		group.Conds = []*HTTPRequestCond{
			{
				IsRequest: true,
				Param:     "456",
				Operator:  "gt",
				Value:     "123",
			},
			{
				IsRequest: true,
				Param:     "456",
				Operator:  "gt",
				Value:     "1234",
			},
			{
				IsRequest: false,
				Param:     "123",
				Operator:  "gt",
				Value:     "456",
			},
		}
		err := group.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(group.MatchRequest(func(source string) string {
			return source
		}))
		a.IsFalse(group.MatchResponse(func(source string) string {
			return source
		}))
	}
}
