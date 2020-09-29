package shared

// 条件配置
// 数据结构：conds -> []groups -> []cond
type HTTPRequestCondsConfig struct {
	IsOn      bool                    `yaml:"isOn" json:"isOn"`
	Connector string                  `yaml:"connector" json:"connector"`
	Groups    []*HTTPRequestCondGroup `yaml:"groups" json:"groups"`
}

// 初始化
func (this *HTTPRequestCondsConfig) Init() error {
	if len(this.Connector) == 0 {
		this.Connector = "or"
	}

	for _, group := range this.Groups {
		err := group.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 判断请求是否匹配
func (this *HTTPRequestCondsConfig) MatchRequest(formatter func(s string) string) bool {
	if !this.IsOn && len(this.Groups) == 0 {
		return true
	}
	ok := false
	for _, group := range this.Groups {
		b := group.MatchRequest(formatter)
		if !b && this.Connector == "and" {
			return false
		}
		if b && this.Connector == "or" {
			return true
		}
		if b {
			// 对于 or 来说至少有一个分组要返回 true
			ok = true
		}
	}
	return ok
}

// 判断响应是否匹配
func (this *HTTPRequestCondsConfig) MatchResponse(formatter func(s string) string) bool {
	if !this.IsOn && len(this.Groups) == 0 {
		return true
	}
	ok := false
	for _, group := range this.Groups {
		b := group.MatchResponse(formatter)
		if !b && this.Connector == "and" {
			return false
		}
		if b && this.Connector == "or" {
			return true
		}
		if b {
			// 对于 or 来说至少有一个分组要返回 true
			ok = true
		}
	}
	return ok
}
