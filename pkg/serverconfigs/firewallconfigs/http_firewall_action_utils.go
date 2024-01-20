package firewallconfigs

import (
	"reflect"
)

var AllActions = []*HTTPFirewallActionDefinition{
	{
		Name:        "显示网页",
		Code:        HTTPFirewallActionPage,
		Description: "显示请求被拦截的网页。",
		Category:    HTTPFirewallActionCategoryBlock,
	},
	{
		Name:        "阻止",
		Code:        HTTPFirewallActionBlock,
		Description: "阻止请求并中断当前连接，并自动将当前客户端IP加入到系统黑名单；使用此动作时，请先自行严格测试设置的规则是否正确，避免因错误封禁而导致用户无法正常访问的严重后果！",
		Category:    HTTPFirewallActionCategoryBlock,
	},
	{
		Name:        "Captcha人机识别",
		Code:        HTTPFirewallActionCaptcha,
		Description: "在浏览器使用人机识别机制（比如验证码）来验证客户端。",
		Category:    HTTPFirewallActionCategoryVerify,
	},
	{
		Name:        "JS Cookie验证",
		Code:        HTTPFirewallActionJavascriptCookie,
		Description: "首次访问网站时通过Javascript设置Cookie来验证请求。",
		Category:    HTTPFirewallActionCategoryVerify,
	},
	{
		Name:        "记录IP",
		Code:        HTTPFirewallActionRecordIP,
		Description: "将此IP记录到某个IP名单中。",
		Category:    HTTPFirewallActionCategoryBlock,
	},
	{
		Name:        "跳转",
		Code:        HTTPFirewallActionRedirect,
		Description: "跳转到新的URL。",
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
		Name:        "标签",
		Code:        HTTPFirewallActionTag,
		Description: "为匹配的请求打上标签。",
		Category:    HTTPFirewallActionCategoryAllow,
	},
	{
		Name:        "告警",
		Code:        HTTPFirewallActionNotify,
		Description: "向集群的消息接收人发送消息通知（商业版）。",
		Category:    HTTPFirewallActionCategoryVerify,
	},
	{
		Name:        "GET 302",
		Code:        HTTPFirewallActionGet302,
		Description: "通过302重定向GET请求验证客户端真实性。",
		Category:    HTTPFirewallActionCategoryVerify,
	},
	{
		Name:        "POST 307",
		Code:        HTTPFirewallActionPost307,
		Description: "通过307重定向POST请求验证客户端真实性。",
		Category:    HTTPFirewallActionCategoryVerify,
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
