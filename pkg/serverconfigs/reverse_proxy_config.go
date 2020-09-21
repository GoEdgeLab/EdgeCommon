package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/schedulingconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"sync"
)

// 反向代理设置
type ReverseProxyConfig struct {
	Id                int64             `yaml:"id" json:"id"`                               // ID
	IsOn              bool              `yaml:"isOn" json:"isOn"`                           // 是否启用
	PrimaryOrigins    []*OriginConfig   `yaml:"primaryOrigins" json:"primaryOrigins"`       // 主要源站列表
	PrimaryOriginRefs []*OriginRef      `yaml:"primaryOriginRefs" json:"primaryOriginRefs"` // 主要源站引用
	BackupOrigins     []*OriginConfig   `yaml:"backupOrigins" json:"backupOrigins"`         // 备用源站列表
	BackupOriginRefs  []*OriginRef      `yaml:"backupOriginRefs" json:"backupOriginRefs"`   // 备用源站引用
	Scheduling        *SchedulingConfig `yaml:"scheduling" json:"scheduling"`               // 调度算法选项

	hasPrimaryOrigins  bool
	hasBackupOrigins   bool
	schedulingIsBackup bool
	schedulingObject   schedulingconfigs.SchedulingInterface
	schedulingLocker   sync.Mutex
}

// 初始化
func (this *ReverseProxyConfig) Init() error {
	this.hasPrimaryOrigins = len(this.PrimaryOrigins) > 0
	this.hasBackupOrigins = len(this.BackupOrigins) > 0

	for _, origin := range this.PrimaryOrigins {
		err := origin.Init()
		if err != nil {
			return err
		}
	}

	for _, origin := range this.BackupOrigins {
		err := origin.Init()
		if err != nil {
			return err
		}
	}

	// scheduling
	this.SetupScheduling(false)

	return nil
}

// 添加主源站配置
func (this *ReverseProxyConfig) AddPrimaryOrigin(origin *OriginConfig) {
	this.PrimaryOrigins = append(this.PrimaryOrigins, origin)
}

// 添加备用源站配置
func (this *ReverseProxyConfig) AddBackupOrigin(origin *OriginConfig) {
	this.BackupOrigins = append(this.BackupOrigins, origin)
}

// 取得下一个可用的后端服务
func (this *ReverseProxyConfig) NextOrigin(call *shared.RequestCall) *OriginConfig {
	this.schedulingLocker.Lock()
	defer this.schedulingLocker.Unlock()

	if this.schedulingObject == nil {
		return nil
	}

	if this.Scheduling != nil && call != nil && call.Options != nil {
		for k, v := range this.Scheduling.Options {
			call.Options[k] = v
		}
	}

	candidate := this.schedulingObject.Next(call)
	if candidate == nil {
		// 启用备用服务器
		if !this.schedulingIsBackup {
			this.SetupScheduling(true)

			candidate = this.schedulingObject.Next(call)
			if candidate == nil {
				return nil
			}
		}

		if candidate == nil {
			return nil
		}
	}

	return candidate.(*OriginConfig)
}

// 设置调度算法
func (this *ReverseProxyConfig) SetupScheduling(isBackup bool) {
	if !isBackup {
		this.schedulingLocker.Lock()
		defer this.schedulingLocker.Unlock()
	}
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
			if origin.IsOn {
				this.schedulingObject.Add(origin)
			}
		}
	} else {
		for _, origin := range this.BackupOrigins {
			if origin.IsOn {
				this.schedulingObject.Add(origin)
			}
		}
	}

	this.schedulingObject.Start()
}

// 获取调度配置对象
func (this *ReverseProxyConfig) FindSchedulingConfig() *SchedulingConfig {
	if this.Scheduling == nil {
		this.Scheduling = &SchedulingConfig{Code: "random"}
	}
	return this.Scheduling
}
