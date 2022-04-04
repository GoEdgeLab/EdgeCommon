// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodeconfigs

type NodeLevel struct {
	Name        string `yaml:"name" json:"name"`
	Code        int    `yaml:"code" json:"code"`
	Description string `yaml:"description" json:"description"`
}

func FindAllNodeLevels() []*NodeLevel {
	return []*NodeLevel{
		{
			Name:        "边缘节点",
			Code:        1,
			Description: "普通的边缘节点。",
		},
		{
			Name:        "L2节点",
			Code:        2,
			Description: "特殊的边缘节点，同时负责同组上一级节点的回源。",
		},
	}
}

func FindNodeLevel(level int) *NodeLevel {
	level--

	var levels = FindAllNodeLevels()
	if level < 0 {
		return levels[0]
	}
	if level < len(levels) {
		return levels[level]
	}
	return levels[0]
}
