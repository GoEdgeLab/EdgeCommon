package firewallconfigs

import (
	"github.com/iwind/TeaGo/maps"
	"regexp"
)

// AllCheckpoints all check points list
var AllCheckpoints = []*HTTPFirewallCheckpointDefinition{
	{
		Name:        "通用请求报头长度限制",
		Prefix:      "requestGeneralHeaderLength",
		Description: "通用报头比如Cache-Control、Accept之类的长度限制，防止缓冲区溢出攻击。",
		IsRequest:   true,
		IsComposed:  true,
		Priority:    100,
	},
	{
		Name:        "通用响应报头长度限制",
		Prefix:      "responseGeneralHeaderLength",
		Description: "通用报头比如Cache-Control、Date之类的长度限制，防止缓冲区溢出攻击。",
		IsRequest:   false,
		IsComposed:  true,
		Priority:    100,
	},
	{
		Name:        "客户端地址（IP）",
		Prefix:      "remoteAddr",
		Description: "试图通过分析X-Forwarded-For等报头获取的客户端地址，比如192.168.1.100，存在伪造的可能。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "客户端源地址（IP）",
		Prefix:      "rawRemoteAddr",
		Description: "直接连接的客户端地址，比如192.168.1.100。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "客户端端口",
		Prefix:      "remotePort",
		Description: "直接连接的客户端地址端口。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "客户端用户名",
		Prefix:      "remoteUser",
		Description: "通过BasicAuth登录的客户端用户名。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "请求URI",
		Prefix:      "requestURI",
		Description: "包含URL参数的请求URI，类似于 /hello/world?lang=go，不包含域名部分。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "请求路径",
		Prefix:      "requestPath",
		Description: "不包含URL参数的请求路径，类似于 /hello/world，不包含域名部分。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "请求完整URL",
		Prefix:      "requestURL",
		Description: "完整的请求URL，包含协议、域名、请求路径、参数等，类似于 https://example.com/hello?name=lily 。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "请求内容长度",
		Prefix:      "requestLength",
		Description: "请求报头中的Content-Length。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "请求体内容",
		Prefix:      "requestBody",
		Description: "通常在POST或者PUT等操作时会附带请求体，最大限制32M。",
		IsRequest:   true,
		Priority:    5,
	},
	{
		Name:        "请求URI和请求体组合",
		Prefix:      "requestAll",
		Description: "${requestURI}和${requestBody}组合。",
		IsRequest:   true,
		Priority:    5,
	},
	{
		Name:        "请求表单参数",
		Prefix:      "requestForm",
		Description: "获取POST或者其他方法发送的表单参数，最大请求体限制32M。",
		IsRequest:   true,
		HasParams:   true,
		Priority:    5,
	},
	{
		Name:        "上传文件",
		Prefix:      "requestUpload",
		Description: "获取POST上传的文件信息，最大请求体限制32M。",
		Params: []*KeyValue{
			NewKeyValue("最小文件尺寸", "minSize"),
			NewKeyValue("最大文件尺寸", "maxSize"),
			NewKeyValue("扩展名(如.txt)", "ext"),
			NewKeyValue("原始文件名", "name"),
			NewKeyValue("表单字段名", "field"),
		},
		IsRequest: true,
		HasParams: true,
		Priority:  20,
	},
	{
		Name:        "请求JSON参数",
		Prefix:      "requestJSON",
		Description: "获取POST或者其他方法发送的JSON，最大请求体限制32M，使用点（.）符号表示多级数据。",
		IsRequest:   true,
		HasParams:   true,
		Priority:    5,
	},
	{
		Name:        "请求方法",
		Prefix:      "requestMethod",
		Description: "比如GET、POST。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "请求协议",
		Prefix:      "scheme",
		Description: "比如http或https。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "HTTP协议版本",
		Prefix:      "proto",
		Description: "比如HTTP/1.1。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "主机名",
		Prefix:      "host",
		Description: "比如example.com。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "CNAME",
		Prefix:      "cname",
		Description: "当前网站服务CNAME，比如38b48e4f.example.com。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "是否为CNAME",
		Prefix:      "isCNAME",
		Description: "是否为CNAME，值为1（是）或0（否）。",
		IsRequest:   true,
		Priority:    100,
		DataType:    "bool",
	},
	{
		Name:        "请求来源",
		Prefix:      "refererOrigin",
		Description: "请求报头中的Referer和Origin值。",
		IsRequest:   true,
		Priority:    100,
		Version:     "1.3.2",
	},
	{
		Name:        "请求来源Referer",
		Prefix:      "referer",
		Description: "请求报头中的Referer值。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "客户端信息",
		Prefix:      "userAgent",
		Description: "比如Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "内容类型",
		Prefix:      "contentType",
		Description: "请求报头的Content-Type。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "所有cookie组合字符串",
		Prefix:      "cookies",
		Description: "比如sid=IxZVPFhE&city=beijing&uid=18237。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "单个cookie值",
		Prefix:      "cookie",
		Description: "单个cookie值。",
		IsRequest:   true,
		HasParams:   true,
		Priority:    100,
	},
	{
		Name:        "所有URL参数组合",
		Prefix:      "args",
		Description: "比如name=lu&age=20。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "单个URL参数值",
		Prefix:      "arg",
		Description: "单个URL参数值。",
		IsRequest:   true,
		HasParams:   true,
		Priority:    100,
	},
	{
		Name:        "所有请求报头内容",
		Prefix:      "headers",
		Description: "使用换行符（\\n）隔开的报头内容字符串，每行均为\"NAME: VALUE格式\"。",
		IsRequest:   true,
		Priority:    100,
	},
	{
		Name:        "所有请求报头名称",
		Prefix:      "headerNames",
		Description: "使用换行符（\\n）隔开的报头名称字符串，每行一个名称。",
		IsRequest:   true,
		Priority:    100,
		Version:     "1.3.2",
	},
	{
		Name:        "单个请求报头值",
		Prefix:      "header",
		Description: "单个报头值。",
		IsRequest:   true,
		HasParams:   true,
		Priority:    100,
	},
	{
		Name:        "国家/地区名称",
		Prefix:      "geoCountryName",
		Description: "当前客户端所处国家/地区名称。",
		IsRequest:   true,
		HasParams:   false,
		Priority:    90,
	},
	{
		Name:        "省份名称",
		Prefix:      "geoProvinceName",
		Description: "当前客户端所处中国省份名称。",
		IsRequest:   true,
		HasParams:   false,
		Priority:    90,
	},
	{
		Name:        "城市名称",
		Prefix:      "geoCityName",
		Description: "当前客户端所处中国城市名称。",
		IsRequest:   true,
		HasParams:   false,
		Priority:    90,
	},
	{
		Name:        "ISP名称",
		Prefix:      "ispName",
		Description: "当前客户端所处ISP名称。",
		IsRequest:   true,
		HasParams:   false,
		Priority:    90,
	},
	{
		Name:        "CC统计",
		Prefix:      "cc2",
		Description: "对统计对象进行统计。",
		HasParams:   false,
		IsRequest:   true,
		IsComposed:  true,
		Priority:    10,
	},
	{
		Name:        "防盗链",
		Prefix:      "refererBlock",
		Description: "对统计对象进行统计。",
		HasParams:   false,
		IsRequest:   true,
		IsComposed:  true,
		Priority:    20,
	},
	{
		Name:        "CC统计（旧）",
		Prefix:      "cc",
		Description: "统计某段时间段内的请求信息（不推荐再使用，请使用新的CC2统计代替）。",
		HasParams:   true,
		Params: []*KeyValue{
			NewKeyValue("请求数", "requests"),
		},
		Options: []OptionInterface{
			&FieldOption{
				Type:        "field",
				Name:        "统计周期",
				Code:        "period",
				Value:       "60",
				IsRequired:  false,
				Size:        8,
				Comment:     "",
				Placeholder: "",
				RightLabel:  "秒",
				MaxLength:   8,
				Validate: func(value string) (ok bool, message string) {
					if regexp.MustCompile(`^\d+$`).MatchString(value) {
						ok = true
						return
					}
					message = "周期需要是一个整数数字"
					return
				},
			},
			&OptionsOption{
				Type:       "options",
				Name:       "用户识别读取来源",
				Code:       "userType",
				Value:      "",
				IsRequired: false,
				Size:       10,
				Comment:    "",
				RightLabel: "",
				Validate:   nil,
				Options: []maps.Map{
					{
						"name":  "IP",
						"value": "ip",
					},
					{
						"name":  "Cookie",
						"value": "cookie",
					},
					{
						"name":  "URL参数",
						"value": "get",
					},
					{
						"name":  "POST参数",
						"value": "post",
					},
					{
						"name":  "HTTP Header",
						"value": "header",
					},
				},
			},
			&FieldOption{
				Type:    "field",
				Name:    "用户识别字段",
				Code:    "userField",
				Comment: "识别用户的唯一性字段，在用户读取来源不是IP时使用",
			},
			&FieldOption{
				Type:      "field",
				Name:      "字段读取位置",
				Code:      "userIndex",
				Size:      5,
				MaxLength: 5,
				Comment:   "读取用户识别字段的位置，从0开始，比如user12345的数字ID 12345的位置就是5，在用户读取来源不是IP时使用",
			},
		},
		IsRequest: true,
		Priority:  10,
	},
	{
		Name:        "响应状态码",
		Prefix:      "status",
		Description: "响应状态码，比如200、404、500。",
		IsRequest:   false,
		Priority:    100,
	},
	{
		Name:        "响应报头",
		Prefix:      "responseHeader",
		Description: "响应报头值。",
		IsRequest:   false,
		HasParams:   true,
		Priority:    100,
	},
	{
		Name:        "响应内容",
		Prefix:      "responseBody",
		Description: "响应内容字符串。",
		IsRequest:   false,
		Priority:    5,
	},
	{
		Name:        "响应内容长度",
		Prefix:      "bytesSent",
		Description: "响应内容长度，通过响应的报头Content-Length获取。",
		IsRequest:   false,
		Priority:    100,
	},
}

// FindCheckpointDefinition 查找Checkpoint定义
func FindCheckpointDefinition(prefix string) *HTTPFirewallCheckpointDefinition {
	for _, checkpoint := range AllCheckpoints {
		if checkpoint.Prefix == prefix {
			return checkpoint
		}
	}
	return nil
}

// CheckCheckpointIsComposed 判断Checkpoint是否为组合的
func CheckCheckpointIsComposed(prefix string) bool {
	for _, checkpoint := range AllCheckpoints {
		if checkpoint.Prefix == prefix {
			return checkpoint.IsComposed
		}
	}
	return false
}
