package shared

import (
	"regexp"
	"strings"
)

// MimeTypeRule mime type规则
type MimeTypeRule struct {
	Value string

	isAll  bool
	regexp *regexp.Regexp
}

func NewMimeTypeRule(mimeType string) (*MimeTypeRule, error) {
	mimeType = strings.ToLower(mimeType)

	var rule = &MimeTypeRule{
		Value: mimeType,
	}
	if mimeType == "*/*" || mimeType == "*" {
		rule.isAll = true
	} else if strings.Contains(mimeType, "*") {
		mimeType = strings.ReplaceAll(regexp.QuoteMeta(mimeType), `\*`, ".+")
		reg, err := regexp.Compile("^(?i)" + mimeType + "$")
		if err != nil {
			return nil, err
		}
		rule.regexp = reg
	}
	return rule, nil
}

func (this *MimeTypeRule) Match(mimeType string) bool {
	if this.isAll {
		return true
	}
	if this.regexp == nil {
		return this.Value == strings.ToLower(mimeType)
	}
	return this.regexp.MatchString(mimeType)
}
