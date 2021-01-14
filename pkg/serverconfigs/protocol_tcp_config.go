package serverconfigs

import "encoding/json"

func NewTCPProtocolConfigFromJSON(configJSON []byte) (*TCPProtocolConfig, error) {
	config := &TCPProtocolConfig{}
	if len(configJSON) > 0 {
		err := json.Unmarshal(configJSON, config)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

type TCPProtocolConfig struct {
	BaseProtocol `yaml:",inline"`
}

func (this *TCPProtocolConfig) Init() error {
	err := this.InitBase()
	if err != nil {
		return err
	}

	return nil
}

// 转换为JSON
func (this *TCPProtocolConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
