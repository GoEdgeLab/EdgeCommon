// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package shared

type BodyType = string

const (
	BodyTypeURL  BodyType = "url"
	BodyTypeHTML BodyType = "html"
)

func FindAllBodyTypes() []*Definition {
	return []*Definition{
		{
			Name: "读取URL",
			Code: BodyTypeURL,
		},
		{
			Name: "HTML",
			Code: BodyTypeHTML,
		},
	}
}
