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
  "cacheRefs": ["缓存条件1", "缓存条件2", ...]
}
~~~
其中：
* `缓存条件` - 参考 {json:http_cache_ref}

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
                "param": "${requestPathExtension}",
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
