package serverconfigs

// TODO 需要实现
type OriginGroupConfig struct {
	Origins []*OriginConfig `yaml:"origins" json:"origins"` // 源站列表
}
