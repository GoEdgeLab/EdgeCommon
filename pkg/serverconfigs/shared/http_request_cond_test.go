package shared

import (
	"bytes"
	"fmt"
	"github.com/iwind/TeaGo/assert"
	"net"
	"regexp"
	"testing"
)

func TestRequestCond_Compare1(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		cond := HTTPRequestCond{
			Param:    "/hello",
			Operator: RequestCondOperatorRegexp,
			Value:    "abc",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello",
			Operator: RequestCondOperatorRegexp,
			Value:    "/\\w+",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/article/123.html",
			Operator: RequestCondOperatorRegexp,
			Value:    `^/article/\d+\.html$`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello",
			Operator: RequestCondOperatorRegexp,
			Value:    "[",
		}
		a.IsNotNil(cond.Init())
		a.IsFalse(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello",
			Operator: RequestCondOperatorNotRegexp,
			Value:    "abc",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello",
			Operator: RequestCondOperatorNotRegexp,
			Value:    "/\\w+",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		var cond = HTTPRequestCond{
			Param:    "/hello",
			Operator: RequestCondOperatorWildcardMatch,
			Value:    "/*",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		var cond = HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorWildcardMatch,
			Value:    "/*/world",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		var cond = HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorWildcardMatch,
			Value:    "/H*/world",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		var cond = HTTPRequestCond{
			Param:    "/hello",
			Operator: RequestCondOperatorWildcardMatch,
			Value:    "/hello/*",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		var cond = HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorWildcardNotMatch,
			Value:    "/hello/*",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		var cond = HTTPRequestCond{
			Param:    "/hello",
			Operator: RequestCondOperatorWildcardNotMatch,
			Value:    "/hello/*",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(format string) string {
			return format
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "123.123",
			Operator: RequestCondOperatorEqInt,
			Value:    "123",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "123",
			Operator: RequestCondOperatorEqInt,
			Value:    "123",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "abc",
			Operator: RequestCondOperatorEqInt,
			Value:    "abc",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "123",
			Operator: RequestCondOperatorEqFloat,
			Value:    "123",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "123.0",
			Operator: RequestCondOperatorEqFloat,
			Value:    "123",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "123.123",
			Operator: RequestCondOperatorEqFloat,
			Value:    "123.12",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "123",
			Operator: RequestCondOperatorGtFloat,
			Value:    "1",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "123",
			Operator: RequestCondOperatorGtFloat,
			Value:    "125",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "125",
			Operator: RequestCondOperatorGteFloat,
			Value:    "125",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "125",
			Operator: RequestCondOperatorLtFloat,
			Value:    "127",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "125",
			Operator: RequestCondOperatorLteFloat,
			Value:    "127",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "125",
			Operator: RequestCondOperatorEqString,
			Value:    "125",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "125",
			Operator: RequestCondOperatorNeqString,
			Value:    "125",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "125",
			Operator: RequestCondOperatorNeqString,
			Value:    "127",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorHasPrefix,
			Value:    "/hello",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorHasPrefix,
			Value:    "/hello2",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorHasSuffix,
			Value:    "world",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorHasSuffix,
			Value:    "world/",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorContainsString,
			Value:    "wo",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorContainsString,
			Value:    "wr",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorNotContainsString,
			Value:    "HELLO",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "/hello/world",
			Operator: RequestCondOperatorNotContainsString,
			Value:    "hello",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}
}

func TestRequestCond_IP(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		cond := HTTPRequestCond{
			Param:    "hello",
			Operator: RequestCondOperatorEqIP,
			Value:    "hello",
		}
		a.IsNotNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorEqIP,
			Value:    "hello",
		}
		a.IsNotNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorEqIP,
			Value:    "192.168.1.100",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorGtIP,
			Value:    "192.168.1.90",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorGteIP,
			Value:    "192.168.1.90",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.80",
			Operator: RequestCondOperatorLtIP,
			Value:    "192.168.1.90",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.0.100",
			Operator: RequestCondOperatorLteIP,
			Value:    "192.168.1.90",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.0.100",
			Operator: RequestCondOperatorIPRange,
			Value:    "192.168.0.90,",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.0.100",
			Operator: RequestCondOperatorIPRange,
			Value:    "192.168.0.90,192.168.1.100",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.0.100",
			Operator: RequestCondOperatorIPRange,
			Value:    ",192.168.1.100",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorIPRange,
			Value:    "192.168.0.90,192.168.1.99",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorIPRange,
			Value:    "192.168.0.90/24",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorIPRange,
			Value:    "192.168.0.90/18",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorIPRange,
			Value:    "a/18",
		}
		a.IsNotNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorIPMod10,
			Value:    "6",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorIPMod100,
			Value:    "76",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "192.168.1.100",
			Operator: RequestCondOperatorIPMod,
			Value:    "10,6",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}
}

func TestRequestCondIPCompare(t *testing.T) {
	{
		ip1 := net.ParseIP("192.168.3.100")
		ip2 := net.ParseIP("192.168.2.100")
		t.Log(bytes.Compare(ip1, ip2))
	}

	{
		ip1 := net.ParseIP("192.168.3.100")
		ip2 := net.ParseIP("a")
		t.Log(bytes.Compare(ip1, ip2))
	}

	{
		ip1 := net.ParseIP("b")
		ip2 := net.ParseIP("192.168.2.100")
		t.Log(bytes.Compare(ip1, ip2))
	}

	{
		ip1 := net.ParseIP("b")
		ip2 := net.ParseIP("a")
		t.Log(ip1 == nil)
		t.Log(bytes.Compare(ip1, ip2))
	}

	{
		cond := HTTPRequestCond{}
		t.Log(cond.ipToInt64(net.ParseIP("192.168.1.100")))
		t.Log(cond.ipToInt64(net.ParseIP("192.168.1.99")))
		t.Log(cond.ipToInt64(net.ParseIP("0.0.0.0")))
		t.Log(cond.ipToInt64(net.ParseIP("127.0.0.1")))
		t.Log(cond.ipToInt64(net.ParseIP("abc")))
		t.Log(cond.ipToInt64(net.ParseIP("192.168")))
		t.Log(cond.ipToInt64(net.ParseIP("2001:0db8:0000:0000:0000:ff00:0042:8329")))
	}
}

func TestRequestCond_In(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		cond := HTTPRequestCond{
			Param:    "a",
			Operator: RequestCondOperatorIn,
			Value:    `a`,
		}
		a.IsNotNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "a",
			Operator: RequestCondOperatorIn,
			Value:    `["a", "b"]`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "c",
			Operator: RequestCondOperatorNotIn,
			Value:    `["a", "b"]`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "a",
			Operator: RequestCondOperatorNotIn,
			Value:    `["a", "b"]`,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}
}

