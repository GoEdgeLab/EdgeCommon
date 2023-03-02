// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeutils

import (
	"github.com/iwind/TeaGo/maps"
	"testing"
)

func TestEncryptMap(t *testing.T) {
	e, err := EncryptMap("a", "b", maps.Map{
		"c": 1,
	}, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("e:", e)

	s, err := DecryptMap("a", "b", e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("s:", s)
}

func TestEncryptData(t *testing.T) {
	encoded, err := EncryptData("a", "b", []byte("Hello, World"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("encoded:", encoded)

	source, err := DecryptData("a", "b", encoded)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("source:", string(source))
}

func BenchmarkEncryptData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = EncryptMap("a", "b", maps.Map{
			"c": 1,
		}, 5)
	}
}
