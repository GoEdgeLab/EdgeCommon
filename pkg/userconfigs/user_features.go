// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"

// UserFeatureCode 用户功能代号
type UserFeatureCode = string

const (
	UserFeatureCodePlan UserFeatureCode = "plan"

	UserFeatureCodeServerTCP           UserFeatureCode = "server.tcp"
	UserFeatureCodeServerTCPPort       UserFeatureCode = "server.tcp.port"
	UserFeatureCodeServerUDP           UserFeatureCode = "server.udp"
	UserFeatureCodeServerUDPPort       UserFeatureCode = "server.udp.port"
	UserFeatureCodeServerAccessLog     UserFeatureCode = "server.accessLog"
	UserFeatureCodeServerViewAccessLog UserFeatureCode = "server.viewAccessLog"
	UserFeatureCodeScript              UserFeatureCode = "server.script"
	UserFeatureCodeServerWAF           UserFeatureCode = "server.waf"
	UserFeatureCodeServerOptimization  UserFeatureCode = "server.optimization"
	UserFeatureCodeServerUAM           UserFeatureCode = "server.uam"
	UserFeatureCodeServerWebP          UserFeatureCode = "server.webp"
	UserFeatureCodeServerCC            UserFeatureCode = "server.cc"
	UserFeatureCodeServerACME          UserFeatureCode = "server.acme"
	UserFeatureCodeServerAuth          UserFeatureCode = "server.auth"
	UserFeatureCodeServerWebsocket     UserFeatureCode = "server.websocket"
	UserFeatureCodeServerHTTP3         UserFeatureCode = "server.http3"
)

// UserFeature 用户功能
type UserFeature struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	SupportPlan bool   `json:"supportPlan"`
}

func (this *UserFeature) ToPB() *pb.UserFeature {
	return &pb.UserFeature{
		Name:        this.Name,
		Code:        this.Code,
		Description: this.Description,
		SupportPlan: this.SupportPlan,
	}
}

// FindAllUserFeatures 所有功能列表
func FindAllUserFeatures() []*UserFeature {
	return []*UserFeature{
		{
			Name:        "记录访问日志",
			Code:        UserFeatureCodeServerAccessLog,
			Description: "用户可以开启服务的访问日志。",
			SupportPlan: true,
		},
		{
			Name:        "查看访问日志",
			Code:        UserFeatureCodeServerViewAccessLog,
			Description: "用户可以查看服务的访问日志。",
			SupportPlan: true,
		},
		/**{
			Name:        "转发访问日志",
			Code:        "server.accessLog.forward",
			Description: "用户可以配置访问日志转发到自定义的API。",
			SupportPlan: true,
		},**/
		{
			Name:        "TCP负载均衡",
			Code:        UserFeatureCodeServerTCP,
			Description: "用户可以添加TCP/TLS负载均衡服务。",
			SupportPlan: true,
		},
		{
			Name:        "自定义TCP负载均衡端口",
			Code:        UserFeatureCodeServerTCPPort,
			Description: "用户可以自定义TCP端口。",
			SupportPlan: true,
		},
		{
			Name:        "UDP负载均衡",
			Code:        UserFeatureCodeServerUDP,
			Description: "用户可以添加UDP负载均衡服务。",
			SupportPlan: true,
		},
		{
			Name:        "自定义UDP负载均衡端口",
			Code:        UserFeatureCodeServerUDPPort,
			Description: "用户可以自定义UDP端口。",
			SupportPlan: true,
		},
		{
			Name:        "申请免费证书",
			Code:        UserFeatureCodeServerACME,
			Description: "用户可以申请ACME免费证书。",
			SupportPlan: false,
		},
		{
			Name:        "开启WAF",
			Code:        UserFeatureCodeServerWAF,
			Description: "用户可以开启WAF功能并可以设置黑白名单等。",
			SupportPlan: true,
		},
		{
			Name:        "边缘脚本",
			Code:        UserFeatureCodeScript,
			Description: "用户可以在使用边缘脚本过滤请求。",
			SupportPlan: true,
		},
		{
			Name:        "5秒盾",
			Code:        UserFeatureCodeServerUAM,
			Description: "用户可以使用5秒盾全站防护功能。",
			SupportPlan: true,
		},
		{
			Name:        "CC防护",
			Code:        UserFeatureCodeServerCC,
			Description: "用户可以使用CC防护功能。",
			SupportPlan: true,
		},
		{
			Name:        "WebP",
			Code:        UserFeatureCodeServerWebP,
			Description: "用户可以开启WebP自动转换功能。",
			SupportPlan: true,
		},
		{
			Name:        "页面优化",
			Code:        UserFeatureCodeServerOptimization,
			Description: "用户可以开启页面优化功能。",
			SupportPlan: true,
		},
		{
			Name:        "访问鉴权",
			Code:        UserFeatureCodeServerAuth,
			Description: "用户可以开启访问鉴权功能。",
			SupportPlan: true,
		},
		{
			Name:        "Websocket",
			Code:        UserFeatureCodeServerWebsocket,
			Description: "用户可以开启Websocket功能。",
			SupportPlan: true,
		},
		{
			Name:        "HTTP/3",
			Code:        UserFeatureCodeServerHTTP3,
			Description: "用户可以开启HTTP/3功能。",
			SupportPlan: true,
		},
		{
			Name:        "套餐",
			Code:        UserFeatureCodePlan,
			Description: "用户可以购买和管理套餐。",
			SupportPlan: false,
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

// CheckUserFeature 检查某个功能代号是否正确
func CheckUserFeature(featureCode string) bool {
	return FindUserFeature(featureCode) != nil
}
