# WebSocket引用

## 定义
~~~json
{
	"isPrior": "是否覆盖上级配置，true|false",
	"isOn": "是否启用，true|false",
	"websocketId": "Websocket配置ID"
}
~~~
其中：
* `Websocket配置ID` - 需要调用 `HTTPWebsocketService.CreateHTTPWebsocketRequest()` 生成 

## 示例
~~~json
{
	"isPrior": true,
	"isOn": true,
	"websocketId": 123
}
~~~