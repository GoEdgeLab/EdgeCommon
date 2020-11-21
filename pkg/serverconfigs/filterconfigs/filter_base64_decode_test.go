package filterconfigs

import (
	"encoding/base64"
	"testing"
)

func TestBase64DecodeFilter_Do(t *testing.T) {
	filter := &Base64DecodeFilter{}
	t.Log(filter.Do("123456", nil))

	encodedString := base64.StdEncoding.EncodeToString([]byte("hello"))
	t.Log(filter.Do(encodedString, nil))
}
