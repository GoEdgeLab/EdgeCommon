// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeutils

import (
	"github.com/iwind/TeaGo/maps"
	"testing"
)

func TestEncryptData(t *testing.T) {
	e, err := EncryptData("a", "b", maps.Map{
		"c": 1,
	}, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("e:", e)

	s, err := DecryptData("a", "b", e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("s:", s)
}

func BenchmarkEncryptData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = EncryptData("a", "b", maps.Map{
			"c": 1,
		}, 5)
	}
}
