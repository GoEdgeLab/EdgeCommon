package filterconfigs

// 过滤接口
type FilterInterface interface {
	// 初始化
	Init() error

	// 执行过滤
	Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error)
}
