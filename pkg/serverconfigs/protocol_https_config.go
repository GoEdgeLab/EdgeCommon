package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"

// TLS Version
type TLSVersion = string

// Cipher Suites
type TLSCipherSuite = string

type HTTPSProtocolConfig struct {
	BaseProtocol `yaml:",inline"`

	SSL *sslconfigs.SSLConfig `yaml:"ssl"`
}

func (this *HTTPSProtocolConfig) Init() error {
	err := this.InitBase()
	if err != nil {
		return err
	}

	return nil
}
