package sslconfigs

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/lists"
	"reflect"
	"strconv"
	"time"
)

// SSLCertConfig SSL证书
type SSLCertConfig struct {
	Id          int64  `yaml:"id" json:"id"`
	IsOn        bool   `yaml:"isOn" json:"isOn"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"` // 说明
	CertData    []byte `yaml:"certData" json:"certData"`       // 证书数据
	KeyData     []byte `yaml:"keyData" json:"keyData"`         // 密钥数据
	ServerName  string `yaml:"serverName" json:"serverName"`   // 证书使用的主机名，在请求TLS服务器时需要
	IsCA        bool   `yaml:"isCA" json:"isCA"`               // 是否为CA证书
	IsACME      bool   `yaml:"isACME" json:"isACME"`           // 是否通过ACME协议免费申请

	// 以下是从证书中分析所得
	TimeBeginAt int64    `yaml:"timeBeginAt" json:"timeBeginAt"`
	TimeEndAt   int64    `yaml:"timeEndAt" json:"timeEndAt"`
	DNSNames    []string `yaml:"dnsNames" json:"dnsNames"`
	CommonNames []string `yaml:"commonNames" json:"commonNames"`

	// OCSP
	OCSP          []byte `yaml:"ocsp" json:"ocsp"`
	OCSPExpiresAt int64  `yaml:"ocspExpiresAt" json:"ocspExpiresAt"`
	OCSPError     string `yaml:"ocspError" json:"ocspError"`

	cert      *tls.Certificate
	caCerts   []*x509.Certificate
	timeBegin time.Time
	timeEnd   time.Time
}

// Init 校验
func (this *SSLCertConfig) Init(ctx context.Context) error {
	// 如果没有指定数据， 则从ctx中读取数据
	if ctx != nil && len(this.CertData) < 128 {
		var dataMapOne = ctx.Value("DataMap")
		if dataMapOne != nil && !reflect.ValueOf(dataMapOne).IsNil() {
			dataMap, ok := dataMapOne.(*shared.DataMap)
			if !ok {
				return errors.New("SSLCertConfig.init(): invalid 'DataMap' in context")
			}
			if dataMap != nil { // 再次检查是否为nil
				this.KeyData = dataMap.Read(this.KeyData)
				this.CertData = dataMap.Read(this.CertData)
				this.OCSP = dataMap.Read(this.OCSP)
			}
		}
	}

	var commonNames []string // 发行组织
	var dnsNames []string    // 域名

	this.caCerts = []*x509.Certificate{}

	// 分析证书
	if this.IsCA { // CA证书
		var data = this.CertData

		var index = -1
		this.cert = &tls.Certificate{
			Certificate: [][]byte{},
		}
		for {
			index++

			block, rest := pem.Decode(data)
			if block == nil {
				break
			}
			this.cert.Certificate = append(this.cert.Certificate, block.Bytes)
			data = rest
			c, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return err
			}
			if c == nil {
				return errors.New("no available certificates in file")
			}
			this.caCerts = append(this.caCerts, c)

			for _, dnsName := range c.DNSNames {
				if !lists.ContainsString(dnsNames, dnsName) {
					dnsNames = append(dnsNames, dnsName)
				}
			}

			commonNames = append(commonNames, c.Issuer.CommonName)

			if index == 0 {
				this.timeBegin = c.NotBefore
				this.timeEnd = c.NotAfter
			}

			if len(rest) == 0 {
				break
			}
		}
	} else { // 证书+私钥
		cert, err := tls.X509KeyPair(this.CertData, this.KeyData)
		if err != nil {
			return fmt.Errorf("load certificate '%s' failed: %w", strconv.FormatInt(this.Id, 10), err)
		}

		for index, data := range cert.Certificate {
			c, err := x509.ParseCertificate(data)
			if err != nil {
				continue
			}

			if cert.Leaf == nil {
				cert.Leaf = c
			}

			for _, dnsName := range c.DNSNames {
				if !lists.ContainsString(dnsNames, dnsName) {
					dnsNames = append(dnsNames, dnsName)
				}
			}

			commonNames = append(commonNames, c.Issuer.CommonName)

			if index == 0 {
				this.timeBegin = c.NotBefore
				this.timeEnd = c.NotAfter
			}
		}

		this.cert = &cert
	}

	// 赋值分析结果
	this.DNSNames = dnsNames
	this.CommonNames = commonNames
	this.TimeBeginAt = this.timeBegin.Unix()
	this.TimeEndAt = this.timeEnd.Unix()

	return nil
}

// MatchDomain 校验是否匹配某个域名
func (this *SSLCertConfig) MatchDomain(domain string) bool {
	if len(this.DNSNames) == 0 {
		return false
	}
	return configutils.MatchDomains(this.DNSNames, domain)
}

// CertObject 获取证书对象
func (this *SSLCertConfig) CertObject() *tls.Certificate {
	return this.cert
}

func (this *SSLCertConfig) CACerts() []*x509.Certificate {
	return this.caCerts
}

// TimeBegin 开始时间
func (this *SSLCertConfig) TimeBegin() time.Time {
	return this.timeBegin
}

// TimeEnd 结束时间
func (this *SSLCertConfig) TimeEnd() time.Time {
	return this.timeEnd
}
