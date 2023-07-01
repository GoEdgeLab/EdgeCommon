# 网络地址定义

## 定义
~~~json
{
	"protocol": "协议",
	"host": "主机地址，通常为空",
	"portRange": "端口或者端口范围"
}
~~~

## 示例
对于 `http://example.com`：
~~~json
{
  "protocol": "http",
  "host": "example.com",
  "portRange": "80"
}
~~~


对于 `https://example.com`：
~~~json
{
  "protocol": "https",
  "host": "example.com",
  "portRange": "443"
}
~~~
