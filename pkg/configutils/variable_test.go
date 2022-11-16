package configutils

import (
	"github.com/iwind/TeaGo/types"
	"runtime"
	"strconv"
	"testing"
)

func TestParseVariables(t *testing.T) {
	{
		v := ParseVariables("hello, ${name}, world", func(s string) string {
			return "Lu"
		})
		t.Log(v)
	}
	{
		v := ParseVariables("hello, world", func(s string) string {
			return "Lu"
		})
		t.Log(v)
	}
	{
		v := ParseVariables("${name}", func(s string) string {
			return "Lu"
		})
		t.Log(v)
	}
}

func TestParseNoVariables(t *testing.T) {
	for i := 0; i < 2; i++ {
		v := ParseVariables("hello, world", func(s string) string {
			return "Lu"
		})
		t.Log(v)
	}
}

func TestParseHolders(t *testing.T) {
	var holders = ParseHolders("hello, ${name}, world")
	for _, h := range holders {
		t.Log(types.String(h))
	}

	t.Log("parse result:", ParseVariablesFromHolders(holders, func(s string) string {
		return "[" + s + "]"
	}))
}

func BenchmarkParseVariables(b *testing.B) {
	_ = ParseVariables("hello, ${name}, ${age}, ${gender}, ${home}, world", func(s string) string {
		return "Lu"
	})

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = ParseVariables("hello, ${name}, ${age}, ${gender}, ${home}, world", func(s string) string {
				return "Lu"
			})
		}
	})
}

func BenchmarkParseVariablesFromHolders(b *testing.B) {
	var holders = ParseHolders("hello, ${name}, ${age}, ${gender}, ${home}, world")

	for i := 0; i < b.N; i++ {
		_ = ParseVariablesFromHolders(holders, func(s string) string {
			return "Lu"
		})
	}
}

func BenchmarkParseVariablesUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ParseVariables("hello, ${name} "+strconv.Itoa(i), func(s string) string {
			return "Lu"
		})
	}
}

func BenchmarkParseVariablesUnique_Single(b *testing.B) {
	runtime.GOMAXPROCS(1)

	for i := 0; i < b.N; i++ {
		_ = ParseVariables("${name}", func(s string) string {
			return "Lu"
		})
	}
}

func BenchmarkParseNoVariables(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ParseVariables("hello, world", func(s string) string {
			return "Lu"
		})
	}
}

func BenchmarkParseEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ParseVariables("", func(s string) string {
			return "Lu"
		})
	}
}
