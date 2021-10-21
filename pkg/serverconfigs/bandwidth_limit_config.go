// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// DefaultBandwidthLimitNoticePageBody 达到带宽限制时默认提示内容
const DefaultBandwidthLimitNoticePageBody = `<!DOCTYPE html>
<html>
<head>
<title>Bandwidth Limit Exceeded Warning/title>
<body>

The site bandwidth has exceeded the limit. Please contact with the site administrator.

</body>
</html>`

// BandwidthLimitConfig 带宽限制
type BandwidthLimitConfig struct {
	IsOn bool `yaml:"isOn" json:"isOn"` // 是否启用

	DailySize   *shared.SizeCapacity `yaml:"dailySize" json:"dailySize"`     // 每日限制
	MonthlySize *shared.SizeCapacity `yaml:"monthlySize" json:"monthlySize"` // 每月限制
	TotalSize   *shared.SizeCapacity `yaml:"totalSize" json:"totalSize"`     // 总限制 TODO 需要实现

	NoticePageBody string `yaml:"noticePageBody" json:"noticePageBody"` // 超出限制时的提醒，支持请求变量
}

// DailyBytes 每天限制
// 不使用Init()来初始化数据，是为了让其他地方不经过Init()也能得到计算值
func (this *BandwidthLimitConfig) DailyBytes() int64 {
	if this.DailySize != nil {
		return this.DailySize.Bytes()
	}
	return -1
}

// MonthlyBytes 每月限制
func (this *BandwidthLimitConfig) MonthlyBytes() int64 {
	if this.MonthlySize != nil {
		return this.MonthlySize.Bytes()
	}
	return -1
}

// TotalBytes 总限制
func (this *BandwidthLimitConfig) TotalBytes() int64 {
	if this.TotalSize != nil {
		return this.TotalSize.Bytes()
	}
	return -1
}

// IsEmpty 检查是否有限制值
func (this *BandwidthLimitConfig) IsEmpty() bool {
	return !this.IsOn || (this.DailyBytes() <= 0 && this.MonthlyBytes() <= 0 && this.TotalBytes() <= 0)
}
