# HTTP防火墙（即WAF）引用
## 定义
~~~json
{
  "isPrior": "是否覆盖上级配置",
  "isOn": "是否启用配置",
  "firewallPolicyId": "WAF策略ID",
  "ignoreGlobalRules": "是否忽略系统定义的全局规则",
  "defaultCaptchaType": "默认人机识别方式，可以选none（不设置）、default（默认）、oneClick（单击验证）、slide（滑动解锁）、geetest（极验）"
}
~~~

## 示例
~~~json
{
  "isPrior": true,
  "isOn": true,
  "firewallPolicyId": 123,
  "ignoreGlobalRules": false,
  "defaultCaptchaType": "none"
}
~~~