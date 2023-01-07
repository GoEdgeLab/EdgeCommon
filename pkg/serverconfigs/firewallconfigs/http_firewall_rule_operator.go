package firewallconfigs

type HTTPFirewallRuleOperator = string
type HTTPFirewallRuleCaseInsensitive = string

const (
	HTTPFirewallRuleOperatorGt           HTTPFirewallRuleOperator = "gt"
	HTTPFirewallRuleOperatorGte          HTTPFirewallRuleOperator = "gte"
	HTTPFirewallRuleOperatorLt           HTTPFirewallRuleOperator = "lt"
	HTTPFirewallRuleOperatorLte          HTTPFirewallRuleOperator = "lte"
	HTTPFirewallRuleOperatorEq           HTTPFirewallRuleOperator = "eq"
	HTTPFirewallRuleOperatorNeq          HTTPFirewallRuleOperator = "neq"
	HTTPFirewallRuleOperatorEqString     HTTPFirewallRuleOperator = "eq string"
	HTTPFirewallRuleOperatorNeqString    HTTPFirewallRuleOperator = "neq string"
	HTTPFirewallRuleOperatorMatch        HTTPFirewallRuleOperator = "match"
	HTTPFirewallRuleOperatorNotMatch     HTTPFirewallRuleOperator = "not match"
	HTTPFirewallRuleOperatorContains     HTTPFirewallRuleOperator = "contains"
	HTTPFirewallRuleOperatorNotContains  HTTPFirewallRuleOperator = "not contains"
	HTTPFirewallRuleOperatorPrefix       HTTPFirewallRuleOperator = "prefix"
	HTTPFirewallRuleOperatorSuffix       HTTPFirewallRuleOperator = "suffix"
	HTTPFirewallRuleOperatorContainsAny  HTTPFirewallRuleOperator = "containsAny"
	HTTPFirewallRuleOperatorContainsAll  HTTPFirewallRuleOperator = "containsAll"
	HTTPFirewallRuleOperatorHasKey       HTTPFirewallRuleOperator = "has key" // has key in slice or map
	HTTPFirewallRuleOperatorVersionGt    HTTPFirewallRuleOperator = "version gt"
	HTTPFirewallRuleOperatorVersionLt    HTTPFirewallRuleOperator = "version lt"
	HTTPFirewallRuleOperatorVersionRange HTTPFirewallRuleOperator = "version range"

	HTTPFirewallRuleOperatorContainsBinary    HTTPFirewallRuleOperator = "contains binary"     // contains binary
	HTTPFirewallRuleOperatorNotContainsBinary HTTPFirewallRuleOperator = "not contains binary" // not contains binary

	// ip
	HTTPFirewallRuleOperatorEqIP       HTTPFirewallRuleOperator = "eq ip"
	HTTPFirewallRuleOperatorGtIP       HTTPFirewallRuleOperator = "gt ip"
	HTTPFirewallRuleOperatorGteIP      HTTPFirewallRuleOperator = "gte ip"
	HTTPFirewallRuleOperatorLtIP       HTTPFirewallRuleOperator = "lt ip"
	HTTPFirewallRuleOperatorLteIP      HTTPFirewallRuleOperator = "lte ip"
	HTTPFirewallRuleOperatorIPRange    HTTPFirewallRuleOperator = "ip range"
	HTTPFirewallRuleOperatorNotIPRange HTTPFirewallRuleOperator = "not ip range"
	HTTPFirewallRuleOperatorIPMod10    HTTPFirewallRuleOperator = "ip mod 10"
	HTTPFirewallRuleOperatorIPMod100   HTTPFirewallRuleOperator = "ip mod 100"
	HTTPFirewallRuleOperatorIPMod      HTTPFirewallRuleOperator = "ip mod"

	HTTPFirewallRuleCaseInsensitiveNone = "none"
	HTTPFirewallRuleCaseInsensitiveYes  = "yes"
	HTTPFirewallRuleCaseInsensitiveNo   = "no"
)

type RuleOperatorDefinition struct {
	Name            string                          `json:"name"`
	Code            string                          `json:"code"`
	Description     string                          `json:"description"`
	CaseInsensitive HTTPFirewallRuleCaseInsensitive `json:"caseInsensitive"` // default caseInsensitive setting
}

