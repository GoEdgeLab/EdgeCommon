package serverconfigs

import "encoding/json"

type UnixProtocolConfig struct {
	BaseProtocol `yaml:",inline"`
}

func (this *UnixProtocolConfig) Init() error {
	err := this.InitBase()
	if err != nil {
		return err
	}

	return nil
}

// 转换为JSON
func (this *UnixProtocolConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
