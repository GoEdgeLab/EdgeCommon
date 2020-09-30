package sslconfigs

type SSLCertRef struct {
	IsOn   bool  `yaml:"isOn" json:"isOn"`
	CertId int64 `yaml:"certId" json:"certId"`
}
