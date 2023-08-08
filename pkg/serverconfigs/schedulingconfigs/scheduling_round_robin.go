package schedulingconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	"sync"
)

// RoundRobinScheduling 轮询调度算法
type RoundRobinScheduling struct {
	Scheduling

	rawWeights     []uint
	currentWeights []uint
	count          uint
	index          uint

	locker sync.Mutex
}

// Start 启动
func (this *RoundRobinScheduling) Start() {
	lists.Sort(this.Candidates, func(i int, j int) bool {
		c1 := this.Candidates[i]
		c2 := this.Candidates[j]
		return c1.CandidateWeight() > c2.CandidateWeight()
	})

	for _, c := range this.Candidates {
		weight := c.CandidateWeight()
		if weight == 0 {
			weight = 1
		} else if weight > 10000 {
			weight = 10000
		}
		this.rawWeights = append(this.rawWeights, weight)
	}

	this.currentWeights = append([]uint{}, this.rawWeights...)
	this.count = uint(len(this.Candidates))
}

// Next 获取下一个候选对象
func (this *RoundRobinScheduling) Next(call *shared.RequestCall) CandidateInterface {
	if this.count == 0 {
		return nil
	}
	this.locker.Lock()
	defer this.locker.Unlock()

	if this.index > this.count-1 {
		this.index = 0
	}
	var weight = this.currentWeights[this.index]

	// 已经一轮了，则重置状态
	if weight == 0 {
		if this.currentWeights[0] == 0 {
			this.currentWeights = append([]uint{}, this.rawWeights...)
		}
		this.index = 0
	}

	c := this.Candidates[this.index]
	this.currentWeights[this.index]--
	this.index++
	return c
}

// Summary 获取简要信息
func (this *RoundRobinScheduling) Summary() maps.Map {
	return maps.Map{
		"code":        "roundRobin",
		"name":        "RoundRobin轮询算法",
		"description": "根据权重，依次分配源站",
		"networks":    []string{"http", "tcp", "udp", "unix"},
	}
}
