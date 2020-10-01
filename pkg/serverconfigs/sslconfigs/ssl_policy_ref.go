package sslconfigs

type SSLPolicyRef struct {
	IsOn        bool  `yaml:"isOn" json:"isOn"`
	SSLPolicyId int64 `yaml:"sslPolicyId" json:"sslPolicyId"`
}
