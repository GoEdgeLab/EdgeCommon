package ipconfigs

type ActionInterface interface {
	// 运行节点
	Node() string

	// 执行对IP信息的处理
	Run(itemConfig *IPItemConfig) error
}
