package serverconfigs

type HTTPWebsocketRef struct {
	IsPrior     bool  `yaml:"isPrior" json:"isPrior"`
	IsOn        bool  `yaml:"isOn" json:"isOn"`
	WebsocketId int64 `yaml:"websocketId" json:"websocketId"`
}
