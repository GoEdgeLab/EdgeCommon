# HSTS

## 定义
~~~json
{
  "isOn": "是否启用",
  "maxAge": "最大有效期，单位秒",
  "includeSubDomains": "可选项，是否包含子域名",
  "preload": "可选项，是否预加载",
  "domains": ["可选项，支持的域名1", "可选项，支持的域名2" ...]
}
~~~

其中：
* `maxAge` 可以填写一天（86400秒）或者更长时间
* 如果不填写 `domains` 则支持所有域名

## 示例
### 不限制任何域名
~~~json
{
  "isOn": true,
  "maxAge": 86400,
  "includeSubDomains":false, 
  "preload":false,
  "domains":[]
}
~~~

### 限制域名
~~~json
{
  "isOn": true,
  "maxAge": 86400,
  "includeSubDomains":false, 
  "preload":false,
  "domains":["example.com", "www.example.com"]
}
~~~