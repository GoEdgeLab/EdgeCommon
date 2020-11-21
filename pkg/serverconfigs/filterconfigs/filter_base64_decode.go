package filterconfigs

import (
	"encoding/base64"
)

type Base64DecodeFilter struct {
}

func (this *Base64DecodeFilter) Init() error {
	return nil
}

func (this *Base64DecodeFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	output, err = base64.StdEncoding.DecodeString(ToString(input))
	if err != nil {
		return
	}
	goNext = true
	return
}
