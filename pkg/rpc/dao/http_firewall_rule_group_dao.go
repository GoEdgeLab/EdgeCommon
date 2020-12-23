package dao

import (
	"context"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
)

var SharedHTTPFirewallRuleGroupDAO = new(HTTPFirewallRuleGroupDAO)

type HTTPFirewallRuleGroupDAO struct {
	BaseDAO
}

// 查找分组配置
func (this *HTTPFirewallRuleGroupDAO) FindRuleGroupConfig(ctx context.Context, groupId int64) (*firewallconfigs.HTTPFirewallRuleGroup, error) {
	groupResp, err := this.RPC().HTTPFirewallRuleGroupRPC().FindEnabledHTTPFirewallRuleGroupConfig(ctx, &pb.FindEnabledHTTPFirewallRuleGroupConfigRequest{FirewallRuleGroupId: groupId})
	if err != nil {
		return nil, err
	}

	if len(groupResp.FirewallRuleGroupJSON) == 0 {
		return nil, nil
	}

	groupConfig := &firewallconfigs.HTTPFirewallRuleGroup{}
	err = json.Unmarshal(groupResp.FirewallRuleGroupJSON, groupConfig)
	if err != nil {
		return nil, err
	}

	return groupConfig, nil
}
