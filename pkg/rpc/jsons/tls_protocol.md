# TLS协议配置

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
	],
  	"sslPolicyRef": {
	  "isOn": "启用SSL策略",
	  "sslPolicyId": "SSL策略ID"
	}
}
~~~

其中 `SSL策略ID` 通过 `/SSLPolicyService/createSSLPolicy` 接口创建。

## 示例

### 监听8443端口
~~~json
{
  "isOn": true,
  "listen": [
    {
      "protocol": "tls",
      "host": "",
      "portRange": "8443"
	}
  ],
  "sslPolicyRef": {
    "isOn": true,
    "sslPolicyId": 123
  }
}
~~~

其中SSL策略ID `123` 通过 `/SSLPolicyService/createSSLPolicy` 接口创建。

### 监听8443和8543端口
~~~json
{
  "isOn": true,
  "listen": [
    {
      "protocol": "tls",
      "portRange": "8443"
	},
    {
      "protocol": "tls",
      "portRange": "8543"
	}
  ],
  "sslPolicyRef": {
    "isOn": true,
    "sslPolicyId": 123
  }
}
~~~

其中SSL策略ID `123` 通过 `/SSLPolicyService/createSSLPolicy` 接口创建。

