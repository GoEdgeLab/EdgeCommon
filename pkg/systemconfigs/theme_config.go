// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package systemconfigs

func DefaultThemeBackgroundColors() []string {
	return []string{
		"14539A",
		"276AC6",
		"0081AF",
		"473BF0",
		"ACADBC",
		"9B9ECE",
		"C96480",
		"B47978",
		"B1AE91",
		"49A078",
		"46237A",
		"000500",
	}
}

// ThemeConfig 风格模板设置
type ThemeConfig struct {
	BackgroundColor string `yaml:"backgroundColor" json:"backgroundColor"` // 背景色，16进制，不需要带井号（#）
}
