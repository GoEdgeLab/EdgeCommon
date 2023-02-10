# 缓存条件设置
## 定义
~~~json
{
  "isOn": "是否启用配置",
  "key": "每个缓存的Key规则，里面可以有变量",
  "life": "缓存时长",
  "expiresTime": "客户端过期时间",
  "status": ["缓存的状态码1", "缓存的状态码2", ...],
  "minSize": "能够缓存的最小尺寸",
  "maxSize": "能够缓存的最大尺寸",
  "methods": ["支持的请求方法1", "支持的请求方法2", ...],
  "skipCacheControlValues": "可以跳过的响应的Cache-Control值",
  "skipSetCookie": "是否跳过响应的Set-Cookie Header",
  "enableRequestCachePragma": "是否支持客户端的Pragma: no-cache",
  "allowChunkedEncoding": "是否允许分片内容",
  "allowPartialContent": "支持分段内容缓存",
  "conds": "请求条件",
  "isReverse": "是否为反向条件，反向条件的不缓存"
}
~~~

## 示例
~~~json
{
	"isOn": true,
	"key": "${scheme}://${host}${requestURI}",
	"life": {
		"count": 1,
		"unit": "day"
	},
	"expiresTime": {
		"isPrior": true,
		"isOn": true,
		"overwrite": true,
		"autoCalculate": false,
		"duration": {
			"count": 1,
			"unit": "day"
		}
	},
	"status": [
		200
	],
	"minSize": {
		"count": 0,
		"unit": "kb"
	},
	"maxSize": {
		"count": 32,
		"unit": "mb"
	},
	"methods": [],
	"skipCacheControlValues": [
		"private",
		"no-cache",
		"no-store"
	],
	"skipSetCookie": true,
	"enableRequestCachePragma": false,
	"allowChunkedEncoding": true,
	"allowPartialContent": false,
	"conds": {
		"isOn": true,
		"connector": "or",
		"groups": [
			{
				"isOn": true,
				"connector": "and",
				"conds": [
					{
						"type": "url-extension",
						"isRequest": true,
						"param": "${requestPathLowerExtension}",
						"operator": "in",
						"value": "[\".css\",\".png\",\".js\",\".woff2\"]",
						"isReverse": false,
						"isCaseInsensitive": false,
						"typeName": "URL扩展名"
					}
				],
				"isReverse": false,
				"description": ""
			}
		]
	},
	"cachePolicy": null,
	"isReverse": false,
	"id": 1
}
~~~