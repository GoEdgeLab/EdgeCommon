// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package shared

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestMimeTypeRule_Match(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		rule, err := NewMimeTypeRule("text/plain")
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(rule.Match("text/plain"))
		a.IsTrue(rule.Match("TEXT/plain"))
		a.IsFalse(rule.Match("text/html"))
	}

	{
		rule, err := NewMimeTypeRule("image/*")
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(rule.Match("image/png"))
		a.IsTrue(rule.Match("IMAGE/jpeg"))
		a.IsFalse(rule.Match("image/"))
		a.IsFalse(rule.Match("image1/png"))
		a.IsFalse(rule.Match("x-image/png"))
	}

	{
		_, err := NewMimeTypeRule("x-image/*")
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		rule, err := NewMimeTypeRule("*/*")
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(rule.Match("any/thing"))
	}
}
