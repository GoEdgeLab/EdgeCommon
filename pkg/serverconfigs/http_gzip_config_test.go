package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestGzipConfig_MatchContentType(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		gzip := &HTTPGzipConfig{}
		a.IsNil(gzip.Init())
		a.IsTrue(gzip.MatchContentType("text/html"))
	}

	{
		gzip := &HTTPGzipConfig{}
		a.IsNil(gzip.Init())
		a.IsTrue(gzip.MatchContentType("text/html; charset=utf-8"))
	}

	{
		gzip := &HTTPGzipConfig{}
		gzip.MimeTypes = []string{"text/*"}
		a.IsNil(gzip.Init())
		a.IsTrue(gzip.MatchContentType("text/html; charset=utf-8"))
	}

	{
		gzip := &HTTPGzipConfig{}
		gzip.MimeTypes = []string{"text/*"}
		a.IsNil(gzip.Init())
		a.IsFalse(gzip.MatchContentType("application/json; charset=utf-8"))
	}

	{
		gzip := &HTTPGzipConfig{}
		gzip.MimeTypes = []string{"text/*", "image/*"}
		a.IsNil(gzip.Init())
		a.IsTrue(gzip.MatchContentType("image/jpeg; charset=utf-8"))
	}
}
