package firewallconfigs

// HTTPFirewallBlockAction url client configure
type HTTPFirewallBlockAction struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	StatusCode int           `yaml:"statusCode" json:"statusCode"`
	Body       string        `yaml:"body" json:"body"` // supports HTML
	URL        string        `yaml:"url" json:"url"`
	Timeout    int32         `yaml:"timeout" json:"timeout"`
	Scope      FirewallScope `yaml:"scope" json:"scope"`
}
