package ipconfigs

type IPItemAction = string

const (
	IPItemActionAdd    IPItemAction = "add"
	IPItemActionUpdate IPItemAction = "update"
	IPItemActionDelete IPItemAction = "delete"
)

type IPItemConfig struct {
	Action    IPItemAction `yaml:"action" json:"action"` // 对当前Item的操作
	Id        int64        `yaml:"id" json:"id"`
	IPFrom    string       `yaml:"ipFrom" json:"ipFrom"`
	IPTo      string       `yaml:"ipTo" json:"ipTo"`
	ExpiredAt int64        `yaml:"expiredAt" json:"expiredAt"`
}
