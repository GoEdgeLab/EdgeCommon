// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeutils

import (
	"github.com/iwind/TeaGo/maps"
	"testing"
)

func TestBase64EncodeMap(t *testing.T) {
	{
		var m maps.Map
		encodedString, err := Base64EncodeMap(m)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("encoded string:", encodedString)

		m, err = Base64DecodeMap(encodedString)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", m)
	}

	{
		var m = maps.Map{}
		encodedString, err := Base64EncodeMap(m)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("encoded string:", encodedString)

		m, err = Base64DecodeMap(encodedString)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", m)
	}

	{
		var m = maps.Map{"userId": 1, "name": "李白"}
		encodedString, err := Base64EncodeMap(m)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("encoded string:", encodedString)

		m, err = Base64DecodeMap(encodedString)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", m)
	}
}
