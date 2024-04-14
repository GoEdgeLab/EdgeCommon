package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/rands"
	"github.com/iwind/TeaGo/types"
	"regexp"
	"strconv"
	"strings"
)

var regexpSinglePort = regexp.MustCompile(`^\d+$`)

// NetworkAddressConfig 网络地址配置
type NetworkAddressConfig struct {
	Protocol  Protocol `yaml:"protocol" json:"protocol"`   // 协议，http、tcp、tcp4、tcp6、unix、udp等
	Host      string   `yaml:"host" json:"host"`           // 主机地址或主机名，支持变量
	PortRange string   `yaml:"portRange" json:"portRange"` // 端口范围，支持 8080、8080-8090、8080:8090

	MinPort int `yaml:"minPort" json:"minPort"` // minPort和maxPort只是用来记录PortRange分解后的结果，不需要用户输入
	MaxPort int `yaml:"maxPort" json:"maxPort"`

	hostHasVariables bool
}

// Init 初始化
func (this *NetworkAddressConfig) Init() error {
	this.hostHasVariables = configutils.HasVariables(this.Host)

	// 特殊端口自动修复，防止有些小白用户不了解HTTP和HTTPS的区别而选择了错误的协议
	if this.PortRange == "80" && this.Protocol == ProtocolHTTPS {
		this.Protocol = ProtocolHTTP
	}
	if this.PortRange == "443" && this.Protocol == ProtocolHTTP {
		this.Protocol = ProtocolHTTPS
	}

	// 8080
	if regexpSinglePort.MatchString(this.PortRange) {
		this.MinPort = types.Int(this.PortRange)
		this.MaxPort = this.MinPort
		return nil
	}

	// 8080:8090
	if strings.Contains(this.PortRange, ":") {
		pieces := strings.SplitN(this.PortRange, ":", 2)
		minPort := types.Int(pieces[0])
		maxPort := types.Int(pieces[1])
		if minPort > maxPort {
			minPort, maxPort = maxPort, minPort
		}
		this.MinPort = minPort
		this.MaxPort = maxPort
		return nil
	}

	// 8080-8090
	if strings.Contains(this.PortRange, "-") {
		pieces := strings.SplitN(this.PortRange, "-", 2)
		minPort := types.Int(pieces[0])
		maxPort := types.Int(pieces[1])
		if minPort > maxPort {
			minPort, maxPort = maxPort, minPort
		}
		this.MinPort = minPort
		this.MaxPort = maxPort
		return nil
	}

	return nil
}

// Addresses 所有的地址列表，不包括scheme
func (this *NetworkAddressConfig) Addresses() []string {
	result := []string{}
	for i := this.MinPort; i <= this.MaxPort; i++ {
		host := this.Host
		result = append(result, configutils.QuoteIP(host)+":"+strconv.Itoa(i))
	}
	return result
}

// FullAddresses 所有的地址列表，包含scheme
func (this *NetworkAddressConfig) FullAddresses() []string {
	result := []string{}
	for i := this.MinPort; i <= this.MaxPort; i++ {
		host := this.Host
		result = append(result, this.Protocol.String()+"://"+configutils.QuoteIP(host)+":"+strconv.Itoa(i))
	}
	return result
}

// PickAddress 选择其中一个地址
func (this *NetworkAddressConfig) PickAddress() string {
	if this.MaxPort > this.MinPort {
		return configutils.QuoteIP(this.Host) + ":" + strconv.Itoa(rands.Int(this.MinPort, this.MaxPort))
	}
	return configutils.QuoteIP(this.Host) + ":" + strconv.Itoa(this.MinPort)
}

// HostHasVariables 判断Host是否包含变量
func (this *NetworkAddressConfig) HostHasVariables() bool {
	return this.hostHasVariables
}
