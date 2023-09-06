// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import timeutil "github.com/iwind/TeaGo/utils/time"

// DefaultPlanExpireNoticePageBody 套餐过期时提示
const DefaultPlanExpireNoticePageBody = `<!DOCTYPE html>
<html>
<head>
<title>套餐已过期</title>
<body>

<h1>套餐已过期，请及时续费。</h1>
<p>Your server plan has been expired, please renew the plan.</p>
<address>Request ID: ${requestId}.</address>

</body>
</html>`

// UserPlanConfig 用户套餐配置
type UserPlanConfig struct {
	Id    int64  `yaml:"id" json:"id"`       // 用户套餐ID
	DayTo string `yaml:"dayTo" json:"dayTo"` // 有效期

	Plan *PlanConfig `yaml:"plan" json:"plan"`
}

// Init 初始化
func (this *UserPlanConfig) Init() error {
	if this.Plan != nil {
		err := this.Plan.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// IsAvailable 是否有效
func (this *UserPlanConfig) IsAvailable() bool {
	return this.DayTo >= timeutil.Format("Y-m-d")
}
