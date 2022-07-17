// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package userconfigs

type UserIdentityStatus = string

const (
	UserIdentityStatusNone      UserIdentityStatus = "none"
	UserIdentityStatusSubmitted UserIdentityStatus = "submitted"
	UserIdentityStatusRejected  UserIdentityStatus = "rejected"
	UserIdentityStatusVerified  UserIdentityStatus = "verified"
)

type UserIdentityType = string

const (
	UserIdentityTypeIDCard UserIdentityType = "idCard"
)
