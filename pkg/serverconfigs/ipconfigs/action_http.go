package ipconfigs

import (
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// TODO 支持自定义Header
type HTTPAction struct {
	URL       string               `yaml:"url" json:"url"`              // URL
	Method    string               `yaml:"method" json:"method"`        // 请求方法
	ParamName string               `yaml:"paramName" json:"paramsName"` // 参数名
	Params    map[string]string    `yaml:"params" json:"params"`        // 附加参数，在请求的时候一起提交
	Timeout   *shared.TimeDuration `yaml:"timeout" json:"timeout"`      // 超时时间
	Tries     int                  `yaml:"tries" json:"tries"`          // 失败尝试次数
}

func (this *HTTPAction) Node() string {
	return "api"
}

func (this *HTTPAction) Run(itemConfig *IPItemConfig) error {
	if itemConfig == nil {
		return errors.New("invalid ip item")
	}
	itemJSON, err := json.Marshal(itemConfig)
	if err != nil {
		return err
	}

	method := this.Method
	if len(method) == 0 {
		method = http.MethodGet
	}

	var body io.Reader = nil
	defaultParamName := "ip"
	apiURL := this.URL

	v := url.Values{}
	for paramName, paramValue := range this.Params {
		v[paramName] = []string{paramValue}
	}
	if len(this.ParamName) == 0 {
		v[defaultParamName] = []string{string(itemJSON)}
	} else {
		v[this.ParamName] = []string{string(itemJSON)}
	}

	if method != http.MethodGet {
		body = strings.NewReader(v.Encode())
	} else {
		if strings.Contains(apiURL, "?") {
			apiURL += "&"
		} else {
			apiURL += "?"
		}
		apiURL += v.Encode()
	}

	req, err := http.NewRequest(method, apiURL, body)
	if err != nil {
		return err
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	client := &http.Client{}
	if this.Timeout != nil {
		timeout := this.Timeout.Duration()
		if timeout > 0 {
			client.Timeout = timeout
		}
	}
	defer func() {
		client.CloseIdleConnections()
	}()

	tries := this.Tries
	if tries <= 0 {
		tries = 1
	}
	for i := 0; i < tries; i++ {
		resp, err := client.Do(req)
		if err == nil {
			_ = resp.Body.Close()
			return nil
		} else if i == tries-1 {
			return err
		}
	}

	return nil
}
