package filterconfigs

import "net/url"

type URLDecodeFilter struct {
}

func (this *URLDecodeFilter) Init() error {
	return nil
}

func (this *URLDecodeFilter) Do(input interface{}, options interface{}) (output interface{}, goNext bool, err error) {
	output, err = url.QueryUnescape(ToString(input))
	if err != nil {
		return
	}
	goNext = true
	return
}
