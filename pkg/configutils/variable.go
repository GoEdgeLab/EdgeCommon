package configutils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	stringutil "github.com/iwind/TeaGo/utils/string"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

// VariableHolder 变量信息存储类型
type VariableHolder struct {
	Param     string
	Modifiers []string
}
type VariableHolders = []any

var variableMapping = map[string][]any{} // source => [holder1, ...]
var variableLocker = &sync.RWMutex{}
var regexpNamedVariable = regexp.MustCompile(`\${[@\w.|-]+}`)

var stringBuilderPool = sync.Pool{
	New: func() any {
		return &strings.Builder{}
	},
}

// ParseVariables 分析变量
func ParseVariables(source string, replacer func(varName string) (value string)) string {
	if len(source) == 0 {
		return ""
	}

	variableLocker.RLock()
	holders, found := variableMapping[source]
	variableLocker.RUnlock()
	if !found {
		holders = ParseHolders(source)
		variableLocker.Lock()
		variableMapping[source] = holders
		variableLocker.Unlock()
	}

	// no variables
	if len(holders) == 0 {
		return source
	}

	// 只有一个占位时，我们快速返回
	if len(holders) == 1 {
		var h = holders[0]
		holder, ok := h.(VariableHolder)
		if ok {
			var value = replacer(holder.Param)
			if holder.Modifiers != nil {
				value = doStringModifiers(value, holder.Modifiers)
			}
			return replacer(value)
		}
		return source
	}

	// 多个占位时，使用Builder
	var builder = stringBuilderPool.Get().(*strings.Builder)
	builder.Reset()
	defer stringBuilderPool.Put(builder)
	for _, h := range holders {
		holder, ok := h.(VariableHolder)
		if ok {
			var value = replacer(holder.Param)
			if holder.Modifiers != nil {
				value = doStringModifiers(value, holder.Modifiers)
			}
			builder.WriteString(value)
		} else {
			builder.Write(h.([]byte))
		}
	}
	return builder.String()
}

func ParseVariablesError(source string, replacer func(varName string) (value string, err error)) (string, error) {
	var resultErr error
	var result = ParseVariables(source, func(varName string) (value string) {
		replacedValue, err := replacer(varName)
		if err != nil {
			resultErr = err
		}
		return replacedValue
	})
	return result, resultErr
}

// ParseVariablesFromHolders 从占位中分析变量
func ParseVariablesFromHolders(holders VariableHolders, replacer func(varName string) (value string)) string {
	// no variables
	if len(holders) == 0 {
		return ""
	}

	// replace
	var result = strings.Builder{}
	for _, h := range holders {
		holder, ok := h.(VariableHolder)
		if ok {
			var value = replacer(holder.Param)
			if holder.Modifiers != nil {
				value = doStringModifiers(value, holder.Modifiers)
			}
			result.WriteString(value)
		} else {
			result.Write(h.([]byte))
		}
	}
	return result.String()
}

// ParseHolders 分析占位
func ParseHolders(source string) (holders VariableHolders) {
	var indexes = regexpNamedVariable.FindAllStringIndex(source, -1)
	var before = 0
	for _, loc := range indexes {
		holders = append(holders, []byte(source[before:loc[0]]))
		var holder = source[loc[0]+2 : loc[1]-1]

		if strings.Contains(holder, "|") {
			var holderPieces = strings.Split(holder, "|")
			holders = append(holders, VariableHolder{
				Param:     holderPieces[0],
				Modifiers: holderPieces[1:],
			})
		} else {
			holders = append(holders, VariableHolder{
				Param:     holder,
				Modifiers: nil,
			})
		}
		before = loc[1]
	}
	if before < len(source) {
		holders = append(holders, []byte(source[before:]))
	}
	return holders
}

// HasVariables 判断是否有变量
func HasVariables(source string) bool {
	if len(source) == 0 {
		return false
	}
	return regexpNamedVariable.MatchString(source)
}

// 执行变量后的修饰符
func doStringModifiers(value string, modifiers []string) string {
	for _, modifier := range modifiers {
		switch modifier {
		case "urlEncode":
			value = url.QueryEscape(value)
		case "urlDecode":
			value2, err := url.QueryUnescape(value)
			if err == nil {
				value = value2
			}
		case "base64Encode":
			value = base64.StdEncoding.EncodeToString([]byte(value))
		case "base64Decode":
			value2, err := base64.StdEncoding.DecodeString(value)
			if err == nil {
				value = string(value2)
			}
		case "md5":
			value = stringutil.Md5(value)
		case "sha1":
			value = fmt.Sprintf("%x", sha1.Sum([]byte(value)))
		case "sha256":
			value = fmt.Sprintf("%x", sha256.Sum256([]byte(value)))
		case "toLowerCase":
			value = strings.ToLower(value)
		case "toUpperCase":
			value = strings.ToUpper(value)
		}
	}
	return value
}
