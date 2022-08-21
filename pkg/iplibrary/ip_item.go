// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"encoding/binary"
)

type ipItem struct {
	CountryId  uint32
	ProvinceId uint32
	CityId     uint32
	TownId     uint32
	ProviderId uint32
	IPFrom     uint64
	IPTo       uint64
}

func (this *ipItem) AsBinary() ([]byte, error) {
	var buf = &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, this)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
