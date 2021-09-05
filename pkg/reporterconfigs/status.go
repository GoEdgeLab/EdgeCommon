// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package reporterconfigs

type Status struct {
	IP               string `json:"ip"`
	OS               string `json:"os"`
	OSName           string `json:"osName"`
	Username         string `json:"username"`
	BuildVersion     string `json:"buildVersion"`     // 编译版本
	BuildVersionCode uint32 `json:"buildVersionCode"` // 版本数字
	UpdatedAt        int64  `json:"updatedAt"`        // 更新时间

	Location string `json:"location"` // 从IP查询到的Location
	ISP      string `json:"isp"`      // 从IP查询到的ISP
}
