package serverconfigs

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
)

func NewTLSProtocolConfigFromJSON(configJSON []byte) (*TLSProtocolConfig, error) {
	config := &TLSProtocolConfig{}
	if len(configJSON) > 0 {
		err := json.Unmarshal(configJSON, config)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

// TLSProtocolConfig TLS协议配置
type TLSProtocolConfig struct {
	BaseProtocol `yaml:",inline"`

	SSLPolicyRef *sslconfigs.SSLPolicyRef `yaml:"sslPolicyRef" json:"sslPolicyRef"`
	SSLPolicy    *sslconfigs.SSLPolicy    `yaml:"sslPolicy" json:"sslPolicy"`
}

// Init 初始化
func (this *TLSProtocolConfig) Init() error {
	err := this.InitBase()
	if err != nil {
		return err
	}

	if this.SSLPolicy != nil {
		err := this.SSLPolicy.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// AsJSON 转换为JSON
func (this *TLSProtocolConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
