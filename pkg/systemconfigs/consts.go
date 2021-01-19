package systemconfigs

type SettingCode = string

const (
	SettingCodeServerGlobalConfig    SettingCode = "serverGlobalConfig"  // 服务相关全局设置
	SettingCodeNodeMonitor           SettingCode = "nodeMonitor"         // 监控节点状态
	SettingCodeClusterHealthCheck    SettingCode = "clusterHealthCheck"  // 集群健康检查
	SettingCodeIPListVersion         SettingCode = "ipListVersion"       // IP名单的版本号
	SettingCodeAdminSecurityConfig   SettingCode = "adminSecurityConfig" // 管理员安全设置
	SettingCodeDatabaseConfigSetting SettingCode = "databaseConfig"      // 数据库相关配置
)
