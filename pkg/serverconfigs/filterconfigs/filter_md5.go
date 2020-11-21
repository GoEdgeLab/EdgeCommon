package filterconfigs

import (
	"crypto/md5"
	"fmt"
)

type Md5Filter struct {
}

func (this *Md5Filter) Init() error {
	return nil
}

func (this *Md5Filter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	data := ToBytes(input)
	m := md5.New()
	m.Write(data)
	result := m.Sum(nil)
	output = fmt.Sprintf("%x", result)
	goNext = true
	return
}
