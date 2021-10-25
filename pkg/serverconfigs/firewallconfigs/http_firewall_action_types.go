package firewallconfigs

type HTTPFirewallActionString = string

const (
	HTTPFirewallActionLog      HTTPFirewallActionString = "log"       // allow and log
	HTTPFirewallActionBlock    HTTPFirewallActionString = "block"     // block
	HTTPFirewallActionCaptcha  HTTPFirewallActionString = "captcha"   // block and show captcha
	HTTPFirewallActionNotify   HTTPFirewallActionString = "notify"    // 告警
	HTTPFirewallActionGet302   HTTPFirewallActionString = "get_302"   // 针对GET的302重定向认证
	HTTPFirewallActionPost307  HTTPFirewallActionString = "post_307"  // 针对POST的307重定向认证
	HTTPFirewallActionRecordIP HTTPFirewallActionString = "record_ip" // 记录IP
	HTTPFirewallActionTag      HTTPFirewallActionString = "tag"       // 标签
	HTTPFirewallActionPage     HTTPFirewallActionString = "page"      // 显示页面
	HTTPFirewallActionAllow    HTTPFirewallActionString = "allow"     // allow
	HTTPFirewallActionGoGroup  HTTPFirewallActionString = "go_group"  // go to next rule group
	HTTPFirewallActionGoSet    HTTPFirewallActionString = "go_set"    // go to next rule set
)
