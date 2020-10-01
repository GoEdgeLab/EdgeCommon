package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"

// HTTPS协议配置
type HTTPSProtocolConfig struct {
	BaseProtocol `yaml:",inline"`

	SSLPolicyRef *sslconfigs.SSLPolicyRef `yaml:"sslPolicyRef" json:"sslPolicyRef"`
	SSLPolicy    *sslconfigs.SSLPolicy    `yaml:"sslPolicy" json:"sslPolicy"`
}

// 初始化
func (this *HTTPSProtocolConfig) Init() error {
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
