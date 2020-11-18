package serverconfigs

import "github.com/iwind/TeaGo/maps"

const (
	DomainMismatchActionPage  = "page"
	DomainMismatchActionClose = "close"
)

type DomainMismatchPageOptions struct {
	StatusCode  int    `yaml:"statusCode" json:"statusCode"`
	ContentHTML string `yaml:"contentHTML" json:"contentHTML"`
}

type DomainMismatchCloseOptions struct {
}

type DomainMismatchAction struct {
	Code    string   `yaml:"code" json:"code"`       // 动作代号
	Options maps.Map `yaml:"options" json:"options"` // 动作选项
}

func (this *DomainMismatchAction) Init() error {
	return nil
}
