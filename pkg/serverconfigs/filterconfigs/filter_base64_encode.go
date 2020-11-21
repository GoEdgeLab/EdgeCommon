package filterconfigs

import (
	"encoding/base64"
)

type Base64EncodeFilter struct {
}

func (this *Base64EncodeFilter) Init() error {
	return nil
}

func (this *Base64EncodeFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	data := ToBytes(input)
	output = base64.StdEncoding.EncodeToString(data)
	goNext = true
	return
}
