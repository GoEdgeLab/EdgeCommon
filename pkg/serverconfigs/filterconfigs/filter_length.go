package filterconfigs

type LengthFilter struct {
}

// 初始化
func (this *LengthFilter) Init() error {
	return nil
}

// 执行过滤
func (this *LengthFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	output = len(ToBytes(input))
	goNext = true
	return
}
