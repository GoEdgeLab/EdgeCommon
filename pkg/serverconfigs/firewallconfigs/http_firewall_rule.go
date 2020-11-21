package firewallconfigs

import (
	"regexp"
	"strings"
)

var namedParamReg = regexp.MustCompile(`^\${\s*(.+)\s*}$`)

type HTTPFirewallRule struct {
	Id                int64                  `yaml:"id" json:"id"`
	IsOn              bool                   `yaml:"isOn" json:"isOn"`
	Param             string                 `yaml:"param" json:"param"`
	ParamFilters      []*ParamFilter         `yaml:"paramFilters" json:"paramFilters"`
	Operator          string                 `yaml:"operator" json:"operator"`
	Value             string                 `yaml:"value" json:"value"`
	IsCaseInsensitive bool                   `yaml:"isCaseInsensitive" json:"isCaseInsensitive"`
	CheckpointOptions map[string]interface{} `yaml:"checkpointOptions" json:"checkpointOptions"`
	Description       string                 `yaml:"description" json:"description"`
}

func (this *HTTPFirewallRule) Init() error {
	// TODO 执行更严谨的校验
	return nil
}

func (this *HTTPFirewallRule) Prefix() string {
	result := namedParamReg.FindStringSubmatch(this.Param)
	if len(result) > 0 {
		param := result[1]
		pieces := strings.Split(param, ".")
		return pieces[0]
	}
	return this.Param
}
