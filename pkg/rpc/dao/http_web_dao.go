package dao

import (
	"context"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
)

var SharedHTTPWebDAO = new(HTTPWebDAO)

type HTTPWebDAO struct {
	BaseDAO
}

// 根据ServerId查找Web配置
func (this *HTTPWebDAO) FindWebConfigWithServerId(ctx context.Context, serverId int64) (*serverconfigs.HTTPWebConfig, error) {
	resp, err := this.RPC().ServerRPC().FindAndInitServerWebConfig(ctx, &pb.FindAndInitServerWebConfigRequest{ServerId: serverId})
	if err != nil {
		return nil, err
	}
	config := &serverconfigs.HTTPWebConfig{}
	err = json.Unmarshal(resp.WebJSON, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// 根据LocationId查找Web配置
func (this *HTTPWebDAO) FindWebConfigWithLocationId(ctx context.Context, locationId int64) (*serverconfigs.HTTPWebConfig, error) {
	resp, err := this.RPC().HTTPLocationRPC().FindAndInitHTTPLocationWebConfig(ctx, &pb.FindAndInitHTTPLocationWebConfigRequest{LocationId: locationId})
	if err != nil {
		return nil, err
	}
	config := &serverconfigs.HTTPWebConfig{}
	err = json.Unmarshal(resp.WebJSON, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// 根据WebId查找Web配置
func (this *HTTPWebDAO) FindWebConfigWithId(ctx context.Context, webId int64) (*serverconfigs.HTTPWebConfig, error) {
	resp, err := this.RPC().HTTPWebRPC().FindEnabledHTTPWebConfig(ctx, &pb.FindEnabledHTTPWebConfigRequest{WebId: webId})
	if err != nil {
		return nil, err
	}
	config := &serverconfigs.HTTPWebConfig{}
	err = json.Unmarshal(resp.WebJSON, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// 初始化防火墙设置
func (this *HTTPWebDAO) InitEmptyHTTPFirewallPolicy(ctx context.Context, serverId int64, webId int64, isOn bool) (int64, error) {
	// 创建FirewallPolicy
	firewallPolicyIdResp, err := this.RPC().HTTPFirewallPolicyRPC().CreateEmptyHTTPFirewallPolicy(ctx, &pb.CreateEmptyHTTPFirewallPolicyRequest{
		ServerId:    serverId,
		IsOn:        true,
		Name:        "用户自定义",
		Description: "",
	})
	if err != nil {
		return 0, errors.Wrap(err)
	}

	policyId := firewallPolicyIdResp.HttpFirewallPolicyId

	firewallRef := &firewallconfigs.HTTPFirewallRef{
		IsPrior:          false,
		IsOn:             isOn,
		FirewallPolicyId: policyId,
	}
	firewallRefJSON, err := json.Marshal(firewallRef)
	if err != nil {
		return 0, errors.Wrap(err)
	}

	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebFirewall(ctx, &pb.UpdateHTTPWebFirewallRequest{
		WebId:        webId,
		FirewallJSON: firewallRefJSON,
	})
	if err != nil {
		return 0, errors.Wrap(err)
	}

	return policyId, nil
}
