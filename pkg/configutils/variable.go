package configutils

import (
	"regexp"
	"strings"
	"sync"
)

// VariableHolder 变量信息存储类型
type VariableHolder string
type VariableHolders = []interface{}

var variableMapping = map[string][]interface{}{} // source => [holder1, ...]
var variableLocker = sync.RWMutex{}
var regexpNamedVariable = regexp.MustCompile(`\${[@\w.-]+}`)

var stringBuilderPool = sync.Pool{
	New: func() interface{} {
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
			return replacer(string(holder))
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
			builder.WriteString(replacer(string(holder)))
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
	result := strings.Builder{}
	for _, h := range holders {
		holder, ok := h.(VariableHolder)
		if ok {
			result.WriteString(replacer(string(holder)))
		} else {
			result.Write(h.([]byte))
		}
	}
	return result.String()
}

// ParseHolders 分析占位
func ParseHolders(source string) (holders VariableHolders) {
	indexes := regexpNamedVariable.FindAllStringIndex(source, -1)
	before := 0
	for _, loc := range indexes {
		holders = append(holders, []byte(source[before:loc[0]]))
		holder := source[loc[0]+2 : loc[1]-1]
		holders = append(holders, VariableHolder(holder))
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
