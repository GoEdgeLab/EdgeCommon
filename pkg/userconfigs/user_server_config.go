// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

const (
	MaxCacheKeysPerTask int32 = 1000
	MaxCacheKeysPerDay  int32 = 10000
)

type HTTPCacheTaskConfig struct {
	MaxKeysPerTask int32 `yaml:"maxKeysPerTask" json:"maxKeysPerTask"`
	MaxKeysPerDay  int32 `yaml:"maxKeysPerDay" json:"maxKeysPerDay"`
}

func DefaultHTTPCacheTaskConfig() *HTTPCacheTaskConfig {
	return &HTTPCacheTaskConfig{
		MaxKeysPerTask: 0,
		MaxKeysPerDay:  0,
	}
}

// UserServerConfig 用户服务设置
type UserServerConfig struct {
	GroupId                  int64                `yaml:"groupId" json:"groupId"`                                   // 分组
	RequirePlan              bool                 `yaml:"requirePlan" json:"requirePlan"`                           // 必须使用套餐
	EnableStat               bool                 `yaml:"enableStat" json:"enableStat"`                             // 开启统计
	HTTPCacheTaskPurgeConfig *HTTPCacheTaskConfig `yaml:"httpCacheTaskPurgeConfig" json:"httpCacheTaskPurgeConfig"` // 缓存任务删除配置
	HTTPCacheTaskFetchConfig *HTTPCacheTaskConfig `yaml:"httpCacheTaskFetchConfig" json:"httpCacheTaskFetchConfig"` // 缓存任务预热配置
}

func DefaultUserServerConfig() *UserServerConfig {
	return &UserServerConfig{
		GroupId:                  0,
		RequirePlan:              false,
		EnableStat:               true,
		HTTPCacheTaskPurgeConfig: DefaultHTTPCacheTaskConfig(),
		HTTPCacheTaskFetchConfig: DefaultHTTPCacheTaskConfig(),
	}
}
