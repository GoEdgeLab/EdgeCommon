package firewallconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/ipconfigs"
)

// HTTPFirewallInboundConfig HTTP防火墙入口配置
type HTTPFirewallInboundConfig struct {
	IsOn      bool                        `yaml:"isOn" json:"isOn"`
	GroupRefs []*HTTPFirewallRuleGroupRef `yaml:"groupRefs" json:"groupRefs"`
	Groups    []*HTTPFirewallRuleGroup    `yaml:"groups" json:"groups"`

	// 地区相关
	Region *HTTPFirewallRegionConfig `yaml:"region" json:"region"`

	// IP名单
	AllowListRef *ipconfigs.IPListRef `yaml:"whiteListRef" json:"whiteListRef"`
	DenyListRef  *ipconfigs.IPListRef `yaml:"blackListRef" json:"blackListRef"`
	GreyListRef  *ipconfigs.IPListRef `yaml:"greyListRef" json:"greyListRef"`

	// 绑定的IP名单
	PublicAllowListRefs []*ipconfigs.IPListRef `yaml:"publicWhiteListRefs" json:"publicWhiteListRefs"`
	PublicDenyListRefs  []*ipconfigs.IPListRef `yaml:"publicBlackListRefs" json:"publicBlackListRefs"`
	PublicGreyListRefs  []*ipconfigs.IPListRef `yaml:"publicGreyListRefs" json:"publicGreyListRefs"`

	allAllowListRefs []*ipconfigs.IPListRef
	allDenyListRefs  []*ipconfigs.IPListRef
	allGreyListRefs  []*ipconfigs.IPListRef
}

// Init 初始化
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

	this.allAllowListRefs = []*ipconfigs.IPListRef{}
	if this.AllowListRef != nil {
		this.allAllowListRefs = append(this.allAllowListRefs, this.AllowListRef)
	}
	if len(this.PublicAllowListRefs) > 0 {
		this.allAllowListRefs = append(this.allAllowListRefs, this.PublicAllowListRefs...)
	}

	this.allDenyListRefs = []*ipconfigs.IPListRef{}
	if this.DenyListRef != nil {
		this.allDenyListRefs = append(this.allDenyListRefs, this.DenyListRef)
	}
	if len(this.PublicDenyListRefs) > 0 {
		this.allDenyListRefs = append(this.allDenyListRefs, this.PublicDenyListRefs...)
	}

	this.allGreyListRefs = []*ipconfigs.IPListRef{}
	if this.GreyListRef != nil {
		this.allGreyListRefs = append(this.allGreyListRefs, this.GreyListRef)
	}
	if len(this.PublicGreyListRefs) > 0 {
		this.allGreyListRefs = append(this.allGreyListRefs, this.PublicGreyListRefs...)
	}

	return nil
}

// FindGroupWithCode 根据Code查找Group
func (this *HTTPFirewallInboundConfig) FindGroupWithCode(code string) *HTTPFirewallRuleGroup {
	for _, group := range this.Groups {
		if group.Code == code {
			return group
		}
	}
	return nil
}

// RemoveRuleGroup 删除某个分组
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

// AddPublicList 绑定公用的IP名单
func (this *HTTPFirewallInboundConfig) AddPublicList(listId int64, listType string) {
	var refs []*ipconfigs.IPListRef
	switch listType {
	case ipconfigs.IPListTypeBlack:
		refs = this.PublicDenyListRefs
	case ipconfigs.IPListTypeWhite:
		refs = this.PublicAllowListRefs
	case ipconfigs.IPListTypeGrey:
		refs = this.PublicGreyListRefs
	}
	var found = false
	for _, ref := range refs {
		if ref.ListId == listId {
			found = true
			ref.IsOn = true
			break
		}
	}
	if !found {
		refs = append(refs, &ipconfigs.IPListRef{
			IsOn:   true,
			ListId: listId,
		})
	}
	switch listType {
	case ipconfigs.IPListTypeBlack:
		this.PublicDenyListRefs = refs
	case ipconfigs.IPListTypeWhite:
		this.PublicAllowListRefs = refs
	case ipconfigs.IPListTypeGrey:
		this.PublicGreyListRefs = refs
	}
}

// RemovePublicList 解绑公用的IP名单
func (this *HTTPFirewallInboundConfig) RemovePublicList(listId int64, listType string) {
	var refs []*ipconfigs.IPListRef
	switch listType {
	case ipconfigs.IPListTypeBlack:
		refs = this.PublicDenyListRefs
	case ipconfigs.IPListTypeWhite:
		refs = this.PublicAllowListRefs
	case ipconfigs.IPListTypeGrey:
		refs = this.PublicGreyListRefs
	}
	var newRefs = []*ipconfigs.IPListRef{}
	for _, ref := range refs {
		if ref.ListId == listId {
			continue
		}
		newRefs = append(newRefs, ref)
	}
	switch listType {
	case ipconfigs.IPListTypeBlack:
		this.PublicDenyListRefs = newRefs
	case ipconfigs.IPListTypeWhite:
		this.PublicAllowListRefs = newRefs
	case ipconfigs.IPListTypeGrey:
		this.PublicGreyListRefs = newRefs
	}
}

// AllAllowListRefs 获取所有允许的IP名单
func (this *HTTPFirewallInboundConfig) AllAllowListRefs() []*ipconfigs.IPListRef {
	return this.allAllowListRefs
}

// AllDenyListRefs 获取所有禁止的IP名单
func (this *HTTPFirewallInboundConfig) AllDenyListRefs() []*ipconfigs.IPListRef {
	return this.allDenyListRefs
}
