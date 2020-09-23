package serverconfigs

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"

type HTTPWebConfig struct {
	Id                 int64                      `yaml:"id" json:"id"`                                 // ID
	IsOn               bool                       `yaml:"isOn" json:"isOn"`                             // 是否启用
	Locations          []*HTTPLocationConfig      `yaml:"locations" json:"locations"`                   // 路径规则 TODO
	LocationRefs       []*HTTPLocationRef         `yaml:"locationRefs" json:"locationRefs"`             // 路径规则应用
	GzipRef            *HTTPGzipRef               `yaml:"gzipRef" json:"gzipRef"`                       // Gzip引用
	Gzip               *HTTPGzipConfig            `yaml:"gzip" json:"gzip"`                             // Gzip配置
	Charset            *HTTPCharsetConfig         `yaml:"charset" json:"charset"`                       // 字符编码
	Shutdown           *HTTPShutdownConfig        `yaml:"shutdown" json:"shutdown"`                     // 临时关闭配置
	Pages              []*HTTPPageConfig          `yaml:"pages" json:"pages"`                           // 特殊页面配置
	RedirectToHttps    *HTTPRedirectToHTTPSConfig `yaml:"redirectToHttps" json:"redirectToHttps"`       // 是否自动跳转到Https
	Root               string                     `yaml:"root" json:"root"`                             // 资源根目录 TODO
	Indexes            []string                   `yaml:"indexes" json:"indexes"`                       // 默认首页文件
	MaxRequestBodySize string                     `yaml:"maxRequestBodySize" json:"maxRequestBodySize"` // 请求body最大尺寸
	AccessLogRef       *HTTPAccessLogRef          `yaml:"accessLog" json:"accessLog"`                   // 访问日志配置
	StatRef            *HTTPStatRef               `yaml:"statRef" json:"statRef"`                       // 统计配置
	CacheRef           *HTTPCacheRef              `yaml:"cacheRef" json:"cacheRef"`                     // 缓存配置
	FirewallRef        *HTTPFirewallRef           `yaml:"firewallRef" json:"firewallRef"`               // 防火墙设置
	WebsocketRef       *HTTPWebsocketRef          `yaml:"websocketRef" json:"websocketRef"`             // Websocket应用配置
	Websocket          *HTTPWebsocketConfig       `yaml:"websocket" json:"websocket"`                   // Websocket配置

	RequestHeaderPolicyRef  *shared.HTTPHeaderPolicyRef `yaml:"requestHeaderPolicyRef" json:"requestHeaderPolicyRef"`   // 请求Header
	RequestHeaderPolicy     *shared.HTTPHeaderPolicy    `yaml:"requestHeaderPolicy" json:"requestHeaderPolicy"`         // 请求Header策略
	ResponseHeaderPolicyRef *shared.HTTPHeaderPolicyRef `yaml:"responseHeaderPolicyRef" json:"responseHeaderPolicyRef"` // 响应Header`
	ResponseHeaderPolicy    *shared.HTTPHeaderPolicy    `yaml:"responseHeaderPolicy" json:"responseHeaderPolicy"`       // 响应Header策略
}

func (this *HTTPWebConfig) Init() error {
	return nil
}

func (this *HTTPWebConfig) RemoveLocationRef(locationId int64) {
	this.LocationRefs = this.removeLocationRef(this.LocationRefs, locationId)
}

func (this *HTTPWebConfig) removeLocationRef(refs []*HTTPLocationRef, locationId int64) []*HTTPLocationRef {
	result := []*HTTPLocationRef{}
	for _, ref := range refs {
		if ref.LocationId == locationId {
			continue
		}

		ref.Children = this.removeLocationRef(ref.Children, locationId)
		result = append(result, ref)
	}
	return result
}
