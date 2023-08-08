package configutils

import (
	"regexp"
	"strings"
)

var whitespaceReg = regexp.MustCompile(`\s+`)

// MatchKeyword 关键词匹配
func MatchKeyword(source, keyword string) bool {
	if len(keyword) == 0 {
		return false
	}

	pieces := whitespaceReg.Split(keyword, -1)
	source = strings.ToLower(source)
	for _, piece := range pieces {
		if strings.Contains(source, strings.ToLower(piece)) {
			return true
		}
	}

	return false
}
