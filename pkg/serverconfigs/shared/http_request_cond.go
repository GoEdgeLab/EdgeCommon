package shared

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/types"
	"github.com/iwind/TeaGo/utils/string"
	"net"
	"path/filepath"
	"regexp"
	"strings"
)

// HTTPRequestCond HTTP请求匹配条件定义
type HTTPRequestCond struct {
	Type      string `yaml:"type" json:"type"`           // 类型，在特殊条件时使用
	IsRequest bool   `yaml:"isRequest" json:"isRequest"` // 是否为请求的条件，用来区分在什么阶段执行

	// 要测试的字符串
	// 其中可以使用跟请求相关的参数，比如：
	// ${arg.name}, ${requestPath}
	Param string `yaml:"param" json:"param"`

	Operator          RequestCondOperator `yaml:"operator" json:"operator"`                   // 运算符
	Value             string              `yaml:"value" json:"value"`                         // 对比值
	IsReverse         bool                `yaml:"isReverse" json:"isReverse"`                 // 是否反向匹配
	IsCaseInsensitive bool                `yaml:"isCaseInsensitive" json:"isCaseInsensitive"` // 大小写是否敏感

	isInt   bool
	isFloat bool
	isIP    bool

	regValue   *regexp.Regexp
	floatValue float64
	ipValue    net.IP
	arrayValue []string
}

// Init 校验配置
func (this *HTTPRequestCond) Init() error {
	this.isInt = RegexpDigitNumber.MatchString(this.Value)
	this.isFloat = RegexpFloatNumber.MatchString(this.Value)

	if lists.ContainsString([]string{
		RequestCondOperatorRegexp,
		RequestCondOperatorNotRegexp,
	}, this.Operator) {
		var value = this.Value
		if this.IsCaseInsensitive && !strings.HasPrefix(this.Value, "(?i)") {
			value = "(?i)" + value
		}
		reg, err := regexp.Compile(value)
		if err != nil {
			return err
		}
		this.regValue = reg
	} else if lists.ContainsString([]string{
		RequestCondOperatorWildcardMatch,
		RequestCondOperatorWildcardNotMatch,
	}, this.Operator) {
		var pieces = strings.Split(this.Value, "*")
		for index, piece := range pieces {
			pieces[index] = regexp.QuoteMeta(piece)
		}
		var pattern = strings.Join(pieces, "(.*)")
		reg, err := regexp.Compile("(?i)" /** 大小写不敏感 **/ + "^" + pattern + "$")
		if err != nil {
			return err
		}
		this.regValue = reg
	} else if lists.ContainsString([]string{
		RequestCondOperatorEqFloat,
		RequestCondOperatorGtFloat,
		RequestCondOperatorGteFloat,
		RequestCondOperatorLtFloat,
		RequestCondOperatorLteFloat,
	}, this.Operator) {
		this.floatValue = types.Float64(this.Value)
	} else if lists.ContainsString([]string{
		RequestCondOperatorEqIP,
		RequestCondOperatorGtIP,
		RequestCondOperatorGteIP,
		RequestCondOperatorLtIP,
		RequestCondOperatorLteIP,
	}, this.Operator) {
		this.ipValue = net.ParseIP(this.Value)
		this.isIP = this.ipValue != nil

		if !this.isIP {
			return errors.New("value should be a valid ip")
		}
	} else if lists.ContainsString([]string{
		RequestCondOperatorIPRange,
	}, this.Operator) {
		if strings.Contains(this.Value, ",") {
			ipList := strings.SplitN(this.Value, ",", 2)
			ipString1 := strings.TrimSpace(ipList[0])
			ipString2 := strings.TrimSpace(ipList[1])

			if len(ipString1) > 0 {
				ip1 := net.ParseIP(ipString1)
				if ip1 == nil {
					return errors.New("start ip is invalid")
				}
			}

			if len(ipString2) > 0 {
				ip2 := net.ParseIP(ipString2)
				if ip2 == nil {
					return errors.New("end ip is invalid")
				}
			}
		} else if strings.Contains(this.Value, "/") {
			_, _, err := net.ParseCIDR(this.Value)
			if err != nil {
				return err
			}
		} else {
			return errors.New("invalid ip range")
		}
	} else if lists.ContainsString([]string{
		RequestCondOperatorIn,
		RequestCondOperatorNotIn,
		RequestCondOperatorFileExt,
	}, this.Operator) {
		stringsValue := []string{}
		err := json.Unmarshal([]byte(this.Value), &stringsValue)
		if err != nil {
			return err
		}
		this.arrayValue = stringsValue
	} else if lists.ContainsString([]string{
		RequestCondOperatorFileMimeType,
	}, this.Operator) {
		stringsValue := []string{}
		err := json.Unmarshal([]byte(this.Value), &stringsValue)
		if err != nil {
			return err
		}
		for k, v := range stringsValue {
			if strings.Contains(v, "*") {
				v = regexp.QuoteMeta(v)
				v = strings.Replace(v, `\*`, ".*", -1)
				stringsValue[k] = v
			}
		}
		this.arrayValue = stringsValue
	}
	return nil
}

