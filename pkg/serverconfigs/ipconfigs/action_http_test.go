package ipconfigs

import (
	"net/http"
	"testing"
	"time"
)

func TestHTTPAction_Run(t *testing.T) {
	action := &HTTPAction{
		URL:       "http://127.0.0.1:1234/get?hello=world",
		Method:    http.MethodGet,
		Params:    map[string]string{"a": "b"},
		ParamName: "IP",
	}
	err := action.Run(&IPItemConfig{
		IPFrom:    "192.168.1.100",
		ExpiredAt: time.Now().Unix() + 3600,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
