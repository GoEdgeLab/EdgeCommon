package dao

import (
	"context"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/ipconfigs"
)

var SharedIPListDAO = new(IPListDAO)

type IPListDAO struct {
	BaseDAO
}

// FindAllowIPListIdWithServerId 查找服务的允许IP列表
func (this *IPListDAO) FindAllowIPListIdWithServerId(ctx context.Context, serverId int64) (int64, error) {
	webConfig, err := SharedHTTPWebDAO.FindWebConfigWithServerId(ctx, serverId)
	if err != nil {
		return 0, err
	}
	if webConfig == nil {
		return 0, nil
	}
	if webConfig.FirewallPolicy == nil || webConfig.FirewallPolicy.Inbound == nil || webConfig.FirewallPolicy.Inbound.AllowListRef == nil {
		return 0, nil
	}
	return webConfig.FirewallPolicy.Inbound.AllowListRef.ListId, nil
}

// FindDenyIPListIdWithServerId 查找服务的禁止IP列表
func (this *IPListDAO) FindDenyIPListIdWithServerId(ctx context.Context, serverId int64) (int64, error) {
	webConfig, err := SharedHTTPWebDAO.FindWebConfigWithServerId(ctx, serverId)
	if err != nil {
		return 0, err
	}
	if webConfig == nil {
		return 0, nil
	}
	if webConfig.FirewallPolicy == nil || webConfig.FirewallPolicy.Inbound == nil || webConfig.FirewallPolicy.Inbound.DenyListRef == nil {
		return 0, nil
	}
	return webConfig.FirewallPolicy.Inbound.DenyListRef.ListId, nil
}

// CreateIPListForServerId 为服务创建IP名单
func (this *IPListDAO) CreateIPListForServerId(ctx context.Context, serverId int64, listType string) (int64, error) {
	webConfig, err := SharedHTTPWebDAO.FindWebConfigWithServerId(ctx, serverId)
	if err != nil {
		return 0, err
	}
	if webConfig == nil {
		return 0, nil
	}
	if webConfig.FirewallPolicy == nil || webConfig.FirewallPolicy.Id == 0 {
		isOn := webConfig.FirewallRef != nil && webConfig.FirewallRef.IsOn
		_, err = SharedHTTPWebDAO.InitEmptyHTTPFirewallPolicy(ctx, serverId, webConfig.Id, isOn)
		if err != nil {
			return 0, errors.Wrap(err)
		}
		webConfig, err = SharedHTTPWebDAO.FindWebConfigWithServerId(ctx, serverId)
		if err != nil {
			return 0, err
		}
		if webConfig == nil {
			return 0, nil
		}
		if webConfig.FirewallPolicy == nil {
			return 0, nil
		}
	}

	inbound := webConfig.FirewallPolicy.Inbound
	if inbound == nil {
		inbound = &firewallconfigs.HTTPFirewallInboundConfig{
			IsOn: true,
		}
	}
	if listType == "white" {
		if inbound.AllowListRef == nil {
			inbound.AllowListRef = &ipconfigs.IPListRef{
				IsOn: true,
			}
		}
		if inbound.AllowListRef.ListId > 0 {
			return inbound.AllowListRef.ListId, nil
		}
	} else if listType == "black" {
		if inbound.DenyListRef == nil {
			inbound.DenyListRef = &ipconfigs.IPListRef{
				IsOn: true,
			}
		}
		if inbound.DenyListRef.ListId > 0 {
			return inbound.DenyListRef.ListId, nil
		}
	}

	ipListResp, err := this.RPC().IPListRPC().CreateIPList(ctx, &pb.CreateIPListRequest{
		Type:        listType,
		Name:        "IP名单",
		Code:        listType,
		TimeoutJSON: nil,
	})
	if err != nil {
		return 0, errors.Wrap(err)
	}

	if listType == "white" {
		inbound.AllowListRef.ListId = ipListResp.IpListId
	} else if listType == "black" {
		inbound.DenyListRef.ListId = ipListResp.IpListId
	}
	inboundJSON, err := json.Marshal(inbound)
	if err != nil {
		return 0, errors.Wrap(err)
	}
	_, err = this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallInboundConfig(ctx, &pb.UpdateHTTPFirewallInboundConfigRequest{
		HttpFirewallPolicyId: webConfig.FirewallPolicy.Id,
		InboundJSON:          inboundJSON,
	})
	if err != nil {
		return 0, errors.Wrap(err)
	}

	return ipListResp.IpListId, nil
}
