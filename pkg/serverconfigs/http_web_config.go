package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type HTTPWebConfig struct {
	Id                 int64                      `yaml:"id" json:"id"`               // ID
	IsOn               bool                       `yaml:"isOn" json:"isOn"`           // 是否启用
	Locations          []*LocationConfig          `yaml:"locations" json:"locations"` // 路径规则 TODO
	GzipRef            *HTTPGzipRef               `yaml:"gzipRef" json:"gzipRef"`
	Gzip               *HTTPGzipConfig            `yaml:"gzip" json:"gzip"`                             // Gzip配置
	Charset            string                     `yaml:"charset" json:"charset"`                       // 字符编码
	Shutdown           *HTTPShutdownConfig        `yaml:"shutdown" json:"shutdown"`                     // 临时关闭配置
	Pages              []*HTTPPageConfig          `yaml:"pages" json:"pages"`                           // 特殊页面配置
	RedirectToHttps    *HTTPRedirectToHTTPSConfig `yaml:"redirectToHttps" json:"redirectToHttps"`       // 是否自动跳转到Https
	Root               string                     `yaml:"root" json:"root"`                             // 资源根目录 TODO
	Indexes            []string                   `yaml:"indexes" json:"indexes"`                       // 默认首页文件
	MaxRequestBodySize string                     `yaml:"maxRequestBodySize" json:"maxRequestBodySize"` // 请求body最大尺寸
	RequestHeaders     *shared.HTTPHeaderPolicy   `yaml:"requestHeaders" json:"requestHeaders"`         // 请求Header
	ResponseHeaders    *shared.HTTPHeaderPolicy   `yaml:"responseHeaders" json:"responseHeaders"`       // 响应Header`
	AccessLogRef       *HTTPAccessLogRef          `yaml:"accessLog" json:"accessLog"`                   // 访问日志配置
	StatRef            *HTTPStatRef               `yaml:"statRef" json:"statRef"`                       // 统计配置
	CacheRef           *HTTPCacheRef              `yaml:"cache" json:"cacheRef"`
}
