package firewallconfigs

import "github.com/iwind/TeaGo/maps"

type OptionsOption struct {
	Type       string                                       `json:"type"`
	Name       string                                       `json:"name"`
	Code       string                                       `json:"code"`
	Value      string                                       `json:"value"` // default value
	IsRequired bool                                         `json:"isRequired"`
	Size       int                                          `json:"size"`
	Comment    string                                       `json:"comment"`
	RightLabel string                                       `json:"rightLabel"`
	Validate   func(value string) (ok bool, message string) `json:"-"`
	Options    []maps.Map                                   `json:"options"`
}

func NewOptionsOption(name string, code string) *OptionsOption {
	return &OptionsOption{
		Type: "options",
		Name: name,
		Code: code,
	}
}

func (this *OptionsOption) SetOptions(options []maps.Map) {
	this.Options = options
}
