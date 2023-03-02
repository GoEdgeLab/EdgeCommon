// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeutils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/iwind/TeaGo/maps"
	"time"
)

// EncryptMap 加密
func EncryptMap(nodeUniqueId string, nodeSecret string, data maps.Map, timeout int32) (string, error) {
	if data == nil {
		data = maps.Map{}
	}

	var expiresAt int64
	if timeout > 0 {
		expiresAt = time.Now().Unix() + int64(timeout)
	}

	dataJSON, err := json.Marshal(maps.Map{
		"expiresAt": expiresAt,
		"data":      data,
	})
	if err != nil {
		return "", errors.New("marshal data to json failed: " + err.Error())
	}

	var method = &AES256CFBMethod{}
	err = method.Init([]byte(nodeUniqueId), []byte(nodeSecret))
	if err != nil {
		return "", err
	}
	result, err := method.Encrypt(dataJSON)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(result), nil
}

// DecryptMap 解密
func DecryptMap(nodeUniqueId string, nodeSecret string, encodedString string) (maps.Map, error) {
	var method = &AES256CFBMethod{}
	err := method.Init([]byte(nodeUniqueId), []byte(nodeSecret))
	if err != nil {
		return nil, err
	}

	encodedData, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return nil, errors.New("base64 decode failed: " + err.Error())
	}

	dataJSON, err := method.Decrypt(encodedData)
	if err != nil {
		return nil, err
	}

	var result = maps.Map{}
	err = json.Unmarshal(dataJSON, &result)
	if err != nil {
		return nil, errors.New("unmarshal data failed: " + err.Error())
	}

	var expiresAt = result.GetInt64("expiresAt")
	if expiresAt > 0 && expiresAt < time.Now().Unix() {
		return nil, errors.New("data is expired")
	}

	return result.GetMap("data"), nil
}

// EncryptData 加密
func EncryptData(nodeUniqueId string, nodeSecret string, data []byte) (string, error) {
	if len(data) == 0 {
		return "", nil
	}

	var method = &AES256CFBMethod{}
	err := method.Init([]byte(nodeUniqueId), []byte(nodeSecret))
	if err != nil {
		return "", err
	}
	result, err := method.Encrypt(data)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(result), nil
}

// DecryptData 解密
func DecryptData(nodeUniqueId string, nodeSecret string, encodedString string) ([]byte, error) {
	if len(encodedString) == 0 {
		return nil, nil
	}

	var method = &AES256CFBMethod{}
	err := method.Init([]byte(nodeUniqueId), []byte(nodeSecret))
	if err != nil {
		return nil, err
	}

	encodedData, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return nil, errors.New("base64 decode failed: " + err.Error())
	}

	return method.Decrypt(encodedData)
}
