// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package ossconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type OSSType = string

func FindAllOSSTypes() []*shared.Definition {
	return []*shared.Definition{}
}

func DecodeOSSOptions(ossType OSSType, optionsJSON []byte) (any, error) {
	return nil, nil
}
