// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package shared

import "github.com/iwind/TeaGo/maps"

// DefaultRequestVariables 默认的请求变量列表
func DefaultRequestVariables() []maps.Map {
	return []maps.Map{
		{"code": "${edgeVersion}", "name": "边缘节点版本", "description": ""},
		{"code": "${remoteAddr}", "name": "客户端地址（IP）", "description": "会依次根据X-Forwarded-For、X-Real-IP、RemoteAddr获取，适合前端有别的反向代理服务时使用，存在伪造的风险"},
		{"code": "${rawRemoteAddr}", "name": "客户端地址（IP）", "description": "返回直接连接服务的客户端原始IP地址"},
		{"code": "${remotePort}", "name": "客户端端口", "description": ""},
		{"code": "${remoteUser}", "name": "客户端用户名", "description": ""},
		{"code": "${requestURI}", "name": "请求URI", "description": "比如/hello?name=lily"},
		{"code": "${requestPath}", "name": "请求路径（不包括参数）", "description": "比如/hello"},
		{"code": "${requestURL}", "name": "完整的请求URL", "description": "比如https://example.com/hello?name=lily"},
		{"code": "${requestLength}", "name": "请求内容长度", "description": ""},
		{"code": "${requestMethod}", "name": "请求方法", "description": "比如GET、POST"},
		{"code": "${requestFilename}", "name": "请求文件路径", "description": ""},
		{"code": "${requestPathExtension}", "name": "请求文件扩展名", "description": "请求路径中的文件扩展名，包括点符号，比如.html、.png"},
		{"code": "${requestPathLowerExtension}", "name": "请求文件小写扩展名", "description": "请求路径中的文件扩展名，其中大写字母会被自动转换为小写，包括点符号，比如.html、.png"},
		{"code": "${scheme}", "name": "请求协议，http或https", "description": ""},
		{"code": "${proto}", "name": "包含版本的HTTP请求协议", "description:": "类似于HTTP/1.0"},
		{"code": "${timeISO8601}", "name": "ISO 8601格式的时间", "description": "比如2018-07-16T23:52:24.839+08:00"},
		{"code": "${timeLocal}", "name": "本地时间", "description": "比如17/Jul/2018:09:52:24 +0800"},
		{"code": "${msec}", "name": "带有毫秒的时间", "description": "比如1531756823.054"},
		{"code": "${timestamp}", "name": "unix时间戳，单位为秒", "description": ""},
		{"code": "${host}", "name": "主机名", "description": ""},
		{"code": "${cname}", "name": "当前网站的CNAME", "description": "比如38b48e4f.goedge.cn"},
		{"code": "${serverName}", "name": "接收请求的服务器名", "description": ""},
		{"code": "${serverPort}", "name": "接收请求的服务器端口", "description": ""},
		{"code": "${referer}", "name": "请求来源URL", "description": ""},
		{"code": "${referer.host}", "name": "请求来源URL域名", "description": ""},
		{"code": "${userAgent}", "name": "客户端信息", "description": ""},
		{"code": "${contentType}", "name": "请求头部的Content-Type", "description": ""},
		{"code": "${cookies}", "name": "所有cookie组合字符串", "description": ""},
		{"code": "${cookie.NAME}", "name": "单个cookie值", "description": ""},
		{"code": "${isArgs}", "name": "问号（?）标记", "description": "如果URL有参数，则值为`?`；否则，则值为空"},
		{"code": "${args}", "name": "所有参数组合字符串", "description": ""},
		{"code": "${arg.NAME}", "name": "单个参数值", "description": ""},
		{"code": "${headers}", "name": "所有Header信息组合字符串", "description": ""},
		{"code": "${header.NAME}", "name": "单个Header值", "description": ""},
		{"code": "${geo.country.name}", "name": "国家/地区名称", "description": ""},
		{"code": "${geo.country.id}", "name": "国家/地区ID", "description": ""},
		{"code": "${geo.province.name}", "name": "省份名称", "description": "目前只包含中国省份"},
		{"code": "${geo.province.id}", "name": "省份ID", "description": "目前只包含中国省份"},
		{"code": "${geo.city.name}", "name": "城市名称", "description": "目前只包含中国城市"},
		{"code": "${geo.city.id}", "name": "城市名称", "description": "目前只包含中国城市"},
		{"code": "${isp.name}", "name": "ISP服务商名称", "description": ""},
		{"code": "${isp.id}", "name": "ISP服务商ID", "description": ""},
		{"code": "${browser.os.name}", "name": "操作系统名称", "description": "客户端所在操作系统名称"},
		{"code": "${browser.os.version}", "name": "操作系统版本", "description": "客户端所在操作系统版本"},
		{"code": "${browser.name}", "name": "浏览器名称", "description": "客户端浏览器名称"},
		{"code": "${browser.version}", "name": "浏览器版本", "description": "客户端浏览器版本"},
		{"code": "${browser.isMobile}", "name": "手机标识", "description": "如果客户端是手机，则值为1，否则为0"},
	}
}
