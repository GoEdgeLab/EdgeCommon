package sslconfigs

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"golang.org/x/net/http2"
	"time"
)

// TLSVersion TLS Version
type TLSVersion = string

// TLSCipherSuite Cipher Suites
type TLSCipherSuite = string

// SSLPolicy SSL配置
type SSLPolicy struct {
	Id   int64 `yaml:"id" json:"id"`     // ID
	IsOn bool  `yaml:"isOn" json:"isOn"` // 是否开启

	CertRefs       []*SSLCertRef     `yaml:"certRefs" json:"certRefs"`
	Certs          []*SSLCertConfig  `yaml:"certs" json:"certs"`
	ClientAuthType SSLClientAuthType `yaml:"clientAuthType" json:"clientAuthType"` // 客户端认证类型
	ClientCARefs   []*SSLCertRef     `yaml:"clientCARefs" json:"clientCARefs"`     // 客户端认证CA证书引用
	ClientCACerts  []*SSLCertConfig  `yaml:"clientCACerts" json:"clientCACerts"`   // 客户端认证CA

	MinVersion       TLSVersion       `yaml:"minVersion" json:"minVersion"`             // 支持的最小版本
	CipherSuitesIsOn bool             `yaml:"cipherSuitesIsOn" json:"cipherSuitesIsOn"` // 是否自定义加密算法套件
	CipherSuites     []TLSCipherSuite `yaml:"cipherSuites" json:"cipherSuites"`         // 加密算法套件

	HSTS         *HSTSConfig `yaml:"hsts" json:"hsts"`                 // HSTS配置
	HTTP2Enabled bool        `yaml:"http2Enabled" json:"http2Enabled"` // 是否启用HTTP2

	OCSPIsOn bool `yaml:"ocspIsOn" json:"ocspIsOn"` // 是否启用OCSP

	nameMapping map[string]*tls.Certificate // dnsName => cert

	minVersion   uint16
	cipherSuites []uint16

	clientCAPool *x509.CertPool

	tlsConfig *tls.Config

	ocspExpiresAt int64 // OCSP最早过期时间
}

// Init 校验配置
func (this *SSLPolicy) Init() error {
	this.nameMapping = map[string]*tls.Certificate{}

	// certs
	var certs = []tls.Certificate{}
	for _, cert := range this.Certs {
		err := cert.Init()
		if err != nil {
			return err
		}
		if this.OCSPIsOn && len(cert.OCSP) > 0 && cert.OCSPExpiresAt > time.Now().Unix() {
			if this.ocspExpiresAt == 0 || cert.OCSPExpiresAt < this.ocspExpiresAt {
				this.ocspExpiresAt = cert.OCSPExpiresAt
			}
			cert.CertObject().OCSPStaple = cert.OCSP
		}
		certs = append(certs, *cert.CertObject())
		for _, dnsName := range cert.DNSNames {
			this.nameMapping[dnsName] = cert.CertObject()
		}
	}

	// CA certs
	for _, cert := range this.ClientCACerts {
		err := cert.Init()
		if err != nil {
			return err
		}
		certs = append(certs, *cert.CertObject())
		for _, dnsName := range cert.DNSNames {
			this.nameMapping[dnsName] = cert.CertObject()
		}
	}

	// min version
	this.convertMinVersion()

	// cipher suite categories
	this.initCipherSuites()

	// hsts
	if this.HSTS != nil {
		err := this.HSTS.Init()
		if err != nil {
			return err
		}
	}

	// tls config
	this.tlsConfig = &tls.Config{}
	cipherSuites := this.TLSCipherSuites()
	if !this.CipherSuitesIsOn || len(cipherSuites) == 0 {
		cipherSuites = nil
	}

	nextProto := []string{}
	if this.HTTP2Enabled {
		nextProto = []string{http2.NextProtoTLS}
	}
	this.tlsConfig = &tls.Config{
		Certificates:   certs,
		MinVersion:     this.TLSMinVersion(),
		CipherSuites:   cipherSuites,
		GetCertificate: nil,
		ClientAuth:     GoSSLClientAuthType(this.ClientAuthType),
		ClientCAs:      this.CAPool(),
		NextProtos:     nextProto,
	}

	return nil
}

