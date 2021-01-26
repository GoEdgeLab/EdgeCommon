package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/schedulingconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/lists"
	"sync"
)

type RequestHostType = int8

const (
	RequestHostTypeProxyServer RequestHostType = 0
	RequestHostTypeOrigin      RequestHostType = 1
	RequestHostTypeCustomized  RequestHostType = 2
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

	StripPrefix     string          `yaml:"stripPrefix" json:"stripPrefix"`         // 去除URL前缀
	RequestHostType RequestHostType `yaml:"requestHostType" json:"requestHostType"` // 请求Host类型
	RequestHost     string          `yaml:"requestHost" json:"requestHost"`         // 请求Host，支持变量
	RequestURI      string          `yaml:"requestURI" json:"requestURI"`           // 请求URI，支持变量，如果同时定义了StripPrefix，则先执行StripPrefix

	AddHeaders []string `yaml:"addHeaders" json:"addHeaders"` // 自动添加的Header

	AutoFlush bool `yaml:"autoFlush" json:"autoFlush"` // 是否自动刷新缓冲区，在比如SSE（server-sent events）场景下很有用

	requestHostHasVariables bool
	requestURIHasVariables  bool

	hasPrimaryOrigins  bool
	hasBackupOrigins   bool
	schedulingIsBackup bool
	schedulingObject   schedulingconfigs.SchedulingInterface
	schedulingLocker   sync.Mutex

	addXRealIPHeader         bool
	addXForwardedForHeader   bool
	addForwardedHeader       bool
	addXForwardedByHeader    bool
	addXForwardedHostHeader  bool
	addXForwardedProtoHeader bool
}

// 初始化
func (this *ReverseProxyConfig) Init() error {
	this.requestHostHasVariables = configutils.HasVariables(this.RequestHost)
	this.requestURIHasVariables = configutils.HasVariables(this.RequestURI)

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

	// Header
	this.addXRealIPHeader = lists.ContainsString(this.AddHeaders, "X-Real-IP")
	this.addXForwardedForHeader = lists.ContainsString(this.AddHeaders, "X-Forwarded-For")
	this.addXForwardedByHeader = lists.ContainsString(this.AddHeaders, "X-Forwarded-By")
	this.addXForwardedHostHeader = lists.ContainsString(this.AddHeaders, "X-Forwarded-Host")
	this.addXForwardedProtoHeader = lists.ContainsString(this.AddHeaders, "X-Forwarded-Proto")

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

// 判断RequestHost是否有变量
func (this *ReverseProxyConfig) RequestHostHasVariables() bool {
	return this.requestHostHasVariables
}

// 判断RequestURI是否有变量
func (this *ReverseProxyConfig) RequestURIHasVariables() bool {
	return this.requestURIHasVariables
}

// 是否添加X-Real-IP
func (this *ReverseProxyConfig) ShouldAddXRealIPHeader() bool {
	return this.addXRealIPHeader
}

// 是否添加X-Forwarded-For
func (this *ReverseProxyConfig) ShouldAddXForwardedForHeader() bool {
	return this.addXForwardedForHeader
}

// 是否添加X-Forwarded-By
func (this *ReverseProxyConfig) ShouldAddXForwardedByHeader() bool {
	return this.addXForwardedByHeader
}

// 是否添加X-Forwarded-Host
func (this *ReverseProxyConfig) ShouldAddXForwardedHostHeader() bool {
	return this.addXForwardedHostHeader
}

// 是否添加X-Forwarded-Proto
func (this *ReverseProxyConfig) ShouldAddXForwardedProtoHeader() bool {
	return this.addXForwardedProtoHeader
}
