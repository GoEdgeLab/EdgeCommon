// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHTTPCompressionConfig_Init(t *testing.T) {
	{
		var config = &HTTPCompressionConfig{
			IsPrior:    false,
			IsOn:       false,
			Types:      nil,
			Level:      0,
			MinLength:  nil,
			MaxLength:  nil,
			MimeTypes:  nil,
			Extensions: nil,
			Conds:      nil,
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}

		a := assert.NewAssertion(t)
		a.IsTrue(config.MatchResponse("text/html", 1024, "", func(s string) string {
			return s
		}))
		a.IsTrue(config.MatchResponse("text/html", 1024, ".html", func(s string) string {
			return s
		}))
	}

	{
		var config = &HTTPCompressionConfig{
			IsPrior:    false,
			IsOn:       false,
			Types:      nil,
			Level:      0,
			MinLength:  nil,
			MaxLength:  nil,
			MimeTypes:  []string{"text/html", "text/plain"},
			Extensions: nil,
			Conds:      nil,
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}

		a := assert.NewAssertion(t)
		a.IsTrue(config.MatchResponse("text/html", 1024, "", func(s string) string {
			return s
		}))
		a.IsTrue(config.MatchResponse("text/html; charset=utf-8", 1024, ".html", func(s string) string {
			return s
		}))
		a.IsFalse(config.MatchResponse("image/png", 1024, ".html", func(s string) string {
			return s
		}))
	}

	{
		var config = &HTTPCompressionConfig{
			IsPrior:    false,
			IsOn:       false,
			Types:      nil,
			Level:      0,
			MinLength:  nil,
			MaxLength:  nil,
			MimeTypes:  []string{"text/html", "text/plain"},
			Extensions: []string{".html", ".txt"},
			Conds:      nil,
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}

		a := assert.NewAssertion(t)
		a.IsTrue(config.MatchResponse("text/html", 1024, "", func(s string) string {
			return s
		}))
		a.IsTrue(config.MatchResponse("text/plain", 1024, ".txt", func(s string) string {
			return s
		}))
		a.IsTrue(config.MatchResponse("text/html; charset=utf-8", 1024, ".html", func(s string) string {
			return s
		}))
		a.IsTrue(config.MatchResponse("image/png", 1024, ".html", func(s string) string {
			return s
		}))
	}
}

func TestHTTPCompressionConfig_MatchAcceptEncoding(t *testing.T) {
	var config = &HTTPCompressionConfig{
		GzipRef: &HTTPCompressionRef{
			//IsOn: true,
		},
		Gzip: &HTTPGzipCompressionConfig{
			IsOn: true,
		},

		DeflateRef: &HTTPCompressionRef{IsOn: true},
		Deflate:    &HTTPDeflateCompressionConfig{IsOn: false},

		BrotliRef: &HTTPCompressionRef{IsOn: true},
		Brotli:    &HTTPBrotliCompressionConfig{IsOn: true},
	}
	err := config.Init()
	if err != nil {
		t.Fatal(err)
	}

	for _, encodings := range []string{"gzip, deflate, br", "gzip, deflate", "deflate", "br", "compress", "br;q=0.8"} {
		result, encoding, ok := config.MatchAcceptEncoding(encodings)
		t.Log(encodings+" -> "+result+"/"+encoding, ok)
	}
}

func TestHTTPCompressionConfig_MatchAcceptEncoding2(t *testing.T) {
	var config = &HTTPCompressionConfig{
		//UseDefaultTypes: true,
		Types: []HTTPCompressionType{"brotli"},
	}
	err := config.Init()
	if err != nil {
		t.Fatal(err)
	}

	for _, encodings := range []string{"gzip, deflate, br", "gzip, deflate", "deflate", "br", "compress", "br;q=0.8"} {
		result, encoding, ok := config.MatchAcceptEncoding(encodings)
		t.Log(encodings+" -> "+result+"/"+encoding, ok)
	}
}
