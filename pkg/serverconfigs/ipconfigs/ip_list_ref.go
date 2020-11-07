package ipconfigs

// IP名单引用
type IPListRef struct {
	IsOn   bool  `yaml:"isOn" json:"isOn"`
	ListId int64 `yaml:"listId" json:"listId"`
}
