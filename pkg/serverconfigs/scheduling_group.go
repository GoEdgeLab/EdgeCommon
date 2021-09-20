// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/schedulingconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
)

// SchedulingGroup 负载均衡分组
type SchedulingGroup struct {
	Scheduling *SchedulingConfig `yaml:"scheduling" json:"scheduling"`

	PrimaryOrigins []*OriginConfig
	BackupOrigins  []*OriginConfig

	hasPrimaryOrigins  bool
	hasBackupOrigins   bool
	schedulingIsBackup bool
	schedulingObject   schedulingconfigs.SchedulingInterface
}

// Init 初始化
func (this *SchedulingGroup) Init() error {
	this.hasPrimaryOrigins = len(this.PrimaryOrigins) > 0
	this.hasBackupOrigins = len(this.BackupOrigins) > 0

	if this.Scheduling == nil {
		this.Scheduling = &SchedulingConfig{
			Code:    "random",
			Options: maps.Map{},
		}
	}

	return nil
}

// NextOrigin 取得下一个可用的后端服务
func (this *SchedulingGroup) NextOrigin(call *shared.RequestCall) *OriginConfig {
	if this.schedulingObject == nil {
		return nil
	}

	if this.Scheduling != nil && call != nil && call.Options != nil {
		for k, v := range this.Scheduling.Options {
			call.Options[k] = v
		}
	}

	candidate := this.schedulingObject.Next(call)

	// 末了重置状态
	defer func() {
		if candidate == nil {
			this.schedulingIsBackup = false
		}
	}()

	if candidate == nil {
		// 启用备用服务器
		if !this.schedulingIsBackup {
			this.SetupScheduling(true, true)
			candidate = this.schedulingObject.Next(call)
			if candidate == nil {
				// 不检查主要源站
				this.SetupScheduling(false, false)
				candidate = this.schedulingObject.Next(call)
				if candidate == nil {
					// 不检查备用源站
					this.SetupScheduling(true, false)
					candidate = this.schedulingObject.Next(call)
					if candidate == nil {
						return nil
					}
				}
			}
		}

		if candidate == nil {
			return nil
		}
	}

	return candidate.(*OriginConfig)
}

// SetupScheduling 设置调度算法
func (this *SchedulingGroup) SetupScheduling(isBackup bool, checkOk bool) {
	this.schedulingIsBackup = isBackup

	if this.Scheduling == nil {
		this.schedulingObject = &schedulingconfigs.RandomScheduling{}
	} else {
		typeCode := this.Scheduling.Code
		s := schedulingconfigs.FindSchedulingType(typeCode)
		if s == nil {
			this.Scheduling = nil
			this.schedulingObject = &schedulingconfigs.RandomScheduling{}
		} else {
			this.schedulingObject = s["instance"].(schedulingconfigs.SchedulingInterface)
		}
	}

	if !isBackup {
		for _, origin := range this.PrimaryOrigins {
			if origin.IsOn && (origin.IsOk || !checkOk) {
				this.schedulingObject.Add(origin)
			}
		}
	} else {
		for _, origin := range this.BackupOrigins {
			if origin.IsOn && (origin.IsOk || !checkOk) {
				this.schedulingObject.Add(origin)
			}
		}
	}

	if !this.schedulingObject.HasCandidates() {
		return
	}

	this.schedulingObject.Start()
}
