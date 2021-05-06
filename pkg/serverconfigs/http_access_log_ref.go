package serverconfigs

// DefaultHTTPAccessLogRef 默认的访问日志配置
var DefaultHTTPAccessLogRef = NewHTTPAccessLogRef()

// HTTPAccessLogRef 代理访问日志配置
type HTTPAccessLogRef struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"` // 是否覆盖
	IsOn    bool `yaml:"isOn" json:"isOn"`       // 是否启用

	Fields []int `yaml:"fields" json:"fields"` // 记录的字段

	Status1 bool `yaml:"status1" json:"status1"` // 1xx
	Status2 bool `yaml:"status2" json:"status2"` // 2xx
	Status3 bool `yaml:"status3" json:"status3"` // 3xx
	Status4 bool `yaml:"status4" json:"status4"` // 4xx
	Status5 bool `yaml:"status5" json:"status5"` // 5xx

	StorageOnly     bool    `yaml:"storageOnly" json:"storageOnly"`         // 是否只输出到存储策略
	StoragePolicies []int64 `yaml:"storagePolicies" json:"storagePolicies"` // 存储策略Ids

	FirewallOnly bool `yaml:"firewallOnly" json:"firewallOnly"` // 是否只记录防火墙相关日志
}

// NewHTTPAccessLogRef 获取新对象
func NewHTTPAccessLogRef() *HTTPAccessLogRef {
	return &HTTPAccessLogRef{
		IsOn:    false,
		Fields:  []int{},
		Status1: true,
		Status2: true,
		Status3: true,
		Status4: true,
		Status5: true,
	}
}

// Init 校验
func (this *HTTPAccessLogRef) Init() error {
	return nil
}

// Match 判断是否应该记录
func (this *HTTPAccessLogRef) Match(status int) bool {
	s := status / 100
	switch s {
	case 1:
		if !this.Status1 {
			return false
		}
	case 2:
		if !this.Status2 {
			return false
		}
	case 3:
		if !this.Status3 {
			return false
		}
	case 4:
		if !this.Status4 {
			return false
		}
	case 5:
		if !this.Status5 {
			return false
		}
	}

	return true
}

// ContainsStoragePolicy 是否包含某个存储策略
func (this *HTTPAccessLogRef) ContainsStoragePolicy(storagePolicyId int64) bool {
	for _, s := range this.StoragePolicies {
		if s == storagePolicyId {
			return true
		}
	}
	return false
}

// ContainsField 检查是否包含某个Field
func (this *HTTPAccessLogRef) ContainsField(field int) bool {
	for _, f := range this.Fields {
		if f == field {
			return true
		}
	}
	return false
}
