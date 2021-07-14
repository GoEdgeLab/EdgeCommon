package firewallconfigs

import (
	"reflect"
)

var AllActions = []*HTTPFirewallActionDefinition{
	{
		Name:        "阻止",
		Code:        HTTPFirewallActionBlock,
		Description: "阻止请求并中断当前连接。",
		Category:    HTTPFirewallActionCategoryBlock,
	},
	{
		Name:        "允许通过",
		Code:        HTTPFirewallActionAllow,
		Description: "允许正常通过，不记录到日志。",
		Category:    HTTPFirewallActionCategoryAllow,
	},
	{
		Name:        "允许并记录日志",
		Code:        HTTPFirewallActionLog,
		Description: "允许正常通过并记录到日志。",
		Category:    HTTPFirewallActionCategoryAllow,
	},
	{
		Name:        "Captcha验证码",
		Code:        HTTPFirewallActionCaptcha,
		Description: "在浏览器使用验证码来验证客户端。",
		Category:    HTTPFirewallActionCategoryVerify,
	},
	{
		Name:        "告警",
		Code:        HTTPFirewallActionNotify,
		Description: "发送消息通知。",
		Category:    HTTPFirewallActionCategoryVerify,
	},
	{
		Name:        "GET 302",
		Code:        HTTPFirewallActionGet302,
		Description: "通过302重定向GET请求验证客户端。",
		Category:    HTTPFirewallActionCategoryVerify,
	},
	{
		Name:        "POST 307",
		Code:        HTTPFirewallActionPost307,
		Description: "通过307重定向POST请求验证客户端。",
		Category:    HTTPFirewallActionCategoryVerify,
	},
	{
		Name:        "记录IP",
		Code:        HTTPFirewallActionRecordIP,
		Description: "将此IP记录到某个IP名单中。",
		Category:    HTTPFirewallActionCategoryBlock,
	},
	{
		Name:        "标签",
		Code:        HTTPFirewallActionTag,
		Description: "为匹配的请求打上标签。",
		Category:    HTTPFirewallActionCategoryAllow,
	},
	{
		Name:     "跳到下一个规则分组",
		Code:     HTTPFirewallActionGoGroup,
		Type:     reflect.TypeOf(new(HTTPFirewallGoGroupAction)).Elem(),
		Category: HTTPFirewallActionCategoryVerify,
	},
	{
		Name:     "跳到下一个规则集",
		Code:     HTTPFirewallActionGoSet,
		Type:     reflect.TypeOf(new(HTTPFirewallGoSetAction)).Elem(),
		Category: HTTPFirewallActionCategoryVerify,
	},
}

func FindActionDefinition(actionCode HTTPFirewallActionString) *HTTPFirewallActionDefinition {
	for _, def := range AllActions {
		if def.Code == actionCode {
			return def
		}
	}
	return nil
}
