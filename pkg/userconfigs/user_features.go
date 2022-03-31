// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"

// UserFeatureCode 用户功能代号
type UserFeatureCode = string

const (
	UserFeatureCodeServerAccessLog     UserFeatureCode = "server.accessLog"
	UserFeatureCodeServerViewAccessLog UserFeatureCode = "server.viewAccessLog"
	UserFeatureCodePlan                UserFeatureCode = "plan"
	UserFeatureCodeScript              UserFeatureCode = "script"
	UserFeatureCodeServerWAF           UserFeatureCode = "server.waf"
	UserFeatureCodeServerUAM           UserFeatureCode = "server.uam"
	UserFeatureCodeServerWebP          UserFeatureCode = "server.webp"
	UserFeatureCodeFinance             UserFeatureCode = "finance"
)

// UserFeature 用户功能
type UserFeature struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (this *UserFeature) ToPB() *pb.UserFeature {
	return &pb.UserFeature{Name: this.Name, Code: this.Code, Description: this.Description}
}

// FindAllUserFeatures 所有功能列表
func FindAllUserFeatures() []*UserFeature {
	return []*UserFeature{
		{
			Name:        "记录访问日志",
			Code:        UserFeatureCodeServerAccessLog,
			Description: "用户可以开启服务的访问日志",
		},
		{
			Name:        "查看访问日志",
			Code:        UserFeatureCodeServerViewAccessLog,
			Description: "用户可以查看服务的访问日志",
		},
		{
			Name:        "转发访问日志",
			Code:        "server.accessLog.forward",
			Description: "用户可以配置访问日志转发到自定义的API",
		},
		{
			Name:        "TCP负载均衡",
			Code:        "server.tcp",
			Description: "用户可以添加TCP/TLS负载均衡服务",
		},
		{
			Name:        "自定义TCP负载均衡端口",
			Code:        "server.tcp.port",
			Description: "用户可以自定义TCP端口",
		},
		{
			Name:        "UDP负载均衡",
			Code:        "server.udp",
			Description: "用户可以添加UDP负载均衡服务",
		},
		{
			Name:        "自定义UDP负载均衡端口",
			Code:        "server.udp.port",
			Description: "用户可以自定义UDP端口",
		},
		{
			Name:        "开启WAF",
			Code:        UserFeatureCodeServerWAF,
			Description: "用户可以开启WAF功能并可以设置黑白名单等",
		},
		{
			Name:        "边缘脚本",
			Code:        UserFeatureCodeScript,
			Description: "用户可以在使用边缘脚本过滤请求",
		},
		{
			Name:        "5秒盾",
			Code:        UserFeatureCodeServerUAM,
			Description: "用户可以使用5秒盾全站防护功能",
		},
		{
			Name:        "WebP",
			Code:        UserFeatureCodeServerWebP,
			Description: "用户可以开启WebP自动转换功能",
		},
		{
			Name:        "费用账单",
			Code:        UserFeatureCodeFinance,
			Description: "开启费用账单相关功能",
		},
		{
			Name:        "套餐",
			Code:        UserFeatureCodePlan,
			Description: "用户可以购买和管理套餐",
		},
	}
}

// FindUserFeature 查询单个功能
func FindUserFeature(code string) *UserFeature {
	for _, feature := range FindAllUserFeatures() {
		if feature.Code == code {
			return feature
		}
	}
	return nil
}
