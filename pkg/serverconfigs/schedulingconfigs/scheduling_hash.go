package schedulingconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
	"hash/crc32"
)

// HashScheduling Hash调度算法
type HashScheduling struct {
	Scheduling

	count uint32
}

// Start 启动
func (this *HashScheduling) Start() {
	this.count = uint32(len(this.Candidates))
}

// Next 获取下一个候选对象
func (this *HashScheduling) Next(call *shared.RequestCall) CandidateInterface {
	if this.count == 0 {
		return nil
	}

	key := call.Options.GetString("key")

	if call.Formatter != nil {
		key = call.Formatter(key)
	}

	sum := crc32.ChecksumIEEE([]byte(key))
	return this.Candidates[sum%this.count]
}

// Summary 获取简要信息
func (this *HashScheduling) Summary() maps.Map {
	return maps.Map{
		"code":        "hash",
		"name":        "Hash算法",
		"description": "根据自定义的键值的Hash值分配源站",
		"networks":    []string{"http"},
	}
}
