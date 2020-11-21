package filterconfigs

import "net/url"

type URLEncodeFilter struct {
}

func (this *URLEncodeFilter) Init() error {
	return nil
}

func (this *URLEncodeFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	output = url.QueryEscape(ToString(input))
	goNext = true
	return
}
