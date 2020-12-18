package serverconfigs

import "encoding/json"

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
