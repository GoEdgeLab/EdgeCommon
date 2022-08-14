// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/types"
	"hash"
	"io"
	"net"
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

type Writer struct {
	writer *hashWriter
	meta   *Meta
}

func NewWriter(writer io.Writer, meta *Meta) *Writer {
	if meta == nil {
		meta = &Meta{}
	}
	meta.Version = Version1
	meta.CreatedAt = time.Now().Unix()

	var libWriter = &Writer{
		writer: newHashWriter(writer),
		meta:   meta,
	}
	return libWriter
}

func (this *Writer) WriteMeta() error {
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

func (this *Writer) Write(ipFrom string, ipTo string, countryId int64, provinceId int64, cityId int64, townId int64, providerId int64) error {
	// validate IP
	var fromIP = net.ParseIP(ipFrom)
	if fromIP == nil {
		return errors.New("invalid 'ipFrom': '" + ipFrom + "'")
	}
	var fromIsIPv4 = configutils.IsIPv4(fromIP)
	var toIP = net.ParseIP(ipTo)
	if toIP == nil {
		return errors.New("invalid 'ipTo': " + ipTo)
	}
	var toIsIPv4 = configutils.IsIPv4(toIP)
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
	var fromIPLong = configutils.IP2Long(fromIP)
	var toIPLong = configutils.IP2Long(toIP)

	if toIPLong < fromIPLong {
		fromIPLong, toIPLong = toIPLong, fromIPLong
	}

	pieces = append(pieces, types.String(fromIPLong))
	if ipFrom == ipTo {
		// 2
		pieces = append(pieces, "")
	} else {
		// 2
		pieces = append(pieces, types.String(toIPLong-fromIPLong))
	}

	// 3
	if countryId > 0 {
		pieces = append(pieces, types.String(countryId))
	} else {
		pieces = append(pieces, "")
	}

	// 4
	if provinceId > 0 {
		pieces = append(pieces, types.String(provinceId))
	} else {
		pieces = append(pieces, "")
	}

	// 5
	if cityId > 0 {
		pieces = append(pieces, types.String(cityId))
	} else {
		pieces = append(pieces, "")
	}

	// 6
	if townId > 0 {
		pieces = append(pieces, types.String(townId))
	} else {
		pieces = append(pieces, "")
	}

	// 7
	if providerId > 0 {
		pieces = append(pieces, types.String(providerId))
	} else {
		pieces = append(pieces, "")
	}

	_, err := this.writer.Write([]byte(strings.Join(pieces, "|")))
	if err != nil {
		return err
	}

	_, err = this.writer.Write([]byte("\n"))
	return err
}

func (this *Writer) Sum() string {
	return this.writer.Sum()
}
