// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package reporterconfigs

// GlobalSetting 全局设置
type GlobalSetting struct {
	MinNotifyConnectivity float64 `json:"minNotifyConnectivity"` // 需要通知的最小连通值
	NotifyWebHookURL      string  `json:"notifyWebHookURL"`      // WebHook通知地址
}

func DefaultGlobalSetting() *GlobalSetting {
	return &GlobalSetting{
		MinNotifyConnectivity: 100,
	}
}
