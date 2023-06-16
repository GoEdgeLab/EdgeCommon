# HTTP协议配置

## 定义
~~~json
{
	"isOn": "是否启用",
  	"listen": [
      {
		"protocol": "协议",
		"host": "主机地址，通常为空",
		"portRange": "端口或者端口范围"
	  },
	  ...
	]
}
~~~

## 示例

### 监听80端口
~~~json
{
  "isOn": true,
  "listen": [
    {
      "protocol": "http",
      "host": "",
      "portRange": "80"
	}
  ]
}
~~~

### 监听80和8080端口
~~~json
{
  "isOn": true,
  "listen": [
    {
      "protocol": "http",
      "portRange": "80"
	},
    {
      "protocol": "http",
      "portRange": "8080"
	}
  ]
}
~~~