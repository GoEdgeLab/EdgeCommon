package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
)

// 日志存储策略
type HTTPAccessLogStoragePolicy struct {
	Id      int64                          `yaml:"id" json:"id"`
	Name    string                         `yaml:"name" json:"name"`
	IsOn    bool                           `yaml:"isOn" json:"isOn"`
	Type    string                         `yaml:"type" json:"type"`       // 存储类型
	Options maps.Map                       `yaml:"options" json:"options"` // 存储选项
	Conds   *shared.HTTPRequestCondsConfig `yaml:"conds" json:"conds"`     // 请求条件
}

// 校验
func (this *HTTPAccessLogStoragePolicy) Init() error {
	// cond
	if this.Conds != nil {
		err := this.Conds.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 匹配关键词
func (this *HTTPAccessLogStoragePolicy) MatchKeyword(keyword string) (matched bool, name string, tags []string) {
	if configutils.MatchKeyword(this.Name, keyword) || configutils.MatchKeyword(this.Type, keyword) {
		matched = true
		name = this.Name
		if len(this.Type) > 0 {
			tags = []string{"类型：" + this.Type}
		}
	}
	return
}
