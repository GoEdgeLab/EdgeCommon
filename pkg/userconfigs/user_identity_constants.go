// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package userconfigs

// 认证状态

type UserIdentityStatus = string

const (
	UserIdentityStatusNone      UserIdentityStatus = "none"
	UserIdentityStatusSubmitted UserIdentityStatus = "submitted"
	UserIdentityStatusRejected  UserIdentityStatus = "rejected"
	UserIdentityStatusVerified  UserIdentityStatus = "verified"
)

// 认证类型

type UserIdentityType = string

const (
	UserIdentityTypeIDCard            UserIdentityType = "idCard"
	UserIdentityTypeEnterpriseLicense UserIdentityType = "enterpriseLicense"
)

// 组织类型

type UserIdentityOrgType = string

const (
	UserIdentityOrgTypeEnterprise UserIdentityOrgType = "enterprise"
	UserIdentityOrgTypeIndividual UserIdentityOrgType = "individual"
)
