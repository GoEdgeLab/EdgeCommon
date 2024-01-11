package configutils_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/assert"
	"github.com/iwind/TeaGo/types"
	"runtime"
	"strconv"
	"testing"
)

func TestParseVariables(t *testing.T) {
	var a = assert.NewAssertion(t)

	{
		var v = configutils.ParseVariables("hello, ${name}, world", func(s string) string {
			return "Lu"
		})
		t.Log(v)
		a.IsTrue(v == "hello, Lu, world")
	}
	{
		var v = configutils.ParseVariables("hello, world", func(s string) string {
			return "Lu"
		})
		t.Log(v)
		a.IsTrue(v == "hello, world")
	}
	{
		var v = configutils.ParseVariables("${name}", func(s string) string {
			return "Lu"
		})
		t.Log(v)
		a.IsTrue(v == "Lu")
	}
}

func TestParseNoVariables(t *testing.T) {
	for i := 0; i < 2; i++ {
		var v = configutils.ParseVariables("hello, world", func(s string) string {
			return "Lu"
		})
		t.Log(v)
	}
}

func TestParseVariables_Modifier(t *testing.T) {
	t.Log(configutils.ParseVariables("${url|urlEncode}", func(varName string) (value string) {
		switch varName {
		case "url":
			return "/hello/world?a=1"
		}
		return "${" + varName + "}"
	}))
	t.Log(configutils.ParseVariables("${url|urlDecode}", func(varName string) (value string) {
		switch varName {
		case "url":
			return "%2Fhello%2Fworld%3Fa%3D1"
		}
		return "${" + varName + "}"
	}))
	t.Log(configutils.ParseVariables("${url|urlDecode|urlEncode}", func(varName string) (value string) {
		switch varName {
		case "url":
			return "%2Fhello%2Fworld%3Fa%3D1"
		}
		return "${" + varName + "}"
	}))
	t.Log(configutils.ParseVariables("${var|base64Encode}", func(varName string) (value string) {
		switch varName {
		case "var":
			return "123456"
		}
		return "${" + varName + "}"
	}))
	t.Log(configutils.ParseVariables("${var|base64Encode|base64Decode}", func(varName string) (value string) {
		switch varName {
		case "var":
			return "123456"
		}
		return "${" + varName + "}"
	}))
	t.Log(configutils.ParseVariables("${var|md5}", func(varName string) (value string) {
		switch varName {
		case "var":
			return "123456"
		}
		return "${" + varName + "}"
	}))
	t.Log(configutils.ParseVariables("${var|sha1}", func(varName string) (value string) {
		switch varName {
		case "var":
			return "123456"
		}
		return "${" + varName + "}"
	}))
	t.Log(configutils.ParseVariables("${var|sha256}", func(varName string) (value string) {
		switch varName {
		case "var":
			return "123456"
		}
		return "${" + varName + "}"
	}))
	t.Log(configutils.ParseVariables("${var|toLowerCase}", func(varName string) (value string) {
		switch varName {
		case "var":
			return "ABC"
		}
		return "${" + varName + "}"
	}))
	t.Log(configutils.ParseVariables("${var|toUpperCase}", func(varName string) (value string) {
		switch varName {
		case "var":
			return "abc"
		}
		return "${" + varName + "}"
	}))
}

func TestParseHolders(t *testing.T) {
	var holders = configutils.ParseHolders("hello, ${name|urlencode}, world")
	t.Log("===holders begin===")
	for _, h := range holders {
		t.Log(types.String(h))
	}
	t.Log("===holders end===")

	t.Log("parse result:", configutils.ParseVariablesFromHolders(holders, func(s string) string {
		return "[" + s + "]"
	}))
}

func BenchmarkParseVariables(b *testing.B) {
	_ = configutils.ParseVariables("hello, ${name}, ${age}, ${gender}, ${home}, world", func(s string) string {
		return "Lu"
	})

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = configutils.ParseVariables("hello, ${name}, ${age}, ${gender}, ${home}, world", func(s string) string {
				return "Lu"
			})
		}
	})
}

func BenchmarkParseVariablesFromHolders(b *testing.B) {
	var holders = configutils.ParseHolders("hello, ${name}, ${age}, ${gender}, ${home}, world")

	for i := 0; i < b.N; i++ {
		_ = configutils.ParseVariablesFromHolders(holders, func(s string) string {
			return "Lu"
		})
	}
}

func BenchmarkParseVariablesUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = configutils.ParseVariables("hello, ${name} "+strconv.Itoa(i%100_000), func(s string) string {
			return "Lu"
		})
	}
}

func BenchmarkParseVariablesUnique_Single(b *testing.B) {
	runtime.GOMAXPROCS(1)

	for i := 0; i < b.N; i++ {
		_ = configutils.ParseVariables("${name}", func(s string) string {
			return "Lu"
		})
	}
}

func BenchmarkParseNoVariables(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = configutils.ParseVariables("hello, world", func(s string) string {
			return "Lu"
		})
	}
}

func BenchmarkParseEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = configutils.ParseVariables("", func(s string) string {
			return "Lu"
		})
	}
}
