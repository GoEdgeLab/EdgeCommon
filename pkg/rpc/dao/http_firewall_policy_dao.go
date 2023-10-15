package dao

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/ipconfigs"
	"github.com/iwind/TeaGo/maps"
)

var SharedHTTPFirewallPolicyDAO = new(HTTPFirewallPolicyDAO)

// HTTPFirewallPolicyDAO WAF策略相关
type HTTPFirewallPolicyDAO struct {
	BaseDAO
}

// FindEnabledHTTPFirewallPolicy 查找WAF策略基本信息
func (this *HTTPFirewallPolicyDAO) FindEnabledHTTPFirewallPolicy(ctx context.Context, policyId int64) (*pb.HTTPFirewallPolicy, error) {
	resp, err := this.RPC().HTTPFirewallPolicyRPC().FindEnabledHTTPFirewallPolicy(ctx, &pb.FindEnabledHTTPFirewallPolicyRequest{HttpFirewallPolicyId: policyId})
	if err != nil {
		return nil, err
	}
	return resp.HttpFirewallPolicy, nil
}

// FindEnabledHTTPFirewallPolicyConfig 查找WAF策略配置
func (this *HTTPFirewallPolicyDAO) FindEnabledHTTPFirewallPolicyConfig(ctx context.Context, policyId int64) (*firewallconfigs.HTTPFirewallPolicy, error) {
	resp, err := this.RPC().HTTPFirewallPolicyRPC().FindEnabledHTTPFirewallPolicyConfig(ctx, &pb.FindEnabledHTTPFirewallPolicyConfigRequest{HttpFirewallPolicyId: policyId})
	if err != nil {
		return nil, err
	}
	if len(resp.HttpFirewallPolicyJSON) == 0 {
		return nil, nil
	}
	firewallPolicy := &firewallconfigs.HTTPFirewallPolicy{}
	err = json.Unmarshal(resp.HttpFirewallPolicyJSON, firewallPolicy)
	if err != nil {
		return nil, err
	}
	return firewallPolicy, nil
}

// FindEnabledHTTPFirewallPolicyInboundConfig 查找WAF的Inbound
func (this *HTTPFirewallPolicyDAO) FindEnabledHTTPFirewallPolicyInboundConfig(ctx context.Context, policyId int64) (*firewallconfigs.HTTPFirewallInboundConfig, error) {
	config, err := this.FindEnabledHTTPFirewallPolicyConfig(ctx, policyId)
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, errors.New("not found")
	}
	return config.Inbound, nil
}

// FindEnabledPolicyIPListIdWithType 根据类型查找WAF的IP名单
func (this *HTTPFirewallPolicyDAO) FindEnabledPolicyIPListIdWithType(ctx context.Context, policyId int64, listType ipconfigs.IPListType) (int64, error) {
	switch listType {
	case ipconfigs.IPListTypeWhite:
		return this.FindEnabledPolicyWhiteIPListId(ctx, policyId)
	case ipconfigs.IPListTypeBlack:
		return this.FindEnabledPolicyBlackIPListId(ctx, policyId)
	default:
		return 0, errors.New("invalid ip list type '" + listType + "'")
	}
}

// FindEnabledPolicyWhiteIPListId 查找WAF的白名单
func (this *HTTPFirewallPolicyDAO) FindEnabledPolicyWhiteIPListId(ctx context.Context, policyId int64) (int64, error) {
	config, err := this.FindEnabledHTTPFirewallPolicyConfig(ctx, policyId)
	if err != nil {
		return 0, err
	}
	if config == nil {
		return 0, errors.New("not found")
	}
	if config.Inbound == nil {
		config.Inbound = &firewallconfigs.HTTPFirewallInboundConfig{IsOn: true}
	}
	if config.Inbound.AllowListRef == nil || config.Inbound.AllowListRef.ListId == 0 {
		createResp, err := this.RPC().IPListRPC().CreateIPList(ctx, &pb.CreateIPListRequest{
			Type:        "white",
			Name:        "白名单",
			Code:        "white",
			TimeoutJSON: nil,
		})
		if err != nil {
			return 0, err
		}
		listId := createResp.IpListId
		config.Inbound.AllowListRef = &ipconfigs.IPListRef{
			IsOn:   true,
			ListId: listId,
		}
		inboundJSON, err := json.Marshal(config.Inbound)
		if err != nil {
			return 0, err
		}
		_, err = this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallInboundConfig(ctx, &pb.UpdateHTTPFirewallInboundConfigRequest{
			HttpFirewallPolicyId: policyId,
			InboundJSON:          inboundJSON,
		})
		if err != nil {
			return 0, err
		}
		return listId, nil
	}

	return config.Inbound.AllowListRef.ListId, nil
}

