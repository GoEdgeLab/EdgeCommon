package serverconfigs

type HTTPLocationRef struct {
	IsOn       bool               `yaml:"isOn" json:"isOn"`             // 是否启用
	LocationId int64              `yaml:"locationId" json:"locationId"` // 路径ID
	Children   []*HTTPLocationRef `yaml:"children" json:"children"`     // 子路径规则
}
