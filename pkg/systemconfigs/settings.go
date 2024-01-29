package systemconfigs

type SettingCode = string

const (
	SettingCodeNodeMonitor           SettingCode = "nodeMonitor"         // 监控节点状态
	SettingCodeClusterHealthCheck    SettingCode = "clusterHealthCheck"  // 集群健康检查
	SettingCodeIPListVersion         SettingCode = "ipListVersion"       // IP名单的版本号
	SettingCodeAdminSecurityConfig   SettingCode = "adminSecurityConfig" // 管理员安全设置
	SettingCodeAdminUIConfig         SettingCode = "adminUIConfig"       // 管理员界面设置
	SettingCodeDatabaseConfigSetting SettingCode = "databaseConfig"      // 数据库相关配置
	SettingCodeAccessLogQueue        SettingCode = "accessLogQueue"      // 访问日志队列
	SettingCodeCheckUpdates          SettingCode = "checkUpdates"        // 检查自动更新配置

	SettingCodeUserServerConfig   SettingCode = "userServerConfig"   // 用户服务设置
	SettingCodeUserRegisterConfig SettingCode = "userRegisterConfig" // 用户注册配置
	SettingCodeUserUIConfig       SettingCode = "userUIConfig"       // 用户界面配置

	SettingCodeStandaloneInstanceInitialized SettingCode = "standaloneInstanceInitialized" // 单体实例是否已经被初始化：0 未被初始化, 1 已经成功初始化
)
