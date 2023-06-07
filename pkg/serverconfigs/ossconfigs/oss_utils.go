// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package ossconfigs

import "strings"

func IsOSSProtocol(protocol string) bool {
	return strings.HasPrefix(protocol, "oss:")
}
