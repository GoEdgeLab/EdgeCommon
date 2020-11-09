package firewallconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/ipconfigs"

type HTTPFirewallInboundConfig struct {
	IsOn      bool                        `yaml:"isOn" json:"isOn"`
	GroupRefs []*HTTPFirewallRuleGroupRef `yaml:"groupRefs" json:"groupRefs"`
	Groups    []*HTTPFirewallRuleGroup    `yaml:"groups" json:"groups"`

	// 地区相关
	Region *HTTPFirewallRegionConfig `yaml:"region" json:"region"`

	// IP名单
	WhiteListRef *ipconfigs.IPListRef `yaml:"whiteListRef" json:"whiteListRef"`
	BlackListRef *ipconfigs.IPListRef `yaml:"blackListRef" json:"blackListRef"`
	GreyListRef  *ipconfigs.IPListRef `yaml:"greyListRef" json:"greyListRef"`
}

// 初始化
func (this *HTTPFirewallInboundConfig) Init() error {
	for _, group := range this.Groups {
		err := group.Init()
		if err != nil {
			return err
		}
	}

	if this.Region != nil {
		err := this.Region.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 根据Code查找Group
func (this *HTTPFirewallInboundConfig) FindGroupWithCode(code string) *HTTPFirewallRuleGroup {
	for _, group := range this.Groups {
		if group.Code == code {
			return group
		}
	}
	return nil
}

// 删除某个分组
func (this *HTTPFirewallInboundConfig) RemoveRuleGroup(groupId int64) {
	groups := []*HTTPFirewallRuleGroup{}
	refs := []*HTTPFirewallRuleGroupRef{}
	for _, g := range this.Groups {
		if g.Id == groupId {
			continue
		}
		groups = append(groups, g)
	}
	for _, ref := range this.GroupRefs {
		if ref.GroupId == groupId {
			continue
		}
		refs = append(refs, ref)
	}
	this.Groups = groups
	this.GroupRefs = refs
}
