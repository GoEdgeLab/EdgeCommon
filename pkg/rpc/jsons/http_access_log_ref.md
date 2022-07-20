# 访问日志引用
## 定义
~~~json
{
  "isPrior": "是否覆盖父级应用",
  "isOn": "是否启用配置",
  "fields": ["字段1", "字段2", ...] // 可以留空
  "status1": "是否启用状态1xx",
  "status2": "是否启用状态2xx",
  "status3": "是否启用状态3xx",
  "status4": "是否启用状态4xx",
  "status5": "是否启用状态5xx",
  "enableClientClosed": "是否记录客户端关闭事件",
  "firewallOnly": "是否只记录防火墙（WAF）相关日志"
}
~~~

### 字段值
* `1` - 请求Header
* `2` - 响应Header
* `3` - 请求URL参数
* `4` - Cookie
* `5` - 扩展信息
* `6` - Referer
* `7` - UserAgent
* `8` - 请求Body
* `9` - 响应Body（目前不支持）

## 示例
~~~json
{
  "isPrior": true,
  "isOn": true,
  "fields": [],
  "status1": true,
  "status2": true,
  "status3": true,
  "status4": true,
  "status5": true,
  "enableClientClosed": true,
  "firewallOnly": true
}
~~~