package firewallconfigs

import (
	"reflect"
)

var AllActions = []*HTTPFirewallActionDefinition{
	{
		Name: "阻止",
		Code: HTTPFirewallActionBlock,
	},
	{
		Name: "允许通过",
		Code: HTTPFirewallActionAllow,
	},
	{
		Name: "允许并记录日志",
		Code: HTTPFirewallActionLog,
	},
	{
		Name: "Captcha验证码",
		Code: HTTPFirewallActionCaptcha,
	},
	{
		Name: "跳到下一个规则分组",
		Code: HTTPFirewallActionGoGroup,
		Type: reflect.TypeOf(new(HTTPFirewallGoGroupAction)).Elem(),
	},
	{
		Name: "跳到下一个规则集",
		Code: HTTPFirewallActionGoSet,
		Type: reflect.TypeOf(new(HTTPFirewallGoSetAction)).Elem(),
	},
}

func FindActionName(action HTTPFirewallActionString) string {
	for _, def := range AllActions {
		if def.Code == action {
			return def.Name
		}
	}
	return ""
}
