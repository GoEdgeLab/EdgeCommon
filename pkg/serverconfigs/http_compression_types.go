// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

type HTTPCompressionType = string

const (
	HTTPCompressionTypeGzip    HTTPCompressionType = "gzip"
	HTTPCompressionTypeDeflate HTTPCompressionType = "deflate"
	HTTPCompressionTypeBrotli  HTTPCompressionType = "brotli"
	HTTPCompressionTypeZSTD    HTTPCompressionType = "zstd"
)
