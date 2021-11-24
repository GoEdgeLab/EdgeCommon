package serverconfigs

import "encoding/json"

func NewUDPProtocolConfigFromJSON(configJSON []byte) (*UDPProtocolConfig, error) {
	config := &UDPProtocolConfig{}
	if len(configJSON) > 0 {
		err := json.Unmarshal(configJSON, config)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

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

// AsJSON 转换为JSON
func (this *UDPProtocolConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
