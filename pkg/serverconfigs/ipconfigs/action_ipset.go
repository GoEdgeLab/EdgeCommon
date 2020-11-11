package ipconfigs

import (
	"os/exec"
	"strconv"
	"time"
)

type IPSetAction struct {
	Exe     string `yaml:"exe" json:"exe"`         // 可执行文件位置
	SetName string `yaml:"setName" json:"setName"` // 集合名称，多个集合使用同一个名称
}

func (this *IPSetAction) Node() string {
	return "node"
}

func (this *IPSetAction) Run(itemConfig *IPItemConfig) error {
	exe := this.Exe
	if len(exe) == 0 {
		path, err := exec.LookPath("ipset")
		if err != nil {
			return err
		}
		exe = path
	}

	var timeout int64 = 0
	if itemConfig.ExpiredAt > 0 {
		timeout = itemConfig.ExpiredAt - time.Now().Unix()
		if timeout <= 0 {
			return nil
		}
	}

	ip := itemConfig.IPFrom
	if len(itemConfig.IPTo) > 0 {
		ip += "-" + itemConfig.IPTo
	}

	switch itemConfig.Action {
	case IPItemActionAdd:
		cmd := exec.Command(exe, "add", this.SetName, ip, "timeout", strconv.FormatInt(timeout, 10))
		_ = cmd.Run()
	case IPItemActionUpdate:
		{
			cmd := exec.Command(exe, "del", this.SetName, ip, "timeout", strconv.FormatInt(timeout, 10))
			_ = cmd.Run()
		}
		{
			cmd := exec.Command(exe, "add", this.SetName, ip, "timeout", strconv.FormatInt(timeout, 10))
			_ = cmd.Run()
		}
	case IPItemActionDelete:
		cmd := exec.Command(exe, "del", this.SetName, ip, "timeout", strconv.FormatInt(timeout, 10))
		_ = cmd.Run()
	}

	return nil
}