// FindEnabledPolicyBlackIPListId 查找WAF的黑名单
func (this *HTTPFirewallPolicyDAO) FindEnabledPolicyBlackIPListId(ctx context.Context, policyId int64) (int64, error) {
	config, err := this.FindEnabledHTTPFirewallPolicyConfig(ctx, policyId)
	if err != nil {
		return 0, err
	}
	if config == nil {
		return 0, errors.New("not found")
	}
	if config.Inbound == nil {
		config.Inbound = &firewallconfigs.HTTPFirewallInboundConfig{IsOn: true}
	}
	if config.Inbound.DenyListRef == nil || config.Inbound.DenyListRef.ListId == 0 {
		createResp, err := this.RPC().IPListRPC().CreateIPList(ctx, &pb.CreateIPListRequest{
			Type:        "black",
			Name:        "黑名单",
			Code:        "black",
			TimeoutJSON: nil,
		})
		if err != nil {
			return 0, err
		}
		listId := createResp.IpListId
		config.Inbound.DenyListRef = &ipconfigs.IPListRef{
			IsOn:   true,
			ListId: listId,
		}
		inboundJSON, err := json.Marshal(config.Inbound)
		if err != nil {
			return 0, err
		}
		_, err = this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallInboundConfig(ctx, &pb.UpdateHTTPFirewallInboundConfigRequest{
			HttpFirewallPolicyId: policyId,
			InboundJSON:          inboundJSON,
		})
		if err != nil {
			return 0, err
		}
		return listId, nil
	}

	return config.Inbound.DenyListRef.ListId, nil
}

// FindEnabledHTTPFirewallPolicyWithServerId 根据服务Id查找WAF策略
func (this *HTTPFirewallPolicyDAO) FindEnabledHTTPFirewallPolicyWithServerId(ctx context.Context, serverId int64) (*pb.HTTPFirewallPolicy, error) {
	serverResp, err := this.RPC().ServerRPC().FindEnabledServer(ctx, &pb.FindEnabledServerRequest{ServerId: serverId})
	if err != nil {
		return nil, err
	}
	server := serverResp.Server
	if server == nil {
		return nil, nil
	}
	if server.NodeCluster == nil {
		return nil, nil
	}
	clusterId := server.NodeCluster.Id
	cluster, err := SharedNodeClusterDAO.FindEnabledNodeCluster(ctx, clusterId)
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return nil, nil
	}
	if cluster.HttpFirewallPolicyId == 0 {
		return nil, nil
	}
	return SharedHTTPFirewallPolicyDAO.FindEnabledHTTPFirewallPolicy(ctx, cluster.HttpFirewallPolicyId)
}

// FindHTTPFirewallActionConfigs 查找动作相关信息
func (this *HTTPFirewallPolicyDAO) FindHTTPFirewallActionConfigs(ctx context.Context, actions []*firewallconfigs.HTTPFirewallActionConfig) ([]maps.Map, error) {
	var actionConfigs = []maps.Map{}
	for _, action := range actions {
		def := firewallconfigs.FindActionDefinition(action.Code)
		if def == nil {
			continue
		}
		if action.Options == nil {
			action.Options = maps.Map{}
		}

		switch action.Code {
		case firewallconfigs.HTTPFirewallActionRecordIP:
			var listId = action.Options.GetInt64("ipListId")
			listResp, err := this.RPC().IPListRPC().FindEnabledIPList(ctx, &pb.FindEnabledIPListRequest{IpListId: listId})
			if err != nil {
				return nil, err
			}
			if listId == 0 {
				action.Options["ipListName"] = "全局黑名单"
			} else if listResp.IpList != nil {
				action.Options["ipListName"] = listResp.IpList.Name
			} else {
				action.Options["ipListName"] = action.Options.GetString("ipListName") + "(已删除)"
			}
		case firewallconfigs.HTTPFirewallActionGoGroup:
			groupId := action.Options.GetInt64("groupId")
			groupResp, err := this.RPC().HTTPFirewallRuleGroupRPC().FindEnabledHTTPFirewallRuleGroup(ctx, &pb.FindEnabledHTTPFirewallRuleGroupRequest{FirewallRuleGroupId: groupId})
			if err != nil {
				return nil, err
			}
			if groupResp.FirewallRuleGroup != nil {
				action.Options["groupName"] = groupResp.FirewallRuleGroup.Name
			} else {
				action.Options["groupName"] = action.Options.GetString("groupName") + "(已删除)"
			}
		case firewallconfigs.HTTPFirewallActionGoSet:
			groupId := action.Options.GetInt64("groupId")
			groupResp, err := this.RPC().HTTPFirewallRuleGroupRPC().FindEnabledHTTPFirewallRuleGroup(ctx, &pb.FindEnabledHTTPFirewallRuleGroupRequest{FirewallRuleGroupId: groupId})
			if err != nil {
				return nil, err
			}
			if groupResp.FirewallRuleGroup != nil {
				action.Options["groupName"] = groupResp.FirewallRuleGroup.Name
			} else {
				action.Options["groupName"] = action.Options.GetString("groupName") + "(已删除)"
			}

			setId := action.Options.GetInt64("setId")
			setResp, err := this.RPC().HTTPFirewallRuleSetRPC().FindEnabledHTTPFirewallRuleSet(ctx, &pb.FindEnabledHTTPFirewallRuleSetRequest{FirewallRuleSetId: setId})
			if err != nil {
				return nil, err
			}
			if setResp.FirewallRuleSet != nil {
				action.Options["setName"] = setResp.FirewallRuleSet.Name
			} else {
				action.Options["setName"] = action.Options.GetString("setName") + "(已删除)"
			}
		}

		actionConfigs = append(actionConfigs, maps.Map{
			"name":     def.Name,
			"code":     def.Code,
			"category": def.Category,
			"options":  action.Options,
		})
	}
	return actionConfigs, nil
}
