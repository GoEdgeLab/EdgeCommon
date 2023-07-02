// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package ossconfigs

import "errors"

type OSSType = string

type OSSTypeDefinition struct {
	Name             string `json:"name"`
	Code             string `json:"code"`
	BucketOptionName string `json:"bucketOptionName"`
	BucketIgnored    bool   `json:"bucketIgnored"` // 是否忽略Bucket名称
}

func FindAllOSSTypes() []*OSSTypeDefinition {
	return []*OSSTypeDefinition{}
}

func FindOSSType(code string) *OSSTypeDefinition {
	for _, t := range FindAllOSSTypes() {
		if t.Code == code {
			return t
		}
	}
	return nil
}

func DecodeOSSOptions(ossType OSSType, optionsJSON []byte) (any, error) {
	return nil, errors.New("'" + ossType + "' has not been supported")
}
