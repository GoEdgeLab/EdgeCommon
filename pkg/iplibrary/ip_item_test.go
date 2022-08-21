// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestIpItem_AsBinary(t *testing.T) {
	var item = &ipItem{
		IPFrom:     123456789,
		IPTo:       123456790,
		CountryId:  1,
		ProvinceId: 2,
		CityId:     3,
		TownId:     4,
		ProviderId: 5,
	}
	b, err := item.AsBinary()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(b), "bytes")

	var item2 = &ipItem{}
	err = binary.Read(bytes.NewReader(b), binary.BigEndian, item2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(item2)
}
