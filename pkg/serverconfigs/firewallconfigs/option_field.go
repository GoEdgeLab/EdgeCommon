package firewallconfigs

// attach option
type FieldOption struct {
	Type        string                                       `json:"type"`
	Name        string                                       `json:"name"`
	Code        string                                       `json:"code"`
	Value       string                                       `json:"value"` // default value
	IsRequired  bool                                         `json:"isRequired"`
	Size        int                                          `json:"size"`
	Comment     string                                       `json:"comment"`
	Placeholder string                                       `json:"placeholder"`
	RightLabel  string                                       `json:"rightLabel"`
	MaxLength   int                                          `json:"maxLength"`
	Validate    func(value string) (ok bool, message string) `json:"-"`
}

func NewFieldOption(name string, code string) *FieldOption {
	return &FieldOption{
		Type: "field",
		Name: name,
		Code: code,
	}
}
