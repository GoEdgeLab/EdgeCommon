package shared

import (
	"github.com/iwind/TeaGo/maps"
	"net/http"
)

// RequestCall 请求调用
type RequestCall struct {
	Formatter func(source string) string // 当前变量格式化函数
	Request   *http.Request              // 当前请求
	Domain    string                     // 当前域名

	ResponseCallbacks []func(resp http.ResponseWriter)
	Options           maps.Map
}

// NewRequestCall 获取新对象
func NewRequestCall() *RequestCall {
	return &RequestCall{
		Options: maps.Map{},
	}
}

// Reset 重置
func (this *RequestCall) Reset() {
	this.Formatter = nil
	this.Request = nil
	this.ResponseCallbacks = nil
	this.Options = maps.Map{}
}

// AddResponseCall 添加响应回调
func (this *RequestCall) AddResponseCall(callback func(resp http.ResponseWriter)) {
	this.ResponseCallbacks = append(this.ResponseCallbacks, callback)
}

// CallResponseCallbacks 执行响应回调
func (this *RequestCall) CallResponseCallbacks(resp http.ResponseWriter) {
	for _, callback := range this.ResponseCallbacks {
		callback(resp)
	}
}
