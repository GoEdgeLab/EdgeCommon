package firewallconfigs

type KeyValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewKeyValue(name string, value string) *KeyValue {
	return &KeyValue{
		Name:  name,
		Value: value,
	}
}

// check point definition
type HTTPFirewallCheckpointDefinition struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Prefix      string            `json:"prefix"`
	IsRequest   bool              `json:"isRequest"`
	Params      []*KeyValue       `json:"params"`
	Options     []OptionInterface `json:"options"`
}
