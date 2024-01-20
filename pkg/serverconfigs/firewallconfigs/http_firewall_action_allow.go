package firewallconfigs

type AllowScope = string

const (
	AllowScopeGroup  AllowScope = "group"
	AllowScopeServer AllowScope = "server"
	AllowScopeGlobal AllowScope = "global"
)

type HTTPFirewallAllowAction struct {
	Scope AllowScope `yaml:"scope" json:"scope"`
}
