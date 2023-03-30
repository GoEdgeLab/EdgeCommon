// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
	"strings"
	"time"
)

type UpdaterSource interface {
	// DataDir 文件目录
	DataDir() string

	// FindLatestFile 检查最新的IP库文件
	FindLatestFile() (code string, fileId int64, err error)

	// DownloadFile 下载文件
	DownloadFile(fileId int64, writer io.Writer) error

	// LogInfo 普通日志
	LogInfo(message string)

	// LogError 错误日志
	LogError(err error)
}

type Updater struct {
	source UpdaterSource

	currentCode string
	ticker      *time.Ticker

	isUpdating bool
}

func NewUpdater(source UpdaterSource, interval time.Duration) *Updater {
	return &Updater{
		source: source,
		ticker: time.NewTicker(interval),
	}
}

func (this *Updater) Start() {
	// 初始化
	err := this.Init()
	if err != nil {
		this.source.LogError(err)
	}

	// 先运行一次
	err = this.Loop()
	if err != nil {
		this.source.LogError(err)
	}

	// 开始定时运行
	for range this.ticker.C {
		err = this.Loop()
		if err != nil {
			this.source.LogError(err)
		}
	}
}

func (this *Updater) Init() error {
	// 检查当前正在使用的IP库
	var path = this.source.DataDir() + "/ip-library.db"
	fp, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return errors.New("read ip library file failed '" + err.Error() + "'")
	}
	defer func() {
		_ = fp.Close()
	}()

	return this.loadFile(fp)
}

func (this *Updater) Loop() error {
	if this.isUpdating {
		return nil
	}

	this.isUpdating = true

	defer func() {
		this.isUpdating = false
	}()

	code, fileId, err := this.source.FindLatestFile()
	if err != nil {
		// 不提示连接错误
		if this.isConnError(err) {
			return nil
		}
		return err
	}
	if len(code) == 0 || fileId <= 0 {
		// 还原到内置IP库
		if len(this.currentCode) > 0 {
			this.currentCode = ""
			this.source.LogInfo("resetting to default ip library ...")

			var defaultPath = this.source.DataDir() + "/ip-library.db"
			_, err = os.Stat(defaultPath)
			if err == nil {
				err = os.Remove(defaultPath)
				if err != nil {
					this.source.LogError(errors.New("can not remove default 'ip-library.db'"))
				}
			}

			err = InitDefault()
			if err != nil {
				this.source.LogError(errors.New("initialize default ip library failed: " + err.Error()))
			}
		}

		return nil
	}

	// 下载
	if this.currentCode == code {
		// 不再重复下载
		return nil
	}

	// 检查是否存在
	var dir = this.source.DataDir()
	var path = dir + "/ip-" + code + ".db"
	stat, err := os.Stat(path)
	if err == nil && !stat.IsDir() && stat.Size() > 0 {
		fp, err := os.Open(path)
		if err != nil {
			return err
		}

		defer func() {
			_ = fp.Close()
		}()

		err = this.loadFile(fp)
		if err != nil {
			// 尝试删除
			_ = os.Remove(path)
		} else {
			this.currentCode = code

			// 拷贝到 ip-library.db
			err = this.createDefaultFile(path, dir)
			if err != nil {
				this.source.LogError(err)
			}
		}
		return err
	}

	// write to file
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return errors.New("create ip library file failed: " + err.Error())
	}

	var isOk = false
	defer func() {
		if !isOk {
			_ = os.Remove(path)
		}
	}()

	err = this.source.DownloadFile(fileId, fp)
	if err != nil {
		_ = fp.Close()
		return err
	}
	err = fp.Close()
	if err != nil {
		return nil
	}

	// load library from file
	fp, err = os.Open(path)
	if err != nil {
		return nil
	}
	err = this.loadFile(fp)
	_ = fp.Close()
	if err != nil {
		return errors.New("load file failed: " + err.Error())
	}

	isOk = true
	this.currentCode = code

	// 拷贝到 ip-library.db
	err = this.createDefaultFile(path, dir)
	if err != nil {
		this.source.LogError(err)
	}

	return nil
}

func (this *Updater) loadFile(fp *os.File) error {
	this.source.LogInfo("load ip library from '" + fp.Name() + "' ...")

	fileReader, err := NewFileDataReader(fp, "")
	if err != nil {
		return errors.New("load ip library from reader failed: " + err.Error())
	}

	var reader = fileReader.RawReader()
	defaultLibrary = NewIPLibraryWithReader(reader)
	this.currentCode = reader.Meta().Code
	return nil
}

func (this *Updater) createDefaultFile(sourcePath string, dir string) error {
	sourceFp, err := os.Open(sourcePath)
	if err != nil {
		return errors.New("prepare to copy file to 'ip-library.db' failed: " + err.Error())
	}
	defer func() {
		_ = sourceFp.Close()
	}()

	dstFp, err := os.Create(dir + "/ip-library.db")
	if err != nil {
		return errors.New("prepare to copy file to 'ip-library.db' failed: " + err.Error())
	}
	defer func() {
		_ = dstFp.Close()
	}()
	_, err = io.Copy(dstFp, sourceFp)
	if err != nil {
		return errors.New("copy file to 'ip-library.db' failed: " + err.Error())
	}
	return nil
}

// isConnError 是否为连接错误
func (this *Updater) isConnError(err error) bool {
	if err == nil {
		return false
	}

	// 检查是否为连接错误
	statusErr, ok := status.FromError(err)
	if ok {
		var errorCode = statusErr.Code()
		return errorCode == codes.Unavailable || errorCode == codes.Canceled
	}

	if strings.Contains(err.Error(), "code = Canceled") {
		return true
	}

	return false
}
