package dao

import (
	"context"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
)

type SettingCode = string

const (
	SettingCodeServerGlobalConfig  SettingCode = "serverGlobalConfig"  // 服务相关全局设置
	SettingCodeNodeMonitor         SettingCode = "nodeMonitor"         // 监控节点状态
	SettingCodeClusterHealthCheck  SettingCode = "clusterHealthCheck"  // 集群健康检查
	SettingCodeIPListVersion       SettingCode = "ipListVersion"       // IP名单的版本号
	SettingCodeAdminSecurityConfig SettingCode = "adminSecurityConfig" // 管理员安全设置
)

var SharedSysSettingDAO = new(SysSettingDAO)

type SysSettingDAO struct {
	BaseDAO
}

// 读取服务全局配置
func (this *SysSettingDAO) ReadGlobalConfig(ctx context.Context) (*serverconfigs.GlobalConfig, error) {
	globalConfigResp, err := this.RPC().SysSettingRPC().ReadSysSetting(ctx, &pb.ReadSysSettingRequest{Code: SettingCodeServerGlobalConfig})
	if err != nil {
		return nil, err
	}
	if len(globalConfigResp.ValueJSON) == 0 {
		return nil, nil
	}
	globalConfig := &serverconfigs.GlobalConfig{}
	err = json.Unmarshal(globalConfigResp.ValueJSON, globalConfig)
	if err != nil {
		return nil, err
	}
	return globalConfig, nil
}
