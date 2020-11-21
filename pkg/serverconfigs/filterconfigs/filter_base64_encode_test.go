package filterconfigs

import (
	"encoding/base64"
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestBase64EncodeFilter_Do(t *testing.T) {
	a := assert.NewAssertion(t)

	filter := &Base64EncodeFilter{}
	t.Log(filter.Do("hello", nil))
	t.Log(filter.Do("=", nil))

	output, goNext, err := filter.Do("123456", nil)
	if err != nil {
		t.Fatal(err)
	}
	a.IsTrue(goNext)

	outputString := output.(string)
	result, err := base64.StdEncoding.DecodeString(outputString)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("origin:", string(result))
}
