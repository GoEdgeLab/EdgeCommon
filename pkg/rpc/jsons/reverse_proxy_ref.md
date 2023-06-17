# 反向代理引用
## 定义
~~~json
{
  "isOn": "是否启用",
  "isPrior": "是否覆盖上级配置，用于路由规则",
  "reverseProxyId": "反向代理ID"
}
~~~
其中：
* `reverseProxyId` - 反向代理ID，可以通过 `/ReverseProxyService/createReverseProxy` 创建

## 示例
~~~json
{
  "isOn": true,
  "reverseProxyId": 123
}
~~~