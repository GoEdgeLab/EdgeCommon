// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

// ReaderV2 IP库Reader V2
type ReaderV2 struct {
	meta *Meta

	regionMap map[string]*ipRegion // 缓存重复的区域用来节约内存

	ipV4Items []ipv4ItemV2
	ipV6Items []ipv6ItemV2

	lastCountryId  uint16
	lastProvinceId uint16
	lastCityId     uint32
	lastTownId     uint32
	lastProviderId uint16
}

// NewReaderV2 创建新Reader对象
func NewReaderV2(reader io.Reader) (*ReaderV2, error) {
	var libReader = &ReaderV2{
		regionMap: map[string]*ipRegion{},
	}

	if runtime.NumCPU() >= 4 /** CPU数量较多的通常有着大内存 **/ {
		libReader.ipV4Items = make([]ipv4ItemV2, 0, 6_000_000)
	} else {
		libReader.ipV4Items = make([]ipv4ItemV2, 0, 600_000)
	}

	err := libReader.load(reader)
	if err != nil {
		return nil, err
	}
	return libReader, nil
}

// 从Reader中加载数据
func (this *ReaderV2) load(reader io.Reader) error {
	var buf = make([]byte, 1024)
	var metaLine []byte
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
			return bytes.Compare(to0[:], to1[:]) < 0
		}
		return bytes.Compare(from0[:], from1[:]) < 0
	})

	sort.Slice(this.ipV6Items, func(i, j int) bool {
		var from0 = this.ipV6Items[i].IPFrom
		var to0 = this.ipV6Items[i].IPTo
		var from1 = this.ipV6Items[j].IPFrom
		var to1 = this.ipV6Items[j].IPTo
		if from0 == from1 {
			return bytes.Compare(to0[:], to1[:]) < 0
		}
		return bytes.Compare(from0[:], from1[:]) < 0
	})

	// 清理内存
	this.regionMap = nil

	return nil
}

func (this *ReaderV2) Lookup(ip net.IP) *QueryResult {
	if ip == nil {
		return &QueryResult{}
	}

	var isV4 = ip.To4() != nil
	var resultItem any
	if isV4 {
		sort.Search(len(this.ipV4Items), func(i int) bool {
			var item = this.ipV4Items[i]
			if bytes.Compare(item.IPFrom[:], ip) <= 0 {
				if bytes.Compare(item.IPTo[:], ip) >= 0 {
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
			if bytes.Compare(item.IPFrom[:], ip) <= 0 {
				if bytes.Compare(item.IPTo[:], ip) >= 0 {
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

func (this *ReaderV2) Meta() *Meta {
	return this.meta
}

func (this *ReaderV2) IPv4Items() []ipv4ItemV2 {
	return this.ipV4Items
}

func (this *ReaderV2) IPv6Items() []ipv6ItemV2 {
	return this.ipV6Items
}

func (this *ReaderV2) Destroy() {
	this.meta = nil
	this.regionMap = nil
	this.ipV4Items = nil
	this.ipV6Items = nil
}

// 分析数据
func (this *ReaderV2) parse(data []byte) (left []byte, err error) {
	if len(data) == 0 {
		return
	}

	for {
		if len(data) == 0 {
			break
		}

		var offset int
		if data[0] == '|' {
			offset = 1 + 8 + 1
		} else if data[0] == '4' {
			offset = 2 + 8 + 1
		} else if data[0] == '6' {
			offset = 2 + 32 + 1
		}

		var index = bytes.IndexByte(data[offset:], '\n')
		if index >= 0 {
			index += offset
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
func (this *ReaderV2) parseLine(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	const maxPieces = 8
	var pieces []string

	var offset int
	if line[0] == '|' {
		offset = 1 + 8 + 1
		pieces = append(pieces, "", string(line[1:5]), string(line[5:9]))
	} else if line[0] == '4' {
		offset = 2 + 8 + 1
		pieces = append(pieces, "", string(line[2:6]), string(line[6:10]))
	} else if line[0] == '6' {
		offset = 2 + 32 + 1
		pieces = append(pieces, "6", string(line[2:18]), string(line[18:34]))
	}

	pieces = append(pieces, strings.Split(string(line[offset:]), "|")...)

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
	var ipFromV4 [4]byte
	var ipToV4 [4]byte

	var ipFromV6 [16]byte
	var ipToV6 [16]byte

	if version == "6" {
		ipFromV6 = [16]byte([]byte(pieces[1]))
		ipToV6 = [16]byte([]byte(pieces[2]))
	} else {
		ipFromV4 = [4]byte([]byte(pieces[1]))
		ipToV4 = [4]byte([]byte(pieces[2]))
	}

	// country
	var countryId uint16
	if pieces[3] == "+" {
		countryId = this.lastCountryId
	} else {
		countryId = uint16(this.decodeUint64(pieces[3]))
	}
	this.lastCountryId = countryId

	var provinceId uint16
	if pieces[4] == "+" {
		provinceId = this.lastProvinceId
	} else {
		provinceId = uint16(this.decodeUint64(pieces[4]))
	}
	this.lastProvinceId = provinceId

	// city
	var cityId uint32
	if pieces[5] == "+" {
		cityId = this.lastCityId
	} else {
		cityId = uint32(this.decodeUint64(pieces[5]))
	}
	this.lastCityId = cityId

	// town
	var townId uint32
	if pieces[6] == "+" {
		townId = this.lastTownId
	} else {
		townId = uint32(this.decodeUint64(pieces[6]))
	}
	this.lastTownId = townId

	// provider
	var providerId uint16
	if pieces[7] == "+" {
		providerId = this.lastProviderId
	} else {
		providerId = uint16(this.decodeUint64(pieces[7]))
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
		this.ipV4Items = append(this.ipV4Items, ipv4ItemV2{
			IPFrom: ipFromV4,
			IPTo:   ipToV4,
			Region: region,
		})
	} else {
		this.ipV6Items = append(this.ipV6Items, ipv6ItemV2{
			IPFrom: ipFromV6,
			IPTo:   ipToV6,
			Region: region,
		})
	}

	return nil
}

func (this *ReaderV2) decodeUint64(s string) uint64 {
	if this.meta != nil && this.meta.Version == Version2 {
		i, _ := strconv.ParseUint(s, 32, 64)
		return i
	}
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}
