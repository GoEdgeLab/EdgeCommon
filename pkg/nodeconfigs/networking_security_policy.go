// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodeconfigs

import (
	"bytes"
	"encoding/json"
)

type NetworkSecurityStatus = string

const (
	NetworkSecurityStatusAuto NetworkSecurityStatus = "auto"
	NetworkSecurityStatusOn   NetworkSecurityStatus = "on"
	NetworkSecurityStatusOff  NetworkSecurityStatus = "off"
)

// NetworkSecurityPolicy 节点网络安全策略
type NetworkSecurityPolicy struct {
	Status NetworkSecurityStatus `json:"status"` // 启用状态

	TCP  struct{} `json:"tcp"`  // TODO
	UDP  struct{} `json:"udp"`  // TODO
	ICMP struct{} `json:"icmp"` // TODO
}

func NewNetworkSecurityPolicy() *NetworkSecurityPolicy {
	var policy = &NetworkSecurityPolicy{}
	policy.Status = NetworkSecurityStatusAuto
	return policy
}

func (this *NetworkSecurityPolicy) Init() error {
	return nil
}

func (this *NetworkSecurityPolicy) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}

func (this *NetworkSecurityPolicy) IsOn() bool {
	return this.Status != NetworkSecurityStatusOff
}

func (this *NetworkSecurityPolicy) IsSame(anotherPolicy *NetworkSecurityPolicy) bool {
	data1, _ := json.Marshal(this)
	data2, _ := json.Marshal(anotherPolicy)
	return bytes.Equal(data1, data2)
}
