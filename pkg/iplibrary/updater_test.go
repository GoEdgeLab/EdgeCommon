// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/iplibrary"
	"github.com/iwind/TeaGo/Tea"
	_ "github.com/iwind/TeaGo/bootstrap"
	"io"
	"testing"
	"time"
)

type updaterSource struct {
	t *testing.T
}

func (this *updaterSource) DataDir() string {
	return Tea.Root + "/data"
}

func (this *updaterSource) FindLatestFile() (code string, fileId int64, err error) {
	return "CODE", 1, nil
}

func (this *updaterSource) DownloadFile(fileId int64, writer io.Writer) error {
	this.t.Log("downloading file:", fileId, "writer:", writer)
	_, err := writer.Write(iplibrary.DefaultIPLibraryData())
	return err
}

func (this *updaterSource) LogInfo(message string) {
	this.t.Log(message)
}

func (this *updaterSource) LogError(err error) {
	this.t.Fatal(err)
}

func TestNewUpdater(t *testing.T) {
	var updater = iplibrary.NewUpdater(&updaterSource{
		t: t,
	}, 1*time.Minute)
	err := updater.Init()
	if err != nil {
		t.Fatal(err)
	}

	err = updater.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
