package firewallconfigs

type HTTPFirewallRule struct {
	Id                int64                  `yaml:"id" json:"id"`
	IsOn              bool                   `yaml:"isOn" json:"isOn"`
	Param             string                 `yaml:"param" json:"param"`
	Operator          string                 `yaml:"operator" json:"operator"`
	Value             string                 `yaml:"value" json:"value"`
	IsCaseInsensitive bool                   `yaml:"isCaseInsensitive" json:"isCaseInsensitive"`
	CheckpointOptions map[string]interface{} `yaml:"checkpointOptions" json:"checkpointOptions"`
	Description       string                 `yaml:"description" json:"description"`
}

func (this *HTTPFirewallRule) Init() error {
	return nil
}
