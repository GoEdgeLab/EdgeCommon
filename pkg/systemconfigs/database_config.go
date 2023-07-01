package systemconfigs

// DatabaseConfig 数据库相关配置
type DatabaseConfig struct {
	ServerAccessLog struct {
		Clean struct {
			Days int `json:"days"` // 日志保留天数，0表示不限制
		} `json:"clean"` // 清理相关配置
	} `json:"serverAccessLog"` // 服务访问日志相关配置

	HTTPCacheTask struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"httpCacheTask"` // 缓存任务

	NodeTrafficDailyStat struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"nodeTrafficDailyStat"`

	ServerBandwidthStat struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"serverBandwidthStat"`

	ServerDailyStat struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"serverDailyStat"`

	UserBandwidthStat struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"userBandwidthStat"`

	NodeClusterTrafficDailyStat struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"nodeClusterTrafficDailyStat"`

	NodeTrafficHourlyStat struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"nodeTrafficHourlyStat"`

	ServerDomainHourlyStat struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"serverDomainHourlyStat"`

	TrafficDailyStat struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"trafficDailyStat"`

	TrafficHourlyStat struct {
		Clean struct {
			Days int `json:"days"`
		} `json:"clean"`
	} `json:"trafficHourlyStat"`
}

func NewDatabaseConfig() *DatabaseConfig {
	var config = &DatabaseConfig{}
	config.ServerAccessLog.Clean.Days = 14
	config.HTTPCacheTask.Clean.Days = 30
	config.NodeTrafficDailyStat.Clean.Days = 32
	config.ServerBandwidthStat.Clean.Days = 100
	config.ServerDailyStat.Clean.Days = 60
	config.UserBandwidthStat.Clean.Days = 100
	config.NodeClusterTrafficDailyStat.Clean.Days = 30
	config.NodeTrafficHourlyStat.Clean.Days = 15
	config.ServerDomainHourlyStat.Clean.Days = 7
	config.TrafficDailyStat.Clean.Days = 30
	config.TrafficHourlyStat.Clean.Days = 15
	return config
}
