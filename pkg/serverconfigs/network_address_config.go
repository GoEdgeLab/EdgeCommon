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

// 网络地址配置
type NetworkAddressConfig struct {
	Protocol  Protocol `yaml:"protocol" json:"protocol"`   // 协议，http、tcp、tcp4、tcp6、unix、udp等
	Host      string   `yaml:"host" json:"host"`           // 主机地址或主机名，支持变量
	PortRange string   `yaml:"portRange" json:"portRange"` // 端口范围，支持 8080、8080-8090、8080:8090

	minPort int
	maxPort int

	hostHasVariables bool
}

// 初始化
func (this *NetworkAddressConfig) Init() error {
	this.hostHasVariables = configutils.HasVariables(this.Host)

	// 8080
	if regexpSinglePort.MatchString(this.PortRange) {
		this.minPort = types.Int(this.PortRange)
		this.maxPort = this.minPort
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
		this.minPort = minPort
		this.maxPort = maxPort
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
		this.minPort = minPort
		this.maxPort = maxPort
		return nil
	}

	return nil
}

// 所有的地址列表，包含scheme
func (this *NetworkAddressConfig) FullAddresses() []string {
	if this.Protocol == ProtocolUnix {
		return []string{this.Protocol.String() + ":" + this.Host}
	}

	result := []string{}
	for i := this.minPort; i <= this.maxPort; i++ {
		host := this.Host
		result = append(result, this.Protocol.String()+"://"+host+":"+strconv.Itoa(i))
	}
	return result
}

// 选择其中一个地址
func (this *NetworkAddressConfig) PickAddress() string {
	if this.maxPort > this.minPort {
		return this.Host + ":" + strconv.Itoa(rands.Int(this.minPort, this.maxPort))
	}
	return this.Host + ":" + strconv.Itoa(this.minPort)
}

// 判断Host是否包含变量
func (this *NetworkAddressConfig) HostHasVariables() bool {
	return this.hostHasVariables
}
