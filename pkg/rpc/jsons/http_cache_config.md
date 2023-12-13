# HTTP缓存配置
## 定义
~~~json
{
  "isPrior": "是否覆盖上级配置",
  "isOn": "是否启用配置",
  "addStatusHeader": "是否增加命中状态Header（X-Cache）",
  "addAgeHeader": "是否增加Age Header",
  "enableCacheControlMaxAge": "是否支持Cache-Control: max-age=...",
  "disablePolicyRefs": "是否停用策略中定义的条件",
  "purgeIsOn": "是否允许使用Purge方法清理",
  "purgeKey": "Purge时使用的X-Edge-Purge-Key",
  "stale": "陈旧缓存使用策略",
  "key": "主域名配置",
  "cacheRefs": ["缓存条件1", "缓存条件2", ...]
}
~~~
其中：
* `缓存条件` - 参考 {json:http_cache_ref}
* `主域名配置` 参考本文“主域名”配置部分

## 示例
### 无缓存条件
~~~json
{
  "isPrior": true,
  "isOn": true,
  "addStatusHeader": true,
  "addAgeHeader": true,
  "enableCacheControlMaxAge": true,
  "disablePolicyRefs": false,
  "purgeIsOn": false,
  "purgeKey": "",
  "stale": null,
  "cacheRefs": []
}
~~~

### 加入缓存条件
~~~json
{
  "isPrior": true,
  "isOn": true,
  "addStatusHeader": true,
  "addAgeHeader": true,
  "enableCacheControlMaxAge": true,
  "disablePolicyRefs": false,
  "purgeIsOn": false,
  "purgeKey": "",
  "stale": null,
  "cacheRefs": [
    {
	  "id": 0,
	  "isOn": true,
	  "key": "${scheme}://${host}${requestPath}${isArgs}${args}",
	  "life": {"count": 2, "unit": "hour"},
	  "status": [200],
      "maxSize": {"count": 32, "unit": "mb"},
      "minSize": {"count": 0, "unit": "kb"},
      "skipCacheControlValues": ["private", "no-cache", "no-store"],
      "skipSetCookie": true,
      "enableRequestCachePragma": false,
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
      "allowChunkedEncoding": true,
      "allowPartialContent": false,
      "isReverse": false,
      "methods": []
	}
  ]
}
~~~


## 主域名配置
~~~json
{
  "isOn": "true|false",
  "scheme": "https|http",
  "host": "域名，必须是当前网站已绑定的域名"
}
~~~

### 示例
#### 不使用主域名
~~~json
{
  "isOn": false
}
~~~

#### 使用主域名
~~~json
{
  "isOn": true,
  "scheme": "https",
  "host": "example.com"
}
~~~

如果启用主域名，则缓存键值中的域名会被自动换成主域名，清理缓存的时候也需要使用此主域名。