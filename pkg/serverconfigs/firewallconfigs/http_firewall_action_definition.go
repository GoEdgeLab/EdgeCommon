package firewallconfigs

import "reflect"

type HTTPFirewallActionCategory = string

const (
	HTTPFirewallActionCategoryBlock  HTTPFirewallActionCategory = "block"
	HTTPFirewallActionCategoryAllow  HTTPFirewallActionCategory = "allow"
	HTTPFirewallActionCategoryVerify HTTPFirewallActionCategory = "verify"
)

// HTTPFirewallActionDefinition action definition
type HTTPFirewallActionDefinition struct {
	Name        string                     `json:"name"`
	Code        HTTPFirewallActionString   `json:"code"`
	Description string                     `json:"description"`
	Type        reflect.Type               `json:"type"`
	Category    HTTPFirewallActionCategory `json:"category"`
}
