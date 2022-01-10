package firewallconfigs

import "encoding/json"

// HTTPFirewallPolicy 防火墙策略
type HTTPFirewallPolicy struct {
	Id               int64                       `yaml:"id" json:"id"`
	IsOn             bool                        `yaml:"isOn" json:"isOn"`
	Name             string                      `yaml:"name" json:"name"`
	Description      string                      `yaml:"description" json:"description"`
	Inbound          *HTTPFirewallInboundConfig  `yaml:"inbound" json:"inbound"`
	Outbound         *HTTPFirewallOutboundConfig `yaml:"outbound" json:"outbound"`
	BlockOptions     *HTTPFirewallBlockAction    `yaml:"blockOptions" json:"blockOptions"`
	Mode             FirewallMode                `yaml:"mode" json:"mode"`
	UseLocalFirewall bool                        `yaml:"useLocalFirewall" json:"useLocalFirewall"`
	SYNFlood         *SYNFloodConfig             `yaml:"synFlood" json:"synFlood"`
}

// Init 初始化
func (this *HTTPFirewallPolicy) Init() error {
	if this.Inbound != nil {
		err := this.Inbound.Init()
		if err != nil {
			return err
		}
	}

	if this.Outbound != nil {
		err := this.Outbound.Init()
		if err != nil {
			return err
		}
	}

	if this.SYNFlood != nil {
		err := this.SYNFlood.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// AllRuleGroups 获取所有分组
func (this *HTTPFirewallPolicy) AllRuleGroups() []*HTTPFirewallRuleGroup {
	result := []*HTTPFirewallRuleGroup{}
	if this.Inbound != nil {
		result = append(result, this.Inbound.Groups...)
	}
	if this.Outbound != nil {
		result = append(result, this.Outbound.Groups...)
	}
	return result
}

// FindRuleGroupWithCode 根据代号查找分组
func (this *HTTPFirewallPolicy) FindRuleGroupWithCode(code string) *HTTPFirewallRuleGroup {
	for _, g := range this.AllRuleGroups() {
		if g.Code == code {
			return g
		}
	}
	return nil
}

// FindRuleGroupWithName 根据名称查找分组
func (this *HTTPFirewallPolicy) FindRuleGroupWithName(name string) *HTTPFirewallRuleGroup {
	for _, g := range this.AllRuleGroups() {
		if g.Name == name {
			return g
		}
	}
	return nil
}

// FindRuleGroup 根据ID查找分组
func (this *HTTPFirewallPolicy) FindRuleGroup(groupId int64) *HTTPFirewallRuleGroup {
	for _, g := range this.AllRuleGroups() {
		if g.Id == groupId {
			return g
		}
	}
	return nil
}

// RemoveRuleGroup 删除某个分组
func (this *HTTPFirewallPolicy) RemoveRuleGroup(groupId int64) {
	if this.Inbound != nil {
		this.Inbound.RemoveRuleGroup(groupId)
	}
	if this.Outbound != nil {
		this.Outbound.RemoveRuleGroup(groupId)
	}
}

// InboundJSON Inbound JSON
func (this *HTTPFirewallPolicy) InboundJSON() ([]byte, error) {
	if this.Inbound == nil {
		return []byte("null"), nil
	}
	groups := this.Inbound.Groups
	this.Inbound.Groups = nil
	defer func() {
		this.Inbound.Groups = groups
	}()
	return json.Marshal(this.Inbound)
}

// OutboundJSON Outbound JSON
func (this *HTTPFirewallPolicy) OutboundJSON() ([]byte, error) {
	if this.Inbound == nil {
		return []byte("null"), nil
	}
	groups := this.Outbound.Groups
	this.Outbound.Groups = nil
	defer func() {
		this.Outbound.Groups = groups
	}()
	return json.Marshal(this.Outbound)
}
