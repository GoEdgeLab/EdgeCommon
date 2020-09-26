package shared

// Header引用
type HTTPHeaderRef struct {
	IsOn     bool  `yaml:"isOn" json:"isOn"`
	HeaderId int64 `yaml:"headerId" json:"headerId"`
}
