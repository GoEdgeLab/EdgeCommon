package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHTTPAccessLogConfig_Match(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		accessLog := NewHTTPAccessLogConfig()
		a.IsNil(accessLog.Init())
		a.IsTrue(accessLog.Match(100))
		a.IsTrue(accessLog.Match(200))
		a.IsTrue(accessLog.Match(300))
		a.IsTrue(accessLog.Match(400))
		a.IsTrue(accessLog.Match(500))
	}

	{
		accessLog := NewHTTPAccessLogConfig()
		accessLog.Status1 = false
		accessLog.Status2 = false
		a.IsNil(accessLog.Init())
		a.IsFalse(accessLog.Match(100))
		a.IsFalse(accessLog.Match(200))
		a.IsTrue(accessLog.Match(300))
		a.IsTrue(accessLog.Match(400))
		a.IsTrue(accessLog.Match(500))
	}

	{
		accessLog := NewHTTPAccessLogConfig()
		accessLog.Status3 = false
		accessLog.Status4 = false
		accessLog.Status5 = false
		a.IsNil(accessLog.Init())
		a.IsTrue(accessLog.Match(100))
		a.IsTrue(accessLog.Match(200))
		a.IsFalse(accessLog.Match(300))
		a.IsFalse(accessLog.Match(400))
		a.IsFalse(accessLog.Match(500))
	}
}
