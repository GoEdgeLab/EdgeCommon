// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/iwind/TeaGo/rands"
	"net/http"
	"strings"
	"testing"
)

func TestHTTPAuthRequestMethod_Filter(t *testing.T) {
	method := &HTTPAuthSubRequestMethod{}
	err := method.Init(map[string]interface{}{
		"url":    "http://127.0.0.1:2345/",
		"method": http.MethodPost,
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodGet, "https://teaos.cn/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Hello", "World")
	req.Header.Set("User-Agent", "GoEdge/1.0")
	b, err := method.Filter(req, func(subReq *http.Request) (status int, err error) {
		return
	}, func(s string) string {
		return s
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result:", b)
}

func TestHTTPAuthRequestMethod_Filter_Path(t *testing.T) {
	method := &HTTPAuthSubRequestMethod{}
	err := method.Init(map[string]interface{}{
		"url":    "/hello?${var}",
		"method": http.MethodGet,
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodGet, "http://teaos.cn/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Hello", "World")
	req.Header.Set("User-Agent", "GoEdge/1.0")
	b, err := method.Filter(req, func(subReq *http.Request) (status int, err error) {
		status = rands.Int(200, 400)
		t.Log("execute sub request:", subReq.URL, status)
		return
	}, func(s string) string {
		return strings.ReplaceAll(s, "${var}", "$VAR")
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result:", b)
}
