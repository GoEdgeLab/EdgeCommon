package shared

// 请求条件分组
type HTTPRequestCondGroup struct {
	IsOn        bool               `yaml:"isOn" json:"isOn"`               // 是否启用
	Connector   string             `yaml:"connector" json:"connector"`     // 条件之间的关系
	Conds       []*HTTPRequestCond `yaml:"conds" json:"conds"`             // 条件列表
	IsReverse   bool               `yaml:"isReverse" json:"isReverse"`     // 是否反向匹配
	Description string             `yaml:"description" json:"description"` // 说明
}

// 初始化
func (this *HTTPRequestCondGroup) Init() error {
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

func (this *HTTPRequestCondGroup) Match(formatter func(source string) string) bool {
	if !this.IsOn {
		return !this.IsReverse
	}
	for _, cond := range this.Conds {
		isMatched := cond.Match(formatter)
		if this.Connector == "or" && isMatched {
			return !this.IsReverse
		}
		if this.Connector == "and" && !isMatched {
			return this.IsReverse
		}
	}
	return !this.IsReverse
}
