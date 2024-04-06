// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

type hashWriter struct {
	rawWriter io.Writer
	hash      hash.Hash
}

func newHashWriter(writer io.Writer) *hashWriter {
	return &hashWriter{
		rawWriter: writer,
		hash:      md5.New(),
	}
}

func (this *hashWriter) Write(p []byte) (n int, err error) {
	n, err = this.rawWriter.Write(p)
	this.hash.Write(p)
	return
}

func (this *hashWriter) Sum() string {
	return fmt.Sprintf("%x", this.hash.Sum(nil))
}

type WriterV2 struct {
	writer *hashWriter
	meta   *Meta

	lastCountryId  int64
	lastProvinceId int64
	lastCityId     int64
	lastTownId     int64
	lastProviderId int64
}

func NewWriterV2(writer io.Writer, meta *Meta) *WriterV2 {
	if meta == nil {
		meta = &Meta{}
	}
	meta.Version = Version2
	meta.CreatedAt = time.Now().Unix()

	var libWriter = &WriterV2{
		writer: newHashWriter(writer),
		meta:   meta,
	}
	return libWriter
}

func (this *WriterV2) WriteMeta() error {
	metaJSON, err := json.Marshal(this.meta)
	if err != nil {
		return err
	}
	_, err = this.writer.Write(metaJSON)
	if err != nil {
		return err
	}
	_, err = this.writer.Write([]byte("\n"))
	return err
}

func (this *WriterV2) Write(ipFrom string, ipTo string, countryId int64, provinceId int64, cityId int64, townId int64, providerId int64) error {
	// validate IP
	var fromIP = net.ParseIP(ipFrom)
	if fromIP == nil {
		return errors.New("invalid 'ipFrom': '" + ipFrom + "'")
	}
	var fromIsIPv4 = fromIP.To4() != nil
	var toIP = net.ParseIP(ipTo)
	if toIP == nil {
		return errors.New("invalid 'ipTo': " + ipTo)
	}
	var toIsIPv4 = toIP.To4() != nil
	if fromIsIPv4 != toIsIPv4 {
		return errors.New("'ipFrom(" + ipFrom + ")' and 'ipTo(" + ipTo + ")' should have the same IP version")
	}

	var pieces = []string{}

	// 0
	if fromIsIPv4 {
		pieces = append(pieces, "")
	} else {
		pieces = append(pieces, "6")
	}

	// 1
	if bytes.Compare(fromIP, toIP) > 0 {
		fromIP, toIP = toIP, fromIP
	}

	if fromIsIPv4 {
		pieces = append(pieces, string(fromIP.To4())+string(toIP.To4()))
	} else {
		pieces = append(pieces, string(fromIP.To16())+string(toIP.To16()))
	}

	// 2
	if countryId > 0 {
		if countryId == this.lastCountryId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(countryId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastCountryId = countryId

	// 3
	if provinceId > 0 {
		if provinceId == this.lastProvinceId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(provinceId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastProvinceId = provinceId

	// 4
	if cityId > 0 {
		if cityId == this.lastCityId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(cityId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastCityId = cityId

	// 5
	if townId > 0 {
		if townId == this.lastTownId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(townId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastTownId = townId

	// 6
	if providerId > 0 {
		if providerId == this.lastProviderId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(providerId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastProviderId = providerId

	_, err := this.writer.Write([]byte(strings.TrimRight(strings.Join(pieces, "|"), "|")))
	if err != nil {
		return err
	}

	_, err = this.writer.Write([]byte("\n"))
	return err
}

func (this *WriterV2) Sum() string {
	return this.writer.Sum()
}

func (this *WriterV2) formatUint64(i uint64) string {
	return strconv.FormatUint(i, 32)
}
