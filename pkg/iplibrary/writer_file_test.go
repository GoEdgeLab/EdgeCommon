// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/iplibrary"
	"github.com/iwind/TeaGo/rands"
	"github.com/iwind/TeaGo/types"
	"testing"
)

func TestNewFileWriter(t *testing.T) {
	writer, err := iplibrary.NewFileWriter("./internal-ip-library.db", &iplibrary.Meta{
		Author: "GoEdge",
	})
	if err != nil {
		t.Fatal(err)
	}

	err = writer.WriteMeta()
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.1.100", "192.168.1.100", 100, 200, 300, 400, 500)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.2.100", "192.168.3.100", 101, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.3.101", "192.168.3.101", 101, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	var n = func() string {
		return types.String(rands.Int(0, 255))
	}

	for i := 0; i < 1_000_000; i++ {
		err = writer.Write(n()+"."+n()+"."+n()+"."+n(), n()+"."+n()+"."+n()+"."+n(), int64(i)+100, 201, 301, 401, 501)
		if err != nil {
			t.Fatal(err)
		}
	}

	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok", writer.Sum())
}
