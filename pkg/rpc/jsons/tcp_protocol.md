# TCP协议配置

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

### 监听1234端口
~~~json
{
  "isOn": true,
  "listen": [
    {
      "protocol": "tcp",
      "host": "",
      "portRange": "1234"
	}
  ]
}
~~~

### 监听1234和2345端口
~~~json
{
  "isOn": true,
  "listen": [
    {
      "protocol": "tcp",
      "portRange": "1234"
	},
    {
      "protocol": "tcp",
      "portRange": "2345"
	}
  ]
}
~~~

### 监听1234到1240之间的所有端口
~~~json
{
  "isOn": true,
  "listen": [
    {
      "protocol": "tcp",
      "host": "",
      "portRange": "1234-1240"
	}
  ]
}
~~~