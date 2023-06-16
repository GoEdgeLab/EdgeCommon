# SSL证书引用

可以用来引用一组证书。

## 定义
~~~json
[
  {
	"isOn": "是否启用",
	"certId": "证书ID 1"
  },
  {
    "isOn": "是否启用",
    "certId": "证书ID 2"
  },
  ...
]
~~~

## 示例
~~~json
[
  {
    "isOn": true,
    "certId": 12345
  },
  {
    "isOn": true,
    "certId": 12346
  }
]
~~~

其中：
* `certId` - 证书的ID