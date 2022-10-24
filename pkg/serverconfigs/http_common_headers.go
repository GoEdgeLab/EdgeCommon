// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

import (
	"net/http"
	"strings"
)

var HTTPCommonRequestHeaders = []string{
	"A-IM",
	"Accept",
	"Accept-Charset",
	"Accept-Datetime",
	"Accept-Encoding",
	"Accept-Language",
	"Access-Control-Request-Method",
	"Access-Control-Request-Headers",
	"Authorization",
	"Cache-Control",
	"Connection",
	"Content-Encoding",
	"Content-Length",
	"Content-MD5",
	"Content-Type",
	"Cookie",
	"Date",
	"Expect",
	"Forwarded",
	"From",
	"Host",
	"HTTP2-Settings",
	"If-Match",
	"If-Modified-Since",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
	"Max-Forwards",
	"Origin",
	"Pragma",
	"Prefer",
	"Proxy-Authorization",
	"Range",
	"Referer",
	"TE",
	"Trailer",
	"Transfer-Encoding",
	"User-Agent",
	"Upgrade",
	"Via",
	"Warning",
	"Keep-Alive",
}

var HTTPCommonRequestHeaders2 = []string{
	"Upgrade-Insecure-Requests",
	"X-Requested-With",
	"DNT",
	"X-Forwarded-For",
	"X-Forwarded-Host",
	"X-Forwarded-Proto",
	"X-Forwarded-By",
	"X-Real-IP",
	"Front-End-Https",
	"X-Http-Method-Override",
	"X-ATT-DeviceId",
	"X-Wap-Profile",
	"Proxy-Connection",
	"X-UIDH",
	"X-Csrf-Token",
	"X-Request-ID",
	"X-Correlation-ID",
	"Correlation-ID",
	"Save-Data",
	"Device-Memory",
	"Downlink",
	"Early-Data",
	"ECT",
	"RTT",
	"Sec-CH-UA",
	"Sec-CH-UA-Arch",
	"Sec-CH-UA-Bitness",
	"Sec-CH-UA-Full-Version",
	"Sec-CH-UA-Full-Version-List",
	"Sec-CH-UA-Mobile",
	"Sec-CH-UA-Model",
	"Sec-CH-UA-Platform",
	"Sec-CH-UA-Platform-Version",
	"Sec-Fetch-Dest",
	"Sec-Fetch-Mode",
	"Sec-Fetch-Site",
	"Sec-Fetch-User",
	"Sec-GPC",
	"Service-Worker-Navigation-Preload",
	"Viewport-Width",
	"Want-Digest",
}

var AllHTTPCommonRequestHeaders = append(append([]string{}, HTTPCommonRequestHeaders...), HTTPCommonRequestHeaders2...)

var HTTPCommonResponseHeaders = []string{
	"Accept-CH",
	"Access-Control-Allow-Origin",
	"Access-Control-Allow-Credentials",
	"Access-Control-Expose-Headers",
	"Access-Control-Max-Age",
	"Access-Control-Allow-Methods",
	"Access-Control-Allow-Headers",
	"Accept-Patch",
	"Accept-Post",
	"Accept-Ranges",
	"Age",
	"Allow",
	"Alt-Svc",
	"Cache-Control",
	"Clear-Site-Data",
	"Connection",
	"Content-Disposition",
	"Content-Encoding",
	"Content-Language",
	"Content-Length",
	"Content-Location",
	"Content-MD5",
	"Content-Range",
	"Content-Type",
	"Date",
	"Delta-Base",
	"ETag",
	"Expires",
	"IM",
	"Last-Modified",
	"Link",
	"Location",
	"P3P",
	"Pragma",
	"Preference-Applied",
	"Proxy-Authenticate",
	"Public-Key-Pins",
	"Retry-After",
	"Server",
	"Server-Timing",
	"Set-Cookie",
	"Strict-Transport-Security",
	"Trailer",
	"Transfer-Encoding",
	"Tk",
	"Upgrade",
	"Vary",
	"Via",
	"Warning",
	"WWW-Authenticate",
	"X-Frame-Options",
	"Keep-Alive",
	"Referrer-Policy",
}

var HTTPCommonResponseHeaders2 = []string{
	"Content-Security-Policy",
	"X-Content-Security-Policy",
	"Content-Security-Policy-Report-Only",
	"Cross-Origin-Embedder-Policy",
	"Cross-Origin-Opener-Policy",
	"Cross-Origin-Resource-Policy",
	"Digest",
	"X-WebKit-CSP",
	"Expect-CT",
	"NEL",
	"Permissions-Policy",
	"Refresh",
	"Report-To",
	"Status",
	"Timing-Allow-Origin",
	"X-Content-Duration",
	"X-Content-Type-Options",
	"X-Powered-By",
	"X-Redirect-By",
	"X-Request-ID",
	"X-Correlation-ID",
	"X-UA-Compatible",
	"X-XSS-Protection",
	"Sec-WebSocket-Accept",
	"SourceMap",
	"X-DNS-Prefetch-Control",
}

var AllHTTPCommonResponseHeaders = append(append([]string{}, HTTPCommonResponseHeaders...), HTTPCommonResponseHeaders2...)

var allRequestHeaderMap = map[string]struct{}{}

func init() {
	for _, headerName := range AllHTTPCommonRequestHeaders {
		allRequestHeaderMap[headerName] = struct{}{}
		allRequestHeaderMap[strings.ToLower(headerName)] = struct{}{}
		allRequestHeaderMap[http.CanonicalHeaderKey(headerName)] = struct{}{}
	}
}

// IsCommonRequestHeader 判断某个HTTP请求Header名称是否为通用
func IsCommonRequestHeader(headerName string) bool {
	_, ok := allRequestHeaderMap[headerName]
	return ok
}
