package serverconfigs

import "github.com/iwind/TeaGo/maps"

type HTTPAccessLogField = int

const (
	HTTPAccessLogFieldHeader       HTTPAccessLogField = 1
	HTTPAccessLogFieldSentHeader   HTTPAccessLogField = 2
	HTTPAccessLogFieldArg          HTTPAccessLogField = 3
	HTTPAccessLogFieldCookie       HTTPAccessLogField = 4
	HTTPAccessLogFieldExtend       HTTPAccessLogField = 5
	HTTPAccessLogFieldReferer      HTTPAccessLogField = 6
	HTTPAccessLogFieldUserAgent    HTTPAccessLogField = 7
	HTTPAccessLogFieldRequestBody  HTTPAccessLogField = 8
	HTTPAccessLogFieldResponseBody HTTPAccessLogField = 9
)

var HTTPAccessLogFieldsCodes = []HTTPAccessLogField{
	HTTPAccessLogFieldHeader,
	HTTPAccessLogFieldSentHeader,
	HTTPAccessLogFieldArg,
	HTTPAccessLogFieldCookie,
	HTTPAccessLogFieldExtend,
	HTTPAccessLogFieldReferer,
	HTTPAccessLogFieldUserAgent,
	HTTPAccessLogFieldRequestBody,
	HTTPAccessLogFieldResponseBody,
}

var HTTPAccessLogDefaultFieldsCodes = []HTTPAccessLogField{
	HTTPAccessLogFieldHeader,
	HTTPAccessLogFieldSentHeader,
	HTTPAccessLogFieldArg,
	HTTPAccessLogFieldCookie,
	HTTPAccessLogFieldExtend,
	HTTPAccessLogFieldReferer,
	HTTPAccessLogFieldUserAgent,
}

var HTTPAccessLogFields = []maps.Map{
	{
		"code": HTTPAccessLogFieldHeader,
		"name": "请求Header列表",
	},
	{
		"code": HTTPAccessLogFieldSentHeader,
		"name": "响应Header列表",
	},
	{
		"code": HTTPAccessLogFieldArg,
		"name": "参数列表",
	},
	{
		"code": HTTPAccessLogFieldCookie,
		"name": "Cookie列表",
	},
	{
		"code": HTTPAccessLogFieldExtend,
		"name": "位置和浏览器分析",
	},
	{
		"code": HTTPAccessLogFieldReferer,
		"name": "请求来源",
	},
	{
		"code": HTTPAccessLogFieldUserAgent,
		"name": "终端信息",
	},
	{
		"code": HTTPAccessLogFieldRequestBody,
		"name": "请求Body",
	},
	{
		"code": HTTPAccessLogFieldResponseBody,
		"name": "响应Body",
	},
}
