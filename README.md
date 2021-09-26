目录结构：
~~~
pkg/
  dnsconfigs - 域名解析和NameServer相关配置
  messageconfigs - 消息通知相关配置
  monitorconfigs - 监控相关配置
  nodeconfigs - 边缘节点相关配置
  serverconfigs - 网站服务相关配置
  systemconfigs - 系统全局配置
  reporterconfigs - 区域监控终端配置
  
  configutils/ - 配置公共函数等
  errors/  - 错误处理
  rpc/  - RPC通讯
     protos/   RPC数据和接口定义
        sevice_*.proto  RPC接口定义
        models/
           model_*.proto RPC数据定义
~~~

开发时需要将 `rpc/protos/` 和 `rpc/protos/models/` 两个目录放入到Proto Buffer检查工具可以找到的位置。