package serverconfigs

type HTTPFilterRef struct {
	IsPrior        bool  `yaml:"isPrior" json:"isPrior"`
	IsOn           bool  `yaml:"isOn" json:"isOn"`
	FilterPolicyId int64 `yaml:"filterPolicyId" json:"filterPolicyId"`
}

func (this *HTTPFilterRef) Init() error {
	return nil
}
