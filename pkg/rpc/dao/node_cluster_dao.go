package dao

import (
	"context"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

var SharedNodeClusterDAO = new(NodeClusterDAO)

// NodeClusterDAO 集群相关操作
type NodeClusterDAO struct {
	BaseDAO
}

// FindEnabledNodeCluster 查找集群
func (this *NodeClusterDAO) FindEnabledNodeCluster(ctx context.Context, clusterId int64) (*pb.NodeCluster, error) {
	clusterResp, err := this.RPC().NodeClusterRPC().FindEnabledNodeCluster(ctx, &pb.FindEnabledNodeClusterRequest{NodeClusterId: clusterId})
	if err != nil {
		return nil, err
	}
	return clusterResp.NodeCluster, nil
}

// FindEnabledNodeClusterConfigInfo 查找集群概要信息
func (this *NodeClusterDAO) FindEnabledNodeClusterConfigInfo(ctx context.Context, clusterId int64) (*pb.FindEnabledNodeClusterConfigInfoResponse, error) {
	return this.RPC().NodeClusterRPC().FindEnabledNodeClusterConfigInfo(ctx, &pb.FindEnabledNodeClusterConfigInfoRequest{NodeClusterId: clusterId})
}
