package serverconfigs

import "encoding/json"

func NewHTTPProtocolConfigFromJSON(configJSON []byte) (*HTTPProtocolConfig, error) {
	config := &HTTPProtocolConfig{}
	if len(configJSON) > 0 {
		err := json.Unmarshal(configJSON, config)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

type HTTPProtocolConfig struct {
	BaseProtocol `yaml:",inline"`
}

func (this *HTTPProtocolConfig) Init() error {
	err := this.InitBase()
	if err != nil {
		return err
	}

	return nil
}

// 转换为JSON
func (this *HTTPProtocolConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
