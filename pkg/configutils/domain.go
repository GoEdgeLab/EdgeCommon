package configutils

import (
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/utils/string"
	"strings"
)

// MatchDomains 从一组规则中匹配域名
// 支持的格式：example.com, www.example.com, .example.com, *.example.com, ~(\d+).example.com
// 更多参考：http://nginx.org/en/docs/http/ngx_http_core_module.html#server_name
func MatchDomains(patterns []string, domain string) (isMatched bool) {
	if len(patterns) == 0 {
		return
	}
	for _, pattern := range patterns {
		if MatchDomain(pattern, domain) {
			return true
		}
	}
	return
}

// MatchDomain 匹配单个域名规则
func MatchDomain(pattern string, domain string) (isMatched bool) {
	if len(pattern) == 0 {
		return
	}

	if pattern == domain {
		return true
	}

	if pattern == "*" {
		return true
	}

	// 正则表达式
	if pattern[0] == '~' {
		reg, err := stringutil.RegexpCompile(strings.TrimSpace(pattern[1:]))
		if err != nil {
			logs.Error(err)
			return false
		}
		return reg.MatchString(domain)
	}

	if pattern[0] == '.' {
		return strings.HasSuffix(domain, pattern)
	}

	// 其他匹配
	var patternPieces = strings.Split(pattern, ".")
	var domainPieces = strings.Split(domain, ".")
	if len(patternPieces) != len(domainPieces) {
		return
	}
	isMatched = true
	for index, patternPiece := range patternPieces {
		if patternPiece == "" || patternPiece == "*" || patternPiece == domainPieces[index] {
			continue
		}
		if strings.HasSuffix(patternPiece, ":*") {
			var portIndex = strings.LastIndex(patternPiece, ":*")
			if portIndex >= 0 {
				var prefix = patternPiece[:portIndex]
				if strings.HasPrefix(domainPieces[index], prefix+":") {
					continue
				}
			}
		}
		isMatched = false
		break
	}
	return isMatched
}

// IsFuzzyDomain 判断是否为特殊域名
func IsFuzzyDomain(domain string) bool {
	if len(domain) == 0 {
		return true
	}
	if domain[0] == '.' || domain[0] == '~' {
		return true
	}
	for _, c := range domain {
		if c == '*' {
			return true
		}
	}
	return false
}