func TestRequestCond_File(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		cond := HTTPRequestCond{
			Param:    "a",
			Operator: RequestCondOperatorFileExt,
			Value:    `["jpeg", "jpg", "png"]`,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "a.gif",
			Operator: RequestCondOperatorFileExt,
			Value:    `["jpeg", "jpg", "png"]`,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "a.png",
			Operator: RequestCondOperatorFileExt,
			Value:    `["jpeg", "jpg", "png"]`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	/**{
		cond := HTTPRequestCond{
			Param:    "a.png",
			Operator: RequestCondOperatorFileExist,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    Tea.Root + "/README.md",
			Operator: RequestCondOperatorFileExist,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    Tea.Root + "/README.md?v=1",
			Operator: RequestCondOperatorFileExist,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    Tea.Root,
			Operator: RequestCondOperatorFileExist,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    Tea.Root,
			Operator: RequestCondOperatorFileExist,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "a.png",
			Operator: RequestCondOperatorFileNotExist,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    Tea.Root + "/README.md",
			Operator: RequestCondOperatorFileNotExist,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}**/
}

func TestRequestCond_MimeType(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		cond := HTTPRequestCond{
			Param:    "text/html; charset=utf-8",
			Operator: RequestCondOperatorFileMimeType,
			Value:    `["text/html"]`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "text/html; charset=utf-8",
			Operator: RequestCondOperatorFileMimeType,
			Value:    `["text/*"]`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "text/html; charset=utf-8",
			Operator: RequestCondOperatorFileMimeType,
			Value:    `["image/*"]`,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "text/plain; charset=utf-8",
			Operator: RequestCondOperatorFileMimeType,
			Value:    `["text/html", "image/jpeg", "image/png"]`,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}
}

func TestRequestCond_Version(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		cond := HTTPRequestCond{
			Param:    "1.0",
			Operator: RequestCondOperatorVersionRange,
			Value:    `1.0,1.1`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "1.0",
			Operator: RequestCondOperatorVersionRange,
			Value:    `1.0,`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "1.0",
			Operator: RequestCondOperatorVersionRange,
			Value:    `,1.1`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "0.9",
			Operator: RequestCondOperatorVersionRange,
			Value:    `1.0,1.1`,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "0.9",
			Operator: RequestCondOperatorVersionRange,
			Value:    `1.0`,
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "1.1",
			Operator: RequestCondOperatorVersionRange,
			Value:    `1.0`,
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}
}

func TestRequestCond_RegexpQuote(t *testing.T) {
	t.Log(regexp.QuoteMeta("a"))
	t.Log(regexp.QuoteMeta("*"))
	t.Log(regexp.QuoteMeta("([\\d]).*"))
}

func TestRequestCond_Mod(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		cond := HTTPRequestCond{
			Param:    "1",
			Operator: RequestCondOperatorMod,
			Value:    "1",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "1",
			Operator: RequestCondOperatorMod,
			Value:    "2",
		}
		a.IsNil(cond.Init())
		a.IsFalse(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "3",
			Operator: RequestCondOperatorMod,
			Value:    "3",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "1",
			Operator: RequestCondOperatorMod,
			Value:    "11,1",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "3",
			Operator: RequestCondOperatorMod,
			Value:    "11,3",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	{
		cond := HTTPRequestCond{
			Param:    "4",
			Operator: RequestCondOperatorMod,
			Value:    "2,0",
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	for i := 0; i < 100; i++ {
		cond := HTTPRequestCond{
			Param:    fmt.Sprintf("%d", i),
			Operator: RequestCondOperatorMod10,
			Value:    fmt.Sprintf("%d", i%10),
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}

	for i := 0; i < 2000; i++ {
		cond := HTTPRequestCond{
			Param:    fmt.Sprintf("%d", i),
			Operator: RequestCondOperatorMod100,
			Value:    fmt.Sprintf("%d", i%100),
		}
		a.IsNil(cond.Init())
		a.IsTrue(cond.Match(func(source string) string {
			return source
		}))
	}
}
