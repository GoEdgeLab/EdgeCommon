package systemconfigs

type SettingCode = string

const (
	SettingCodeServerGlobalConfig    SettingCode = "serverGlobalConfig"  // 服务相关全局设置
	SettingCodeNodeMonitor           SettingCode = "nodeMonitor"         // 监控节点状态
	SettingCodeClusterHealthCheck    SettingCode = "clusterHealthCheck"  // 集群健康检查
	SettingCodeIPListVersion         SettingCode = "ipListVersion"       // IP名单的版本号
	SettingCodeAdminSecurityConfig   SettingCode = "adminSecurityConfig" // 管理员安全设置
	SettingCodeDatabaseConfigSetting SettingCode = "databaseConfig"      // 数据库相关配置
	SettingCodeNSAccessLogSetting    SettingCode = "nsAccessLogSetting"  // NS相关全局配置

	SettingCodeNSNodeMonitor SettingCode = "nsNodeMonitor" // 监控NS节点状态

	SettingCodeReportNodeGlobalSetting SettingCode = "reportNodeGlobalSetting" // 区域监控节点全局配置
)
