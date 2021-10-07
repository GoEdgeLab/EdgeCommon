package dao

import (
	"context"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
)

var SharedServerDAO = new(ServerDAO)

type ServerDAO struct {
	BaseDAO
}

// FindEnabledServerConfig 查找服务配置
func (this *ServerDAO) FindEnabledServerConfig(ctx context.Context, serverId int64) (*serverconfigs.ServerConfig, error) {
	resp, err := this.RPC().ServerRPC().FindEnabledServerConfig(ctx, &pb.FindEnabledServerConfigRequest{ServerId: serverId})
	if err != nil {
		return nil, err
	}
	if len(resp.ServerJSON) == 0 {
		return nil, nil
	}
	config := &serverconfigs.ServerConfig{}
	err = json.Unmarshal(resp.ServerJSON, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// FindEnabledServer 查找服务
func (this *ServerDAO) FindEnabledServer(ctx context.Context, serverId int64) (*pb.Server, error) {
	resp, err := this.RPC().ServerRPC().FindEnabledServer(ctx, &pb.FindEnabledServerRequest{ServerId: serverId})
	if err != nil {
		return nil, err
	}
	return resp.Server, nil
}
