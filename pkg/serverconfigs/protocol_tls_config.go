package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"

// TLS协议配置
type TLSProtocolConfig struct {
	BaseProtocol `yaml:",inline"`

	SSLPolicyRef *sslconfigs.SSLRef    `yaml:"sslPolicyRef" json:"sslPolicyRef"`
	SSLPolicy    *sslconfigs.SSLPolicy `yaml:"sslPolicy" json:"sslPolicy"`
}

// 初始化
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
