package firewallconfigs

// HTTPFirewallBlockAction url client configure
type HTTPFirewallBlockAction struct {
	StatusCode int    `yaml:"statusCode" json:"statusCode"`
	Body       string `yaml:"body" json:"body"` // supports HTML
	URL        string `yaml:"url" json:"url"`
	Timeout    int32  `yaml:"timeout" json:"timeout"`
}
