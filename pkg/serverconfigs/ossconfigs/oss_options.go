// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package ossconfigs

type OSSOptions interface {
	Init() error      // 初始化
	Summary() string  // 内容简述
	UniqueId() string // 唯一标识
}
