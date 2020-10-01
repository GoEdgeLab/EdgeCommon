package serverconfigs

type FTPServerRef struct {
	IsOn        bool  `yaml:"isOn" json:"isOn"`
	FTPServerId int64 `yaml:"ftpServerId" json:"ftpServerId"`
}
