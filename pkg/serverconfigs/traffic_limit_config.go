// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

// DefaultTrafficLimitNoticePageBody 达到流量限制时默认提示内容
const DefaultTrafficLimitNoticePageBody = `<!DOCTYPE html>
<html>
<head>
<title>Traffic Limit Exceeded Warning</title>
<body>

<h1>Traffic Limit Exceeded Warning</h1>
<p>The site traffic has exceeded the limit. Please contact with the site administrator.</p>
<address>Request ID: ${requestId}.</address>

</body>
</html>`

// TrafficLimitConfig 流量限制
type TrafficLimitConfig struct {
	IsOn bool `yaml:"isOn" json:"isOn"` // 是否启用

	DailySize   *shared.SizeCapacity `yaml:"dailySize" json:"dailySize"`     // 每日限制
	MonthlySize *shared.SizeCapacity `yaml:"monthlySize" json:"monthlySize"` // 每月限制
	TotalSize   *shared.SizeCapacity `yaml:"totalSize" json:"totalSize"`     // 总限制 TODO 需要实现

	NoticePageBody string `yaml:"noticePageBody" json:"noticePageBody"` // 超出限制时的提醒，支持请求变量
}

func (this *TrafficLimitConfig) Init() error {
	return nil
}

// DailyBytes 每天限制
// 不使用Init()来初始化数据，是为了让其他地方不经过Init()也能得到计算值
func (this *TrafficLimitConfig) DailyBytes() int64 {
	if this.DailySize != nil {
		return this.DailySize.Bytes()
	}
	return -1
}

// MonthlyBytes 每月限制
func (this *TrafficLimitConfig) MonthlyBytes() int64 {
	if this.MonthlySize != nil {
		return this.MonthlySize.Bytes()
	}
	return -1
}

// TotalBytes 总限制
func (this *TrafficLimitConfig) TotalBytes() int64 {
	if this.TotalSize != nil {
		return this.TotalSize.Bytes()
	}
	return -1
}

// IsEmpty 检查是否有限制值
func (this *TrafficLimitConfig) IsEmpty() bool {
	return !this.IsOn || (this.DailyBytes() <= 0 && this.MonthlyBytes() <= 0 && this.TotalBytes() <= 0)
}

func (this *TrafficLimitConfig) Equals(another *TrafficLimitConfig) bool {
	if another == nil {
		return false
	}

	if this.IsOn != another.IsOn {
		return false
	}

	if !this.equalCapacity(this.DailySize, another.DailySize) {
		return false
	}

	if !this.equalCapacity(this.MonthlySize, another.MonthlySize) {
		return false
	}

	if !this.equalCapacity(this.TotalSize, another.TotalSize) {
		return false
	}

	if this.NoticePageBody != another.NoticePageBody {
		return false
	}

	return true
}

func (this *TrafficLimitConfig) equalCapacity(size1 *shared.SizeCapacity, size2 *shared.SizeCapacity) bool {
	if size1 == size2 { // all are nil
		return true
	}

	if size1 != nil {
		if size2 == nil {
			return false
		}
		return size1.Bytes() == size2.Bytes()
	}

	return false
}
