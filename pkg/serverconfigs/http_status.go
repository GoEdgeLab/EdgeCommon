package serverconfigs

import "net/http"

type HTTPStatus struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func AllHTTPRedirectStatusList() []*HTTPStatus {
	return []*HTTPStatus{
		{
			Code: http.StatusMovedPermanently,
			Text: http.StatusText(http.StatusMovedPermanently),
		},
		{
			Code: http.StatusPermanentRedirect,
			Text: http.StatusText(http.StatusPermanentRedirect),
		},
		{
			Code: http.StatusFound,
			Text: http.StatusText(http.StatusFound),
		},
		{
			Code: http.StatusSeeOther,
			Text: http.StatusText(http.StatusSeeOther),
		},
		{
			Code: http.StatusTemporaryRedirect,
			Text: http.StatusText(http.StatusTemporaryRedirect),
		},
	}
}
