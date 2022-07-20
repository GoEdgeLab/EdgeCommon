# HTTP获取客户端IP地址方式配置
## 定义
~~~json
{
  "isPrior": "是否覆盖父级应用",
  "isOn": "是否启用配置",
  "value": "自定义值变量",
  "isCustomized": "是否自定义"
}
~~~

## 示例
### 不启用自定义
~~~json
{
  "isPrior": false,
  "isOn": false,
  "value": "",
  "isCustomized": false
}
~~~

### 启用自定义
~~~json
{
  "isPrior": true,
  "isOn": true,
  "value": "${remoteAddr}",
  "isCustomized": true
}
~~~