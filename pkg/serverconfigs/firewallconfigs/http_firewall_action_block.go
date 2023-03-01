package firewallconfigs

import "net/http"

// HTTPFirewallBlockAction url client configure
type HTTPFirewallBlockAction struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	StatusCode int           `yaml:"statusCode" json:"statusCode"`
	Body       string        `yaml:"body" json:"body"` // supports HTML
	URL        string        `yaml:"url" json:"url"`
	Timeout    int32         `yaml:"timeout" json:"timeout"`       // 最小封禁时长
	TimeoutMax int32         `yaml:"timeoutMax" json:"timeoutMax"` // 最大封禁时长
	Scope      FirewallScope `yaml:"scope" json:"scope"`
}

func DefaultHTTPFirewallBlockAction() *HTTPFirewallBlockAction {
	return &HTTPFirewallBlockAction{
		StatusCode: http.StatusForbidden,
		Body:       "Blocked By WAF",
		Timeout:    300,
	}
}
