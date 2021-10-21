// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import timeutil "github.com/iwind/TeaGo/utils/time"

// BandwidthLimitStatus 带宽限制状态
type BandwidthLimitStatus struct {
	UntilDay string `yaml:"untilDay" json:"untilDay"` // 有效日期，格式YYYYMMDD
}

func (this *BandwidthLimitStatus) IsValid() bool {
	if len(this.UntilDay) == 0 {
		return false
	}
	return this.UntilDay >= timeutil.Format("Ymd")
}
