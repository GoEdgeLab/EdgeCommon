// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/types"
	"io"
	"net"
	"sort"
	"strings"
)

type Reader struct {
	meta *Meta

	ipV4Items []*ipItem
	ipV6Items []*ipItem
}

func NewReader(reader io.Reader) (*Reader, error) {
	var libReader = &Reader{}
	err := libReader.load(reader)
	if err != nil {
		return nil, err
	}
	return libReader, nil
}

func (this *Reader) load(reader io.Reader) error {
	var buf = make([]byte, 1024)
	var metaLine = []byte{}
	var metaLineFound = false
	var dataBuf = []byte{}
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			var data = buf[:n]
			dataBuf = append(dataBuf, data...)
			if metaLineFound {
				left, err := this.parse(dataBuf)
				if err != nil {
					return err
				}
				dataBuf = left
			} else {
				var index = bytes.IndexByte(dataBuf, '\n')
				if index > 0 {
					metaLine = dataBuf[:index]
					dataBuf = dataBuf[index+1:]
					metaLineFound = true
					var meta = &Meta{}
					err = json.Unmarshal(metaLine, &meta)
					if err != nil {
						return err
					}
					meta.Init()
					this.meta = meta

					left, err := this.parse(dataBuf)
					if err != nil {
						return err
					}
					dataBuf = left
				}
			}
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}

	sort.Slice(this.ipV4Items, func(i, j int) bool {
		var from0 = this.ipV4Items[i].ipFrom
		var to0 = this.ipV4Items[i].ipTo
		var from1 = this.ipV4Items[j].ipFrom
		var to1 = this.ipV4Items[j].ipTo
		if from0 == from1 {
			return to0 < to1
		}
		return from0 < from1
	})

	sort.Slice(this.ipV6Items, func(i, j int) bool {
		var from0 = this.ipV6Items[i].ipFrom
		var to0 = this.ipV6Items[i].ipTo
		var from1 = this.ipV6Items[j].ipFrom
		var to1 = this.ipV6Items[j].ipTo
		if from0 == from1 {
			return to0 < to1
		}
		return from0 < from1
	})

	return nil
}

func (this *Reader) Lookup(ip net.IP) *QueryResult {
	var ipLong = configutils.IP2Long(ip)
	var isV4 = configutils.IsIPv4(ip)
	var resultItem *ipItem
	if isV4 {
		sort.Search(len(this.ipV4Items), func(i int) bool {
			var item = this.ipV4Items[i]
			if item.ipFrom <= ipLong {
				if item.ipTo >= ipLong {
					resultItem = item
					return false
				}
				return false
			}
			return true
		})
	} else {
		sort.Search(len(this.ipV6Items), func(i int) bool {
			var item = this.ipV6Items[i]
			if item.ipFrom <= ipLong {
				if item.ipTo >= ipLong {
					resultItem = item
					return false
				}
				return false
			}
			return true
		})
	}

	return &QueryResult{
		item: resultItem,
		meta: this.meta,
	}
}

func (this *Reader) Meta() *Meta {
	return this.meta
}

func (this *Reader) IPv4Items() []*ipItem {
	return this.ipV4Items
}

func (this *Reader) IPv6Items() []*ipItem {
	return this.ipV6Items
}

func (this *Reader) parse(data []byte) (left []byte, err error) {
	if len(data) == 0 {
		return
	}
	for {
		var index = bytes.IndexByte(data, '\n')
		if index >= 0 {
			var line = data[:index]
			var pieces = strings.Split(string(line), "|")
			if len(pieces) != 8 {
				return nil, errors.New("invalid ip definition '" + string(line) + "'")
			}

			var version = pieces[0]
			if len(version) == 0 {
				version = "4"
			}

			if version != "4" && version != "6" {
				return nil, errors.New("invalid ip version '" + string(line) + "'")
			}

			var ipFrom uint64
			var ipTo uint64
			if len(pieces[2]) == 0 {
				pieces[2] = pieces[1]
				ipFrom = types.Uint64(pieces[1])
				ipTo = types.Uint64(pieces[2])
			} else {
				ipFrom = types.Uint64(pieces[1])
				ipTo = types.Uint64(pieces[2]) + ipFrom
			}

			if version == "4" {
				this.ipV4Items = append(this.ipV4Items, &ipItem{
					ipFrom:     ipFrom,
					ipTo:       ipTo,
					countryId:  types.Int64(pieces[3]),
					provinceId: types.Int64(pieces[4]),
					cityId:     types.Int64(pieces[5]),
					townId:     types.Int64(pieces[6]),
					providerId: types.Int64(pieces[7]),
				})
			} else {
				this.ipV6Items = append(this.ipV6Items, &ipItem{
					ipFrom:     ipFrom,
					ipTo:       ipTo,
					countryId:  types.Int64(pieces[3]),
					provinceId: types.Int64(pieces[4]),
					cityId:     types.Int64(pieces[5]),
					townId:     types.Int64(pieces[6]),
					providerId: types.Int64(pieces[7]),
				})
			}

			data = data[index+1:]
		} else {
			left = data
			break
		}
	}
	return
}
