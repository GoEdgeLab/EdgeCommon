package shared

type HTTPHeaderPolicyRef struct {
	IsPrior        bool  `yaml:"isPrior" json:"isPrior"`
	IsOn           bool  `yaml:"isOn" json:"isOn"`
	HeaderPolicyId int64 `yaml:"headerPolicyId" json:"headerPolicyId"`
}

func (this *HTTPHeaderPolicyRef) Init() error {
	return nil
}
