package serverconfigs

// BaseProtocol 协议基础数据结构
type BaseProtocol struct {
	IsOn   bool                    `yaml:"isOn" json:"isOn"`     // 是否开启
	Listen []*NetworkAddressConfig `yaml:"listen" json:"listen"` // 绑定的网络地址
}

// InitBase 初始化
func (this *BaseProtocol) InitBase() error {
	for _, addr := range this.Listen {
		err := addr.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// FullAddresses 获取完整的地址列表
func (this *BaseProtocol) FullAddresses() []string {
	result := []string{}
	for _, addr := range this.Listen {
		result = append(result, addr.FullAddresses()...)
	}
	return result
}

// AddListen 添加地址
func (this *BaseProtocol) AddListen(addr ...*NetworkAddressConfig) {
	this.Listen = append(this.Listen, addr...)
}

// AllPorts 获取所有端口号
func (this *BaseProtocol) AllPorts() []int {
	var ports = []int{}
	var portMap = map[int]bool{}
	for _, listen := range this.Listen {
		if listen.MinPort > 0 && listen.MaxPort > 0 {
			for port := listen.MinPort; port <= listen.MaxPort; port++ {
				_, ok := portMap[port]
				if ok {
					continue
				}
				ports = append(ports, port)
				portMap[port] = true
			}
		}
	}
	return ports
}