// Match 将此条件应用于请求，检查是否匹配
func (this *HTTPRequestCond) Match(formatter func(source string) string) bool {
	b := this.match(formatter)
	if this.IsReverse {
		return !b
	}
	return b
}

func (this *HTTPRequestCond) match(formatter func(source string) string) bool {
	paramValue := formatter(this.Param)
	switch this.Operator {
	case RequestCondOperatorRegexp:
		if this.regValue == nil {
			return false
		}
		return this.regValue.MatchString(paramValue)
	case RequestCondOperatorNotRegexp:
		if this.regValue == nil {
			return false
		}
		return !this.regValue.MatchString(paramValue)
	case RequestCondOperatorWildcardMatch:
		if this.regValue == nil {
			return false
		}
		return this.regValue.MatchString(paramValue)
	case RequestCondOperatorWildcardNotMatch:
		if this.regValue == nil {
			return false
		}
		return !this.regValue.MatchString(paramValue)
	case RequestCondOperatorEqInt:
		return this.isInt && paramValue == this.Value
	case RequestCondOperatorEqFloat:
		return this.isFloat && types.Float64(paramValue) == this.floatValue
	case RequestCondOperatorGtFloat:
		return this.isFloat && types.Float64(paramValue) > this.floatValue
	case RequestCondOperatorGteFloat:
		return this.isFloat && types.Float64(paramValue) >= this.floatValue
	case RequestCondOperatorLtFloat:
		return this.isFloat && types.Float64(paramValue) < this.floatValue
	case RequestCondOperatorLteFloat:
		return this.isFloat && types.Float64(paramValue) <= this.floatValue
	case RequestCondOperatorMod:
		pieces := strings.SplitN(this.Value, ",", 2)
		if len(pieces) == 1 {
			rem := types.Int64(pieces[0])
			return types.Int64(paramValue)%10 == rem
		}
		div := types.Int64(pieces[0])
		if div == 0 {
			return false
		}
		rem := types.Int64(pieces[1])
		return types.Int64(paramValue)%div == rem
	case RequestCondOperatorMod10:
		return types.Int64(paramValue)%10 == types.Int64(this.Value)
	case RequestCondOperatorMod100:
		return types.Int64(paramValue)%100 == types.Int64(this.Value)
	case RequestCondOperatorEqString:
		if this.IsCaseInsensitive {
			return strings.EqualFold(paramValue, this.Value)
		}
		return paramValue == this.Value
	case RequestCondOperatorNeqString:
		if this.IsCaseInsensitive {
			return !strings.EqualFold(paramValue, this.Value)
		}
		return paramValue != this.Value
	case RequestCondOperatorHasPrefix:
		if this.IsCaseInsensitive {
			return strings.HasPrefix(strings.ToUpper(paramValue), strings.ToUpper(this.Value))
		}
		return strings.HasPrefix(paramValue, this.Value)
	case RequestCondOperatorHasSuffix:
		if this.IsCaseInsensitive {
			return strings.HasSuffix(strings.ToUpper(paramValue), strings.ToUpper(this.Value))
		}
		return strings.HasSuffix(paramValue, this.Value)
	case RequestCondOperatorContainsString:
		if this.IsCaseInsensitive {
			return strings.Contains(strings.ToUpper(paramValue), strings.ToUpper(this.Value))
		}
		return strings.Contains(paramValue, this.Value)
	case RequestCondOperatorNotContainsString:
		if this.IsCaseInsensitive {
			return !strings.Contains(strings.ToUpper(paramValue), strings.ToUpper(this.Value))
		}
		return !strings.Contains(paramValue, this.Value)
	case RequestCondOperatorEqIP:
		var ip = net.ParseIP(paramValue)
		if ip == nil {
			return false
		}
		return this.isIP && ip.Equal(this.ipValue)
	case RequestCondOperatorGtIP:
		ip := net.ParseIP(paramValue)
		if ip == nil {
			return false
		}
		return this.isIP && bytes.Compare(ip, this.ipValue) > 0
	case RequestCondOperatorGteIP:
		ip := net.ParseIP(paramValue)
		if ip == nil {
			return false
		}
		return this.isIP && bytes.Compare(ip, this.ipValue) >= 0
	case RequestCondOperatorLtIP:
		ip := net.ParseIP(paramValue)
		if ip == nil {
			return false
		}
		return this.isIP && bytes.Compare(ip, this.ipValue) < 0
	case RequestCondOperatorLteIP:
		ip := net.ParseIP(paramValue)
		if ip == nil {
			return false
		}
		return this.isIP && bytes.Compare(ip, this.ipValue) <= 0
	case RequestCondOperatorIPRange:
		ip := net.ParseIP(paramValue)
		if ip == nil {
			return false
		}

		// 检查IP范围格式
		if strings.Contains(this.Value, ",") {
			ipList := strings.SplitN(this.Value, ",", 2)
			ipString1 := strings.TrimSpace(ipList[0])
			ipString2 := strings.TrimSpace(ipList[1])

			if len(ipString1) > 0 {
				ip1 := net.ParseIP(ipString1)
				if ip1 == nil {
					return false
				}

				if bytes.Compare(ip, ip1) < 0 {
					return false
				}
			}

			if len(ipString2) > 0 {
				ip2 := net.ParseIP(ipString2)
				if ip2 == nil {
					return false
				}

				if bytes.Compare(ip, ip2) > 0 {
					return false
				}
			}

			return true
		} else if strings.Contains(this.Value, "/") {
			_, ipNet, err := net.ParseCIDR(this.Value)
			if err != nil {
				return false
			}
			return ipNet.Contains(ip)
		} else {
			return false
		}
	case RequestCondOperatorIn:
		if this.IsCaseInsensitive {
			paramValue = strings.ToUpper(paramValue)
			for _, v := range this.arrayValue {
				if strings.ToUpper(v) == paramValue {
					return true
				}
			}
			return false
		} else {
			return lists.ContainsString(this.arrayValue, paramValue)
		}
	case RequestCondOperatorNotIn:
		if this.IsCaseInsensitive {
			paramValue = strings.ToUpper(paramValue)
			for _, v := range this.arrayValue {
				if strings.ToUpper(v) == paramValue {
					return false
				}
			}
			return true
		} else {
			return !lists.ContainsString(this.arrayValue, paramValue)
		}
	case RequestCondOperatorFileExt:
		ext := filepath.Ext(paramValue)
		if len(ext) > 0 {
			ext = ext[1:] // remove dot
		}
		return lists.ContainsString(this.arrayValue, strings.ToLower(ext))
	case RequestCondOperatorFileMimeType:
		index := strings.Index(paramValue, ";")
		if index >= 0 {
			paramValue = strings.TrimSpace(paramValue[:index])
		}
		if len(this.arrayValue) == 0 {
			return false
		}
		for _, v := range this.arrayValue {
			if strings.Contains(v, "*") {
				reg, err := stringutil.RegexpCompile("^" + v + "$")
				if err == nil && reg.MatchString(paramValue) {
					return true
				}
			} else if paramValue == v {
				return true
			}
		}
	case RequestCondOperatorVersionRange:
		if strings.Contains(this.Value, ",") {
			versions := strings.SplitN(this.Value, ",", 2)
			version1 := strings.TrimSpace(versions[0])
			version2 := strings.TrimSpace(versions[1])
			if len(version1) > 0 && stringutil.VersionCompare(paramValue, version1) < 0 {
				return false
			}
			if len(version2) > 0 && stringutil.VersionCompare(paramValue, version2) > 0 {
				return false
			}
			return true
		} else {
			return stringutil.VersionCompare(paramValue, this.Value) >= 0
		}
	case RequestCondOperatorIPMod:
		pieces := strings.SplitN(this.Value, ",", 2)
		if len(pieces) == 1 {
			rem := types.Int64(pieces[0])
			return this.ipToInt64(net.ParseIP(paramValue))%10 == rem
		}
		div := types.Int64(pieces[0])
		if div == 0 {
			return false
		}
		rem := types.Int64(pieces[1])
		return this.ipToInt64(net.ParseIP(paramValue))%div == rem
	case RequestCondOperatorIPMod10:
		return this.ipToInt64(net.ParseIP(paramValue))%10 == types.Int64(this.Value)
	case RequestCondOperatorIPMod100:
		return this.ipToInt64(net.ParseIP(paramValue))%100 == types.Int64(this.Value)
		/**case RequestCondOperatorFileExist:
			index := strings.Index(paramValue, "?")
			if index > -1 {
				paramValue = paramValue[:index]
			}
			if len(paramValue) == 0 {
				return false
			}
			if !filepath.IsAbs(paramValue) {
				paramValue = Tea.Root + Tea.DS + paramValue
			}
			stat, err := os.Stat(paramValue)
			return err == nil && !stat.IsDir()
		case RequestCondOperatorFileNotExist:
			index := strings.Index(paramValue, "?")
			if index > -1 {
				paramValue = paramValue[:index]
			}
			if len(paramValue) == 0 {
				return true
			}
			if !filepath.IsAbs(paramValue) {
				paramValue = Tea.Root + Tea.DS + paramValue
			}
			stat, err := os.Stat(paramValue)
			return err != nil || stat.IsDir()**/
	}

	return false
}

func (this *HTTPRequestCond) ipToInt64(ip net.IP) int64 {
	if len(ip) == 0 {
		return 0
	}
	if len(ip) == 16 {
		return int64(binary.BigEndian.Uint32(ip[12:16]))
	}
	return int64(binary.BigEndian.Uint32(ip))
}
