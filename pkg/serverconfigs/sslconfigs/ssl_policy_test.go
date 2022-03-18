// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package sslconfigs_test

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/iwind/TeaGo/assert"
	"testing"
	"time"
)

func TestSSLPolicy_MatchDomain(t *testing.T) {
	var a = assert.NewAssertion(t)

	var policy = &sslconfigs.SSLPolicy{}
	policy.Certs = []*sslconfigs.SSLCertConfig{
		{
			Id:       1,
			DNSNames: []string{"a.com", "b.com"},
		},
		{
			Id:       2,
			DNSNames: []string{"c.com", "d.com"},
		},
		{
			Id:       3,
			DNSNames: []string{"e.com", "f.com"},
		},
	}

	{
		_, ok := policy.MatchDomain("c.com")
		a.IsTrue(ok)
	}
}

func TestSSLPolicy_CheckOCSP(t *testing.T) {
	var certData = []byte(`-----BEGIN CERTIFICATE-----
MIIEcTCCA9qgAwIBAgIDbhMuMA0GCSqGSIb3DQEBBQUAMIGKMQswCQYDVQQGEwJD
TjESMBAGA1UECBMJR3Vhbmdkb25nMREwDwYDVQQHEwhTaGVuemhlbjEQMA4GA1UE
ChMHVGVuY2VudDEMMAoGA1UECxMDV1hHMRMwEQYDVQQDEwpNbXBheW1jaENBMR8w
HQYJKoZIhvcNAQkBFhBtbXBheW1jaEB0ZW5jZW50MB4XDTE2MTIxMjA5NDAwM1oX
DTI2MTIxMDA5NDAwM1owgaExCzAJBgNVBAYTAkNOMRIwEAYDVQQIEwlHdWFuZ2Rv
bmcxETAPBgNVBAcTCFNoZW56aGVuMRAwDgYDVQQKEwdUZW5jZW50MQ4wDAYDVQQL
EwVNTVBheTE2MDQGA1UEAxQt5YyX5Lqs5LiJ55m+5YWt5Y2B6KGM5LqS6IGU56eR
5oqA5pyJ6ZmQ5YWs5Y+4MREwDwYDVQQEEwgxNzIyNzc0NDCCASIwDQYJKoZIhvcN
AQEBBQADggEPADCCAQoCggEBAN2/1axdFhLKgMAGpkM9kpBfz88IvVYLFLaRrsIO
aM4RLDup5ye0GrOvQtq8gvPFbn+GuekyBfoVRNHW1OSv/uQfDYd5tcmAy/0BDZSL
OfPHaYOS2fj2y9KvLZTFTMBszG9kwV/FFlHgK4SJKbikdqTPd9vnt6Yr7FyfTIws
K9RQ77vetOTduWZttON+RK/Tlz6AepiVfl9LZ/XOVveYI/6TfEbI6uUoeXrlSKCf
w8/yfo69tcZV0g9yjSnVYDvgp6BFXJ1QK1CnJB4Dnol8XoBgUIrUyJqO+LvPr2Qy
wsnyONc15AJK/23vebDGGvTvYtu47qRywISD4ioW15YBK1UCAwEAAaOCAUYwggFC
MAkGA1UdEwQCMAAwLAYJYIZIAYb4QgENBB8WHSJDRVMtQ0EgR2VuZXJhdGUgQ2Vy
dGlmaWNhdGUiMB0GA1UdDgQWBBQVQCAalLY0TuS+z80biOcWb0QkzjCBvwYDVR0j
BIG3MIG0gBQ+BSb2ImK0FVuIzWR+sNRip+WGdKGBkKSBjTCBijELMAkGA1UEBhMC
Q04xEjAQBgNVBAgTCUd1YW5nZG9uZzERMA8GA1UEBxMIU2hlbnpoZW4xEDAOBgNV
BAoTB1RlbmNlbnQxDDAKBgNVBAsTA1dYRzETMBEGA1UEAxMKTW1wYXltY2hDQTEf
MB0GCSqGSIb3DQEJARYQbW1wYXltY2hAdGVuY2VudIIJALtUlyu8AOhXMA4GA1Ud
DwEB/wQEAwIGwDAWBgNVHSUBAf8EDDAKBggrBgEFBQcDAjANBgkqhkiG9w0BAQUF
AAOBgQA/Zr9PIRE8c3mAnb0lmx/DToFJrUB4Sr51szjiX5XiKymBoC2hnwJvI+7B
EkRdNv4S7rvu33GS7BcZvjEwyrZdA9ZRIQz1MiaBIXdayIkkUCxaStB1junI8Jfc
dG6S+JIMJU8y0tG53vEG2JRw8Mmm1qloAxs1Zl92UtlZoiHHCQ==
-----END CERTIFICATE-----
`)
	var keyData = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDdv9WsXRYSyoDA
BqZDPZKQX8/PCL1WCxS2ka7CDmjOESw7qecntBqzr0LavILzxW5/hrnpMgX6FUTR
1tTkr/7kHw2HebXJgMv9AQ2Uiznzx2mDktn49svSry2UxUzAbMxvZMFfxRZR4CuE
iSm4pHakz3fb57emK+xcn0yMLCvUUO+73rTk3blmbbTjfkSv05c+gHqYlX5fS2f1
zlb3mCP+k3xGyOrlKHl65Uign8PP8n6OvbXGVdIPco0p1WA74KegRVydUCtQpyQe
A56JfF6AYFCK1Miajvi7z69kMsLJ8jjXNeQCSv9t73mwxhr072LbuO6kcsCEg+Iq
FteWAStVAgMBAAECggEAVkqTfMqQj2lsJs2vn5TzVulh9cAB5dzUB6OzbOKsmBwI
qYMZZ9LnXSsDihk3oGMg99FWwU9tEf9602mVWRS/zMfkvOZ4/lv3hZIGVdrEB4B/
J+talU58zJTM2QraLjtoZqS/t2P7porkhGPX73lYjhQKIXIPfkOza+u1nwqFV84Z
YowiRTowuBHAAduW7W5uv4MaGG6P9w/JzR4zHCUjc5rnh/3a3+TN6KRxkXDh+1yi
6wg0S54qtTyAEIeGMjIqhzUgN0fxlhyMgtROi8h3DN/tvBoOCT9jGFeTsBcfk6Ib
p4sMDo/OcC1NXsENsccVprH297jKmwV0vZFGUebPAQKBgQD9bfrWU0TvLlLILJmT
52HRy6HCddKV6SdCBF04Rz3a5L32epKREql6l8KewHo05wlty90UL4sltHwZo9h6
QuukNMMuLvaye2qOAkuFw1x5qD4R2VvbQsPDHoJt0zOzzF77/Faob+3NSHk9Yt3h
s7/LrU9vDfoPVROMatJR01XzFQKBgQDf/5kofDYQ/qcddosktkgxIyZBFuE4C/s+
nhiXl/Kd5Q+AP2o6kPsl5o4Jz2s3zBrmyRb733Zhb/rx/gbebTvqLjrTpyxXovmQ
8ecKeAS+IlrvAEDDT4c6ecAXR4zHZER00g0zbL4sX+fpKzON+jL6poA/el4MQySR
/DLJUx1nQQKBgQCLNQFG/2BrPXfNaupFWyDZW9CT/6JYJEUjN0B5bHCmr2VFYdjm
hWjA5WHLUBEQxCPiwsvCjccSRAyzDNQZfG7xuOXJlZR/P9ms/ce8Ry6hyO+nYEzb
qNXddQHSD+RjjAxUwCxdw3XNgFTQimE03EarO5zZdMT57RKa3AaBWePpbQKBgQCq
D4fcMNFrfaqqt8FUEgAlLiZw7En5Hz+Ufrr0/Kt6LNnj6EFiTYgfcjcMQ6mHJzKV
XL5SY4mg2D+RUectH4mJdae74QPNVTJcVQuv6wbOw45+PZbtsYddYenwwqWjDADd
IExdaoXHctjDMcVmWTozCg38I48biC5Pl0WHi86bAQKBgGBK6XUJPRYOsQFshunq
edxSbZBiYFDUj6SfOdaTSuU61KOWRTXJyuOBaB77usmZdwOrB4vy1XUT1uuPWKlx
SKmNoe/mk2xYiGdKvFDRRHh25zCxDWsQ2nMQfUFczTZ9wBwGs40wzm36fSgHZybq
Z3NIV2eNt6YBwkC69DzdazXT
-----END PRIVATE KEY-----
`)

	var policy = &sslconfigs.SSLPolicy{
		OCSPIsOn: true,
	}

	var nowTime = time.Now().Unix()

	policy.Certs = append(policy.Certs, &sslconfigs.SSLCertConfig{
		Id:            1,
		CertData:      certData,
		KeyData:       keyData,
		OCSP:          []byte("ocsp"),
		OCSPExpiresAt: nowTime + 1,
	})
	policy.Certs = append(policy.Certs, &sslconfigs.SSLCertConfig{
		Id:            1,
		CertData:      certData,
		KeyData:       keyData,
		OCSP:          []byte("ocsp"),
		OCSPExpiresAt: nowTime + 3,
	})
	policy.Certs = append(policy.Certs, &sslconfigs.SSLCertConfig{
		Id:            1,
		CertData:      certData,
		KeyData:       keyData,
		OCSP:          []byte("ocsp"),
		OCSPExpiresAt: nowTime + 2,
	})

	err := policy.Init()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(policy.OcspExpiresAt(), policy.OcspExpiresAt() == nowTime+1)

	time.Sleep(1 * time.Second)
	policy.CheckOCSP()
	t.Log(policy.OcspExpiresAt(), policy.OcspExpiresAt() == nowTime+2)
}
