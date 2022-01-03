package systemconfigs

// DatabaseConfig 数据库相关配置
type DatabaseConfig struct {
	ServerAccessLog struct {
		Clean struct {
			Days int `json:"days"` // 日志保留天数，0表示不限制
		} `json:"clean"` // 清理相关配置
	} `json:"serverAccessLog"` // 服务访问日志相关配置
}
