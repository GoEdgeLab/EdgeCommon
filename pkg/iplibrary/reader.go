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

// Reader IP库Reader
type Reader struct {
	meta *Meta

	regionMap map[string]*ipRegion

	ipV4Items []*ipItem
	ipV6Items []*ipItem

	lastIPFrom     uint64
	lastCountryId  uint32
	lastProvinceId uint32
	lastCityId     uint32
	lastTownId     uint32
	lastProviderId uint32
}

// NewReader 创建新Reader对象
func NewReader(reader io.Reader) (*Reader, error) {
	var libReader = &Reader{
		regionMap: map[string]*ipRegion{},
	}
	err := libReader.load(reader)
	if err != nil {
		return nil, err
	}
	return libReader, nil
}

// 从Reader中加载数据
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
		var from0 = this.ipV4Items[i].IPFrom
		var to0 = this.ipV4Items[i].IPTo
		var from1 = this.ipV4Items[j].IPFrom
		var to1 = this.ipV4Items[j].IPTo
		if from0 == from1 {
			return to0 < to1
		}
		return from0 < from1
	})

	sort.Slice(this.ipV6Items, func(i, j int) bool {
		var from0 = this.ipV6Items[i].IPFrom
		var to0 = this.ipV6Items[i].IPTo
		var from1 = this.ipV6Items[j].IPFrom
		var to1 = this.ipV6Items[j].IPTo
		if from0 == from1 {
			return to0 < to1
		}
		return from0 < from1
	})

	return nil
}

func (this *Reader) Lookup(ip net.IP) *QueryResult {
	if ip == nil {
		return &QueryResult{}
	}

	var ipLong = configutils.IP2Long(ip)
	var isV4 = configutils.IsIPv4(ip)
	var resultItem *ipItem
	if isV4 {
		sort.Search(len(this.ipV4Items), func(i int) bool {
			var item = this.ipV4Items[i]
			if item.IPFrom <= ipLong {
				if item.IPTo >= ipLong {
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
			if item.IPFrom <= ipLong {
				if item.IPTo >= ipLong {
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

// 分析数据
func (this *Reader) parse(data []byte) (left []byte, err error) {
	if len(data) == 0 {
		return
	}

	for {
		var index = bytes.IndexByte(data, '\n')
		if index >= 0 {
			var line = data[:index]
			err = this.parseLine(line)
			if err != nil {
				return nil, err
			}
			data = data[index+1:]
		} else {
			left = data
			break
		}
	}
	return
}

// 单行分析
func (this *Reader) parseLine(line []byte) error {
	const maxPieces = 8
	var pieces = strings.Split(string(line), "|")
	var countPieces = len(pieces)
	if countPieces < maxPieces { // 补足一行
		for i := 0; i < maxPieces-countPieces; i++ {
			pieces = append(pieces, "")
		}
	} else if countPieces > maxPieces {
		return errors.New("invalid ip definition '" + string(line) + "'")
	}

	var version = pieces[0]
	if len(version) == 0 {
		version = "4"
	}

	if version != "4" && version != "6" {
		return errors.New("invalid ip version '" + string(line) + "'")
	}

	// ip range
	var ipFrom uint64
	var ipTo uint64
	if strings.HasPrefix(pieces[1], "+") {
		ipFrom = this.lastIPFrom + types.Uint64(pieces[1][1:])
	} else {
		ipFrom = types.Uint64(pieces[1])
	}
	if len(pieces[2]) == 0 {
		ipTo = ipFrom
	} else {
		ipTo = types.Uint64(pieces[2]) + ipFrom
	}
	this.lastIPFrom = ipFrom

	// country
	var countryId uint32
	if pieces[3] == "+" {
		countryId = this.lastCountryId
	} else {
		countryId = types.Uint32(pieces[3])
	}
	this.lastCountryId = countryId

	var provinceId uint32
	if pieces[4] == "+" {
		provinceId = this.lastProvinceId
	} else {
		provinceId = types.Uint32(pieces[4])
	}
	this.lastProvinceId = provinceId

	// city
	var cityId uint32
	if pieces[5] == "+" {
		cityId = this.lastCityId
	} else {
		cityId = types.Uint32(pieces[5])
	}
	this.lastCityId = cityId

	// town
	var townId uint32
	if pieces[6] == "+" {
		townId = this.lastTownId
	} else {
		townId = types.Uint32(pieces[6])
	}
	this.lastTownId = townId

	// provider
	var providerId uint32
	if pieces[7] == "+" {
		providerId = this.lastProviderId
	} else {
		providerId = types.Uint32(pieces[7])
	}
	this.lastProviderId = providerId

	var hash = HashRegion(countryId, provinceId, cityId, townId, providerId)

	region, ok := this.regionMap[hash]
	if !ok {
		region = &ipRegion{
			CountryId:  countryId,
			ProvinceId: provinceId,
			CityId:     cityId,
			TownId:     townId,
			ProviderId: providerId,
		}
		this.regionMap[hash] = region
	}

	if version == "4" {
		this.ipV4Items = append(this.ipV4Items, &ipItem{
			IPFrom: ipFrom,
			IPTo:   ipTo,
			Region: region,
		})
	} else {
		this.ipV6Items = append(this.ipV6Items, &ipItem{
			IPFrom: ipFrom,
			IPTo:   ipTo,
			Region: region,
		})
	}

	return nil
}
