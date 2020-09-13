package shared

// 状态吗
type HTTPStatusConfig struct {
	Always bool  `yaml:"always" json:"always"`
	Codes  []int `yaml:"codes" json:"codes"`
}

func (this *HTTPStatusConfig) Init() error {
	// TODO
	return nil
}

func (this *HTTPStatusConfig) Match(statusCode int) bool {
	if this.Always {
		return true
	}
	if len(this.Codes) == 0 {
		return false
	}
	for _, c := range this.Codes {
		if c == statusCode {
			return true
		}
	}
	return false
}
