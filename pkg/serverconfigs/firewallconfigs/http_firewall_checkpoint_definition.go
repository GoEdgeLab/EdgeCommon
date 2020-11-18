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
	Name        string            `json:"name"`        // 名称
	Description string            `json:"description"` // 描述
	Prefix      string            `json:"prefix"`      // 前缀
	IsRequest   bool              `json:"isRequest"`   // 是否为请求
	Params      []*KeyValue       `json:"params"`      // 参数
	Options     []OptionInterface `json:"options"`     // 选项
	IsComposed  bool              `json:"isComposed"`  // 是否为组合的checkpoint
}
