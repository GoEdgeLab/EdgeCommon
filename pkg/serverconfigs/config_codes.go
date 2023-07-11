// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

type ConfigCode = string

const (
	ConfigCodeUAM            ConfigCode = "uam"
	ConfigCodeCC             ConfigCode = "cc"
	ConfigCodeHostRedirects  ConfigCode = "hostRedirects"
	ConfigCodeLocations      ConfigCode = "locations"
	ConfigCodeRewrites       ConfigCode = "rewrites"
	ConfigCodeWAF            ConfigCode = "waf"
	ConfigCodeCache          ConfigCode = "cache"
	ConfigCodeAuth           ConfigCode = "auth"
	ConfigCodeReferers       ConfigCode = "referers"
	ConfigCodeUserAgent      ConfigCode = "userAgent"
	ConfigCodeCharset        ConfigCode = "charset"
	ConfigCodeAccessLog      ConfigCode = "accessLog"
	ConfigCodeStat           ConfigCode = "stat"
	ConfigCodeCompression    ConfigCode = "compression"
	ConfigCodeOptimization   ConfigCode = "optimization"
	ConfigCodePages          ConfigCode = "pages"
	ConfigCodeHeaders        ConfigCode = "headers"
	ConfigCodeWebsocket      ConfigCode = "websocket"
	ConfigCodeWebp           ConfigCode = "webp"
	ConfigCodeRoot           ConfigCode = "root"
	ConfigCodeFastcgi        ConfigCode = "fastcgi"
	ConfigCodeRemoteAddr     ConfigCode = "remoteAddr"
	ConfigCodeRequestLimit   ConfigCode = "requestLimit"
	ConfigCodeTraffic        ConfigCode = "traffic"
	ConfigCodeRequestScripts ConfigCode = "requestScripts"
	ConfigCodeReverseProxy   ConfigCode = "reverseProxy"
)
