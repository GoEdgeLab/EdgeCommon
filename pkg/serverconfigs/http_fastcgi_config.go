package serverconfigs

import (
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
	"net"
	"path/filepath"
	"regexp"
	"time"
)

// HTTPFastcgiParam Fastcgi参数
type HTTPFastcgiParam struct {
	Name  string `yaml:"name" json:"name"`
	Value string `yaml:"value" json:"value"`
}

// HTTPFastcgiConfig Fastcgi配置
type HTTPFastcgiConfig struct {
	Id   int64 `yaml:"id" json:"id"`
	IsOn bool  `yaml:"isOn" json:"isOn"`

	// fastcgi地址配置
	// 支持unix:/tmp/php-fpm.sock ...
	Address string `yaml:"address" json:"address"`

	Index           string               `yaml:"index" json:"index"`                     // @TODO
	Params          []*HTTPFastcgiParam  `yaml:"params" json:"params"`                   // 参数
	ReadTimeout     *shared.TimeDuration `yaml:"readTimeout" json:"readTimeout"`         // @TODO 读取超时时间
	SendTimeout     *shared.TimeDuration `yaml:"sendTimeout" json:"sendTimeout"`         // @TODO 发送超时时间
	ConnTimeout     *shared.TimeDuration `yaml:"connTimeout" json:"connTimeout"`         // @TODO 连接超时时间
	Weight          int                  `yaml:"weight" json:"weight"`                   // TODO 权重
	PoolSize        int                  `yaml:"poolSize" json:"poolSize"`               // 连接池尺寸
	PathInfoPattern string               `yaml:"pathInfoPattern" json:"pathInfoPattern"` // PATH_INFO匹配正则

	network string // 协议：tcp, unix
	address string // 地址

	paramsMap      maps.Map
	readTimeout    time.Duration
	connTimeout    time.Duration
	pathInfoRegexp *regexp.Regexp
}

// Init 初始化
func (this *HTTPFastcgiConfig) Init() error {
	params := map[string]string{}
	for _, p := range this.Params {
		params[p.Name] = p.Value
	}
	this.paramsMap = maps.NewMap(params)
	if !this.paramsMap.Has("SCRIPT_FILENAME") {
		this.paramsMap["SCRIPT_FILENAME"] = ""
	}
	if !this.paramsMap.Has("REDIRECT_STATUS") {
		this.paramsMap["REDIRECT_STATUS"] = "200"
	}
	if !this.paramsMap.Has("GATEWAY_INTERFACE") {
		this.paramsMap["GATEWAY_INTERFACE"] = "CGI/1.1"
	}

	// 校验地址
	if regexp.MustCompile(`^\d+$`).MatchString(this.Address) {
		this.network = "tcp"
		this.address = "127.0.0.1:" + this.Address
	} else if regexp.MustCompile(`^(.*):(\d+)$`).MatchString(this.Address) {
		var matches = regexp.MustCompile(`^(.*):(\d+)$`).FindStringSubmatch(this.Address)
		ip := matches[1]
		port := matches[2]
		if len(ip) == 0 {
			ip = "127.0.0.1"
		}
		this.network = "tcp"
		this.address = ip + ":" + port
	} else if net.ParseIP(this.address) != nil {
		this.network = "tcp"
		this.address = configutils.QuoteIP(this.Address) + ":9000"
	} else if regexp.MustCompile("^unix:(.+)$").MatchString(this.Address) {
		matches := regexp.MustCompile("^unix:(.+)$").FindStringSubmatch(this.Address)
		path := matches[1]
		this.network = "unix"
		this.address = path
	} else if regexp.MustCompile("^[./].+$").MatchString(this.Address) {
		this.network = "unix"
		this.address = this.Address
	} else {
		return errors.New("invalid 'pass' format")
	}

	// 超时时间
	if this.ReadTimeout != nil {
		this.readTimeout = this.ReadTimeout.Duration()
	} else {
		this.readTimeout = 3 * time.Second
	}

	if this.ConnTimeout != nil {
		this.connTimeout = this.ConnTimeout.Duration()
	} else {
		this.connTimeout = 10 * time.Second
	}

	// PATH_INFO
	if len(this.PathInfoPattern) > 0 {
		reg, err := regexp.Compile(this.PathInfoPattern)
		if err != nil {
			return err
		}
		this.pathInfoRegexp = reg
	}

	return nil
}

// FilterParams 过滤参数
func (this *HTTPFastcgiConfig) FilterParams() maps.Map {
	params := maps.NewMap(this.paramsMap)

	// 自动添加参数
	script := params.GetString("SCRIPT_FILENAME")
	if len(script) > 0 {
		if !params.Has("SCRIPT_NAME") {
			params["SCRIPT_NAME"] = filepath.Base(script)
		}
		if !params.Has("DOCUMENT_ROOT") {
			params["DOCUMENT_ROOT"] = filepath.Dir(script)
		}
		if !params.Has("PWD") {
			params["PWD"] = filepath.Dir(script)
		}
	}

	return params
}

// ReadTimeoutDuration 超时时间
func (this *HTTPFastcgiConfig) ReadTimeoutDuration() time.Duration {
	if this.readTimeout <= 0 {
		this.readTimeout = 30 * time.Second
	}
	return this.readTimeout
}

// Network 网络协议
func (this *HTTPFastcgiConfig) Network() string {
	return this.network
}

// RealAddress 网络地址
func (this *HTTPFastcgiConfig) RealAddress() string {
	return this.address
}

// Param 读取参数
func (this *HTTPFastcgiConfig) Param(paramName string) string {
	for _, p := range this.Params {
		if p.Name == paramName {
			return p.Value
		}
	}
	return ""
}

// PathInfoRegexp PATH_INFO正则
func (this *HTTPFastcgiConfig) PathInfoRegexp() *regexp.Regexp {
	return this.pathInfoRegexp
}
