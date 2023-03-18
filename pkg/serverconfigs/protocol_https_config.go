package serverconfigs

import (
	"context"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
)

func NewHTTPSProtocolConfigFromJSON(configJSON []byte) (*HTTPSProtocolConfig, error) {
	config := &HTTPSProtocolConfig{}
	if len(configJSON) > 0 {
		err := json.Unmarshal(configJSON, config)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

// HTTPSProtocolConfig HTTPS协议配置
type HTTPSProtocolConfig struct {
	BaseProtocol `yaml:",inline"`

	SSLPolicyRef *sslconfigs.SSLPolicyRef `yaml:"sslPolicyRef" json:"sslPolicyRef"`
	SSLPolicy    *sslconfigs.SSLPolicy    `yaml:"sslPolicy" json:"sslPolicy"`
}

// Init 初始化
func (this *HTTPSProtocolConfig) Init(ctx context.Context) error {
	err := this.InitBase()
	if err != nil {
		return err
	}

	if this.SSLPolicy != nil {
		err := this.SSLPolicy.Init(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// AsJSON 转换为JSON
func (this *HTTPSProtocolConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
