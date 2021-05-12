package shared

// HTTPRequestCondGroup 请求条件分组
type HTTPRequestCondGroup struct {
	IsOn        bool               `yaml:"isOn" json:"isOn"`               // 是否启用
	Connector   string             `yaml:"connector" json:"connector"`     // 条件之间的关系
	Conds       []*HTTPRequestCond `yaml:"conds" json:"conds"`             // 条件列表
	IsReverse   bool               `yaml:"isReverse" json:"isReverse"`     // 是否反向匹配
	Description string             `yaml:"description" json:"description"` // 说明

	requestConds  []*HTTPRequestCond
	responseConds []*HTTPRequestCond
}

// Init 初始化
func (this *HTTPRequestCondGroup) Init() error {
	if len(this.Connector) == 0 {
		this.Connector = "or"
	}

	this.requestConds = []*HTTPRequestCond{}
	this.responseConds = []*HTTPRequestCond{}
	if len(this.Conds) > 0 {
		for _, cond := range this.Conds {
			err := cond.Init()
			if err != nil {
				return err
			}

			if cond.IsRequest {
				this.requestConds = append(this.requestConds, cond)
			} else {
				this.responseConds = append(this.responseConds, cond)
			}
		}
	}
	return nil
}

func (this *HTTPRequestCondGroup) MatchRequest(formatter func(source string) string) bool {
	return this.match(this.requestConds, formatter)
}

func (this *HTTPRequestCondGroup) MatchResponse(formatter func(source string) string) bool {
	return this.match(this.responseConds, formatter)
}

func (this *HTTPRequestCondGroup) HasRequestConds() bool {
	return len(this.requestConds) > 0
}

func (this *HTTPRequestCondGroup) HasResponseConds() bool {
	return len(this.responseConds) > 0
}

func (this *HTTPRequestCondGroup) match(conds []*HTTPRequestCond, formatter func(source string) string) bool {
	if !this.IsOn || len(conds) == 0 {
		return !this.IsReverse
	}
	ok := false
	for _, cond := range conds {
		isMatched := cond.Match(formatter)
		if this.Connector == "or" && isMatched {
			return !this.IsReverse
		}
		if this.Connector == "and" && !isMatched {
			return this.IsReverse
		}
		if isMatched {
			// 对于OR来说至少要有一个返回true
			ok = true
		}
	}
	if this.IsReverse {
		return !ok
	}
	return ok
}
