package serverconfigs

import "encoding/json"

type UDPProtocolConfig struct {
	BaseProtocol `yaml:",inline"`
}

func (this *UDPProtocolConfig) Init() error {
	err := this.InitBase()
	if err != nil {
		return err
	}

	return nil
}

// 转换为JSON
func (this *UDPProtocolConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
