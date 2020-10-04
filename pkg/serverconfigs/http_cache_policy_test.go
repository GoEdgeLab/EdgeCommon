package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHTTPCachePolicy_IsSame(t *testing.T) {
	a := assert.NewAssertion(t)
	{
		p1 := &HTTPCachePolicy{}
		p2 := &HTTPCachePolicy{}
		a.IsTrue(p1.IsSame(p2))
	}
	{
		p1 := &HTTPCachePolicy{
			Capacity: &shared.SizeCapacity{
				Count: 0,
				Unit:  "",
			},
		}
		p2 := &HTTPCachePolicy{}
		a.IsFalse(p1.IsSame(p2))
	}
	{
		p1 := &HTTPCachePolicy{
			Capacity: &shared.SizeCapacity{
				Count: 0,
				Unit:  "",
			},
		}
		p2 := &HTTPCachePolicy{
			Capacity: &shared.SizeCapacity{
				Count: 0,
				Unit:  "",
			},
		}
		a.IsTrue(p1.IsSame(p2))
	}
	{
		p1 := &HTTPCachePolicy{
			Options: map[string]interface{}{},
		}
		p2 := &HTTPCachePolicy{}
		a.IsFalse(p1.IsSame(p2))
	}
	{
		p1 := &HTTPCachePolicy{
			Options: map[string]interface{}{},
		}
		p2 := &HTTPCachePolicy{
			Options: map[string]interface{}{},
		}
		a.IsTrue(p1.IsSame(p2))
	}
	{
		p1 := &HTTPCachePolicy{
			Options: map[string]interface{}{
				"c": 3,
				"a": 1,
				"d": "abc",
				"b": 2,
			},
		}
		p2 := &HTTPCachePolicy{
			Options: map[string]interface{}{
				"c": 3,
				"a": 1,
				"d": "abc",
				"b": 2,
			},
		}
		a.IsTrue(p1.IsSame(p2))
	}
}
