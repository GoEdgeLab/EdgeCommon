# 域名信息列表

## 定义
~~~
[ 域名信息1, 域名信息2, ... ]
~~~
其中 `域名信息N` 等是单个域名信息定义，具体请参考 {json:server_name}

## 示例
### 示例1：单个域名
~~~json
[
	{
		"name": "example.com",
		"type": "full"
	}
]
~~~

### 示例2：多个域名
~~~json
[
	{
		"name": "example.com",
		"type": "full"
	},
	{
      	"name": "google.com",
      	"type": "full"
  	},
  	{
		"name": "facebook.com",
	  	"type": "full"
  	}
]
~~~

### 示例3：域名合集
域名合集效果跟多个域名是一样的，只不过在界面上以一个目录的形式呈现。
~~~json
[
  	{
		"name": "",
	  	"type": "full",
	  	"subNames": ["example.com", "google.com", "facebook.com"]
  	}
]
~~~
