# 反向代理调度
## 定义
~~~json
{
	"code": "调度方法代号",
    "options": "调度选项"
}
~~~

其中：
* `code` 调度方法代号
  * `random` - 随机
  * `roundRobin` - 轮询
  * `hash` - Hash算法
    * `key` - 自定义Key，可以使用请求变量，比如 `${remoteAddr}`
  * `sticky` - Sticky算法
    * `type` - 类型：cookie、header、argument
    * `param` - 参数值


## 示例
~~~json
{
  "code": "random",
  "options": null
}
~~~