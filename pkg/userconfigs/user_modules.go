// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package userconfigs

type UserModule = string

const (
	UserModuleCDN      UserModule = "cdn"
	UserModuleAntiDDoS UserModule = "antiDDoS"
	UserModuleNS       UserModule = "ns"
)

var DefaultUserModules = []UserModule{UserModuleCDN}