// TLSMinVersion 取得最小版本
func (this *SSLPolicy) TLSMinVersion() uint16 {
	return this.minVersion
}

// TLSCipherSuites 套件
func (this *SSLPolicy) TLSCipherSuites() []uint16 {
	return this.cipherSuites
}

// MatchDomain 校验是否匹配某个域名
func (this *SSLPolicy) MatchDomain(domain string) (cert *tls.Certificate, ok bool) {
	cert, ok = this.nameMapping[domain]
	if ok {
		return
	}

	for name, cert := range this.nameMapping {
		if configutils.MatchDomain(name, domain) {
			return cert, true
		}
	}
	return nil, false
}

// FirstCert 取得第一个证书
func (this *SSLPolicy) FirstCert() *tls.Certificate {
	for _, cert := range this.Certs {
		return cert.CertObject()
	}
	return nil
}

// CAPool CA证书Pool，用于TLS对客户端进行认证
func (this *SSLPolicy) CAPool() *x509.CertPool {
	return this.clientCAPool
}

func (this *SSLPolicy) TLSConfig() *tls.Config {
	return this.tlsConfig
}

// ContainsCert 检查是否包括某个证书
func (this *SSLPolicy) ContainsCert(certId int64) bool {
	for _, cert := range this.Certs {
		if cert.Id == certId {
			return true
		}
	}
	return false
}

// UpdateCertOCSP 修改某个证书的OCSP
func (this *SSLPolicy) UpdateCertOCSP(certId int64, ocsp []byte, expiresAt int64) {
	var nowTime = time.Now().Unix()

	for _, cert := range this.Certs {
		if cert.Id == certId {
			cert.OCSP = ocsp
			cert.OCSPExpiresAt = expiresAt
			cert.CertObject().OCSPStaple = cert.OCSP

			// 修改tlsConfig中的cert
			for index, certObj := range this.tlsConfig.Certificates {
				if this.certIsEqual(*cert.CertObject(), certObj) {
					if len(cert.OCSP) > 0 && cert.OCSPExpiresAt > nowTime {
						this.tlsConfig.Certificates[index].OCSPStaple = ocsp

						// 重置过期时间
						if this.ocspExpiresAt == 0 || cert.OCSPExpiresAt < this.ocspExpiresAt {
							this.ocspExpiresAt = cert.OCSPExpiresAt
						}
					} else {
						this.tlsConfig.Certificates[index].OCSPStaple = nil
					}
				}
			}
			break
		}
	}
}

// CheckOCSP 检查OCSP过期时间
func (this *SSLPolicy) CheckOCSP() {
	if !this.OCSPIsOn || this.ocspExpiresAt == 0 {
		return
	}

	var nowTime = time.Now().Unix()
	if this.ocspExpiresAt > nowTime {
		return
	}
	this.ocspExpiresAt = 0

	for _, cert := range this.Certs {
		if cert.OCSPExpiresAt > 0 && cert.OCSPExpiresAt < nowTime+1 {
			// 重置OCSP
			cert.OCSP = nil
			cert.OCSPExpiresAt = 0
			for index, certObj := range this.tlsConfig.Certificates {
				if this.certIsEqual(*cert.CertObject(), certObj) {
					this.tlsConfig.Certificates[index].OCSPStaple = nil
				}
			}
		} else if len(cert.OCSP) > 0 && cert.OCSPExpiresAt > nowTime && (this.ocspExpiresAt == 0 || cert.OCSPExpiresAt < this.ocspExpiresAt) {
			// 重置过期时间
			this.ocspExpiresAt = cert.OCSPExpiresAt
		}
	}
}

// OcspExpiresAt OCSP最近过期时间
func (this *SSLPolicy) OcspExpiresAt() int64 {
	return this.ocspExpiresAt
}

func (this *SSLPolicy) certIsEqual(cert1 tls.Certificate, cert2 tls.Certificate) bool {
	var b1 = cert1.Certificate
	var b2 = cert2.Certificate
	if len(b1) != len(b2) {
		return false
	}

	for index, b := range b1 {
		if bytes.Compare(b, b2[index]) != 0 {
			return false
		}
	}

	return true
}
