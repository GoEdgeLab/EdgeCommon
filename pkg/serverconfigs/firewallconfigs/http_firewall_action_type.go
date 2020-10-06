package firewallconfigs

type HTTPFirewallActionString = string

const (
	HTTPFirewallActionLog     = "log"      // allow and log
	HTTPFirewallActionBlock   = "block"    // block
	HTTPFirewallActionCaptcha = "captcha"  // block and show captcha
	HTTPFirewallActionAllow   = "allow"    // allow
	HTTPFirewallActionGoGroup = "go_group" // go to next rule group
	HTTPFirewallActionGoSet   = "go_set"   // go to next rule set
)
