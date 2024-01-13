# 比特位尺寸
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
  * `b` - 比特
  * `kb` - Kb
  * `mb` - Mb
  * `gb` - Gb
  * `tb` - Tb
  * `pb` - Pb
  * `eb` - Eb

## 示例
100Mb：
~~~
{
  "count": 100,
  "unit": "mb"
}
~~~


32Gb：
~~~
{
  "count": 32,
  "unit": "gb"
}
~~~