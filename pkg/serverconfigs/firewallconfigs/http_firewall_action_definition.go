package firewallconfigs

import "reflect"

// action definition
type HTTPFirewallActionDefinition struct {
	Name        string
	Code        HTTPFirewallActionString
	Description string
	Type        reflect.Type
}
