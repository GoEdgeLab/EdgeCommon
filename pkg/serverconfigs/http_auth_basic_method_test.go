// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"encoding/base64"
	"github.com/iwind/TeaGo/assert"
	"github.com/iwind/TeaGo/maps"
	stringutil "github.com/iwind/TeaGo/utils/string"
	"net/http"
	"testing"
)

func TestHTTPAuthBasicMethodUser_Validate(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		user := &HTTPAuthBasicMethodUser{
			Password: "123456",
		}
		b, err := user.Validate("123456")
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(b)
	}

	{
		user := &HTTPAuthBasicMethodUser{
			Password: "654321",
		}
		b, err := user.Validate("123456")
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(b)
	}

	{
		user := &HTTPAuthBasicMethodUser{
			Password: stringutil.Md5("123456"),
			Encoder:  "md5",
		}
		b, err := user.Validate("123456")
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(b)
	}

	{
		user := &HTTPAuthBasicMethodUser{
			Password: base64.StdEncoding.EncodeToString([]byte("123456")),
			Encoder:  "base64",
		}
		b, err := user.Validate("123456")
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(b)
	}
}

func TestHTTPAuthBasicMethod_Filter(t *testing.T) {

	method := &HTTPAuthBasicMethod{}
	err := method.Init(map[string]interface{}{
		"users": []maps.Map{
			{
				"username": "hello",
				"password": "world",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodGet, "http://teaos.cn/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("hello:world")))
	t.Log(method.Filter(req, nil, nil))
}
