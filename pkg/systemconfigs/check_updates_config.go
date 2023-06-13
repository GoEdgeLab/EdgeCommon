// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package systemconfigs

// CheckUpdatesConfig 检查更新配置
type CheckUpdatesConfig struct {
	AutoCheck      bool   `yaml:"autoCheck" json:"autoCheck"`           // 是否开启自动检查
	IgnoredVersion string `yaml:"ignoredVersion" json:"ignoredVersion"` // 上次忽略的版本
}

func NewCheckUpdatesConfig() *CheckUpdatesConfig {
	return &CheckUpdatesConfig{}
}
