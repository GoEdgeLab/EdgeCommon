# 容量
## 定义
~~~json
{
  "count": "数量",
  "unit": "单位"
}
~~~

其中：
* `数量` - 必须是一个整数数字
* `单位` - 有以下几个值：
  * `byte` - 字节
  * `kb` - KB
  * `mb` - MB
  * `gb` - GB
  * `tb` - TB
  * `pb` - PB
  * `eb` - EB

## 示例
100MB：
~~~
{
  "count": 100,
  "unit": "mb"
}
~~~


32GB：
~~~
{
  "count": 32,
  "unit": "gb"
}
~~~