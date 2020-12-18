package serverconfigs

import "encoding/json"

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