var AllRuleOperators = []*RuleOperatorDefinition{
	{
		Name:            "数值大于",
		Code:            HTTPFirewallRuleOperatorGt,
		Description:     "使用数值对比大于，对比值需要是一个数字",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值大于等于",
		Code:            HTTPFirewallRuleOperatorGte,
		Description:     "使用数值对比大于等于，对比值需要是一个数字",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值小于",
		Code:            HTTPFirewallRuleOperatorLt,
		Description:     "使用数值对比小于，对比值需要是一个数字",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值小于等于",
		Code:            HTTPFirewallRuleOperatorLte,
		Description:     "使用数值对比小于等于，对比值需要是一个数字",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值等于",
		Code:            HTTPFirewallRuleOperatorEq,
		Description:     "使用数值对比等于，对比值需要是一个数字",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值不等于",
		Code:            HTTPFirewallRuleOperatorNeq,
		Description:     "使用数值对比不等于，对比值需要是一个数字",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "字符串等于",
		Code:            HTTPFirewallRuleOperatorEqString,
		Description:     "使用字符串对比等于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "字符串不等于",
		Code:            HTTPFirewallRuleOperatorNeqString,
		Description:     "使用字符串对比不等于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "正则匹配",
		Code:            HTTPFirewallRuleOperatorMatch,
		Description:     "使用正则表达式匹配，在头部使用(?i)表示不区分大小写，<a href=\"https://goedge.cn/docs/Appendix/Regexp/Index.md\" target=\"_blank\">正则表达式语法 &raquo;</a>",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveYes,
	},
	{
		Name:            "正则不匹配",
		Code:            HTTPFirewallRuleOperatorNotMatch,
		Description:     "使用正则表达式不匹配，在头部使用(?i)表示不区分大小写，<a href=\"https://goedge.cn/docs/Appendix/Regexp/Index.md\" target=\"_blank\">正则表达式语法 &raquo;</a>",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveYes,
	},
	{
		Name:            "包含字符串",
		Code:            HTTPFirewallRuleOperatorContains,
		Description:     "包含某个字符串，比如Hello World包含了World",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "不包含字符串",
		Code:            HTTPFirewallRuleOperatorNotContains,
		Description:     "不包含某个字符串，比如Hello字符串中不包含Hi",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含任一字符串",
		Code:            HTTPFirewallRuleOperatorContainsAny,
		Description:     "包含字符串列表中的任意一个，比如/hello/world包含/hello和/hi中的/hello，每行一个字符串",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含所有字符串",
		Code:            HTTPFirewallRuleOperatorContainsAll,
		Description:     "包含字符串列表中的所有字符串，比如/hello/world必须包含/hello和/world，每行一个字符串",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含前缀",
		Code:            HTTPFirewallRuleOperatorPrefix,
		Description:     "包含字符串前缀部分，比如/hello前缀会匹配/hello, /hello/world等",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含后缀",
		Code:            HTTPFirewallRuleOperatorSuffix,
		Description:     "包含字符串后缀部分，比如/hello后缀会匹配/hello, /hi/hello等",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含二进制数据",
		Code:            HTTPFirewallRuleOperatorContainsBinary,
		Description:     "包含一组二进制数据",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "不包含二进制数据",
		Code:            HTTPFirewallRuleOperatorNotContainsBinary,
		Description:     "不包含一组二进制数据",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含索引",
		Code:            HTTPFirewallRuleOperatorHasKey,
		Description:     "对于一组数据拥有某个键值或者索引",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "版本号大于",
		Code:            HTTPFirewallRuleOperatorVersionGt,
		Description:     "对比版本号大于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "版本号小于",
		Code:            HTTPFirewallRuleOperatorVersionLt,
		Description:     "对比版本号小于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "版本号范围",
		Code:            HTTPFirewallRuleOperatorVersionRange,
		Description:     "判断版本号在某个范围内，格式为version1,version2",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP等于",
		Code:            HTTPFirewallRuleOperatorEqIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP大于",
		Code:            HTTPFirewallRuleOperatorGtIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP大于等于",
		Code:            HTTPFirewallRuleOperatorGteIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP小于",
		Code:            HTTPFirewallRuleOperatorLtIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP小于等于",
		Code:            HTTPFirewallRuleOperatorLteIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP范围",
		Code:            HTTPFirewallRuleOperatorIPRange,
		Description:     "IP在某个范围之内，范围格式可以是英文逗号分隔的<code-label>开始IP,结束IP</code-label>，比如<code-label>192.168.1.100,192.168.2.200</code-label>；或者CIDR格式的ip/bits，比如<code-label>192.168.2.1/24</code-label>；或者单个IP。可以填写多行，每行一个IP范围。",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "不在IP范围",
		Code:            HTTPFirewallRuleOperatorNotIPRange,
		Description:     "IP不在某个范围之内，范围格式可以是英文逗号分隔的<code-label>开始IP,结束IP</code-label>，比如<code-label>192.168.1.100,192.168.2.200</code-label>；或者CIDR格式的ip/bits，比如<code-label>192.168.2.1/24</code-label>；或者单个IP。可以填写多行，每行一个IP范围。",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP取模10",
		Code:            HTTPFirewallRuleOperatorIPMod10,
		Description:     "对IP参数值取模，除数为10，对比值为余数",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP取模100",
		Code:            HTTPFirewallRuleOperatorIPMod100,
		Description:     "对IP参数值取模，除数为100，对比值为余数",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP取模",
		Code:            HTTPFirewallRuleOperatorIPMod,
		Description:     "对IP参数值取模，对比值格式为：除数,余数，比如10,1",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
}

func FindRuleOperatorName(code string) string {
	for _, operator := range AllRuleOperators {
		if operator.Code == code {
			return operator.Name
		}
	}
	return ""
}
