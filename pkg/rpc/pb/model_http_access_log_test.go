package pb

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestHTTPAccessLog_Marshal(t *testing.T) {
	data, err := proto.Marshal(&HTTPAccessLog{
		ServerId:   1,
		LocationId: 1,
		RewriteId:  1,
		NodeId:     1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(data), "bytes")

	accessLog := &HTTPAccessLog{}
	err = proto.Unmarshal(data, accessLog)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("accessLog:", accessLog)
}

func TestHTTPAccessLog_Memory(t *testing.T) {
	s := []*HTTPAccessLog{}
	for i := 0; i < 100000; i++ {
		s = append(s, &HTTPAccessLog{
			ServerId:        1,
			LocationId:      1,
			RewriteId:       1,
			NodeId:          1,
			RequestPath:     "/hello",
			RequestURI:      "/hello?name=lu&age=20",
			RequestMethod:   "POST",
			RequestFilename: "/hello.html",
			Header: map[string]*Strings{
				"User-Agent": {
					Values: []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"},
				},
			},
		})
	}
	time.Sleep(10 * time.Second)
}

func TestHTTPAccessLog_RequestId(t *testing.T) {
	u := time.Now().UnixNano()
	t.Logf("%d, %d", u, len(strconv.FormatInt(u, 10)))
}

func BenchmarkHTTPAccessLog_Proto_Marshal(b *testing.B) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(&HTTPAccessLog{
			ServerId:        1,
			LocationId:      1,
			RewriteId:       1,
			NodeId:          1,
			RequestPath:     "/hello",
			RequestURI:      "/hello?name=lu&age=20",
			RequestMethod:   "POST",
			RequestFilename: "/hello.html",
			Header: map[string]*Strings{
				"User-Agent": {
					Values: []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"},
				},
			},
		})
	}
}

func BenchmarkHTTPAccessLog_JSON_Marshal(b *testing.B) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&HTTPAccessLog{
			ServerId:        1,
			LocationId:      1,
			RewriteId:       1,
			NodeId:          1,
			RequestPath:     "/hello",
			RequestURI:      "/hello?name=lu&age=20",
			RequestMethod:   "POST",
			RequestFilename: "/hello.html",
			Header: map[string]*Strings{
				"User-Agent": {
					Values: []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"},
				},
			},
		})
	}
}

func BenchmarkHTTPAccessLog_JSON_Marshal_Map(b *testing.B) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < b.N; i++ {
		m := map[string]interface{}{
			"ServerId":        "1",
			"LocationId":      "1",
			"RewriteId":       "1",
			"NodeId":          "1",
			"RequestPath":     "/hello",
			"RequestURI":      "/hello?name=lu&age=20",
			"RequestMethod":   "POST",
			"RequestFilename": "/hello.html",
			"Header": map[string]interface{}{
				"User-Agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"},
			},
		}
		_, _ = json.Marshal(m)
	}
}
