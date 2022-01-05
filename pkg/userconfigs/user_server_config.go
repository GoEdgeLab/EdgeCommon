// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

// UserServerConfig 用户服务设置
type UserServerConfig struct {
	GroupId     int64 `yaml:"groupId" json:"groupId"`         // 分组
	RequirePlan bool  `yaml:"requirePlan" json:"requirePlan"` // 必须使用套餐
	EnableStat  bool  `yaml:"enableStat" json:"enableStat"`   // 开启统计
}

func DefaultUserServerConfig() *UserServerConfig {
	return &UserServerConfig{
		GroupId:     0,
		RequirePlan: false,
		EnableStat:  true,
	}
}
