package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/logs"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
)

// 日志存储策略
// 存储在configs/accesslog.storage.$id.conf
type HTTPAccessLogStoragePolicy struct {
	Id      int64                  `yaml:"id" json:"id"`
	Name    string                 `yaml:"name" json:"name"`
	IsOn    bool                   `yaml:"isOn" json:"isOn"`
	Type    string                 `yaml:"type" json:"type"`       // 存储类型
	Options map[string]interface{} `yaml:"options" json:"options"` // 存储选项
	Conds   []*shared.RequestCond  `yaml:"conds" json:"conds"`     // 请求条件
}

// 创建新策略
func NewHTTPAccessLogStoragePolicy() *HTTPAccessLogStoragePolicy {
	return &HTTPAccessLogStoragePolicy{
		IsOn: true,
	}
}

// 从文件中加载策略
func NewAccessLogStoragePolicyFromId(id string) *HTTPAccessLogStoragePolicy {
	filename := "accesslog.storage." + id + ".conf"
	data, err := ioutil.ReadFile(Tea.ConfigFile(filename))
	if err != nil {
		logs.Error(err)
		return nil
	}
	policy := NewHTTPAccessLogStoragePolicy()
	err = yaml.Unmarshal(data, policy)
	if err != nil {
		logs.Error(err)
		return nil
	}
	return policy
}

// 校验
func (this *HTTPAccessLogStoragePolicy) Init() error {
	// cond
	if len(this.Conds) > 0 {
		for _, cond := range this.Conds {
			err := cond.Init()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// 保存
func (this *HTTPAccessLogStoragePolicy) Save() error {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	data, err := yaml.Marshal(this)
	if err != nil {
		return err
	}

	filename := "accesslog.storage." + this.IdString() + ".conf"
	return ioutil.WriteFile(Tea.ConfigFile(filename), data, 0666)
}

// 删除
func (this *HTTPAccessLogStoragePolicy) Delete() error {
	filename := "accesslog.storage." + this.IdString() + ".conf"
	return os.Remove(Tea.ConfigFile(filename))
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

// 匹配条件
func (this *HTTPAccessLogStoragePolicy) MatchConds(formatter func(string) string) bool {
	if len(this.Conds) == 0 {
		return true
	}
	for _, cond := range this.Conds {
		if !cond.Match(formatter) {
			return false
		}
	}
	return true
}

// 将ID转换为字符串
func (this *HTTPAccessLogStoragePolicy) IdString() string {
	return strconv.FormatInt(this.Id, 10)
}
