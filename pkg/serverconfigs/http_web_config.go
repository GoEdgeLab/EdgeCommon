package serverconfigs

import (
	"context"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
)

type HTTPWebConfig struct {
	Id                 int64                               `yaml:"id" json:"id"`                                 // ID
	IsOn               bool                                `yaml:"isOn" json:"isOn"`                             // 是否启用
	Locations          []*HTTPLocationConfig               `yaml:"locations" json:"locations"`                   // 路径规则 TODO
	LocationRefs       []*HTTPLocationRef                  `yaml:"locationRefs" json:"locationRefs"`             // 路径规则应用
	Compression        *HTTPCompressionConfig              `yaml:"compression" json:"compression"`               // 压缩配置
	Optimization       *HTTPPageOptimizationConfig         `yaml:"optimization" json:"optimization"`             // 页面优化配置
	WebP               *WebPImageConfig                    `yaml:"webp" json:"webp"`                             // WebP配置
	Charset            *HTTPCharsetConfig                  `yaml:"charset" json:"charset"`                       // 字符编码
	Shutdown           *HTTPShutdownConfig                 `yaml:"shutdown" json:"shutdown"`                     // 临时关闭配置
	Pages              []*HTTPPageConfig                   `yaml:"pages" json:"pages"`                           // 特殊页面配置
	RedirectToHttps    *HTTPRedirectToHTTPSConfig          `yaml:"redirectToHTTPS" json:"redirectToHTTPS"`       // 是否自动跳转到Https
	Root               *HTTPRootConfig                     `yaml:"root" json:"root"`                             // 资源根目录 TODO
	MaxRequestBodySize string                              `yaml:"maxRequestBodySize" json:"maxRequestBodySize"` // 请求body最大尺寸 TODO 需要实现
	AccessLogRef       *HTTPAccessLogRef                   `yaml:"accessLog" json:"accessLog"`                   // 访问日志配置
	StatRef            *HTTPStatRef                        `yaml:"statRef" json:"statRef"`                       // 统计配置
	Cache              *HTTPCacheConfig                    `yaml:"cache" json:"cache"`                           // 缓存配置
	FirewallRef        *firewallconfigs.HTTPFirewallRef    `yaml:"firewallRef" json:"firewallRef"`               // 防火墙设置
	FirewallPolicy     *firewallconfigs.HTTPFirewallPolicy `yaml:"firewallPolicy" json:"firewallPolicy"`         // 防火墙策略
	WebsocketRef       *HTTPWebsocketRef                   `yaml:"websocketRef" json:"websocketRef"`             // Websocket应用配置
	Websocket          *HTTPWebsocketConfig                `yaml:"websocket" json:"websocket"`                   // Websocket配置
	RewriteRefs        []*HTTPRewriteRef                   `yaml:"rewriteRefs" json:"rewriteRefs"`               // 重写规则配置
	RewriteRules       []*HTTPRewriteRule                  `yaml:"rewriteRules" json:"rewriteRules"`             // 重写规则
	FastcgiRef         *HTTPFastcgiRef                     `yaml:"fastcgiRef" json:"fastcgiRef"`                 // Fastcgi引用
	FastcgiList        []*HTTPFastcgiConfig                `yaml:"fastcgiList" json:"fastcgiList"`               // Fastcgi配置
	UserAgent          *UserAgentConfig                    `yaml:"userAgent" json:"userAgent"`                   // UserAgent配置

	RequestHeaderPolicyRef  *shared.HTTPHeaderPolicyRef `yaml:"requestHeaderPolicyRef" json:"requestHeaderPolicyRef"`   // 请求Header
	RequestHeaderPolicy     *shared.HTTPHeaderPolicy    `yaml:"requestHeaderPolicy" json:"requestHeaderPolicy"`         // 请求Header策略
	ResponseHeaderPolicyRef *shared.HTTPHeaderPolicyRef `yaml:"responseHeaderPolicyRef" json:"responseHeaderPolicyRef"` // 响应Header`
	ResponseHeaderPolicy    *shared.HTTPHeaderPolicy    `yaml:"responseHeaderPolicy" json:"responseHeaderPolicy"`       // 响应Header策略

	FilterRefs     []*HTTPFilterRef    `yaml:"filterRefs" json:"filterRefs"`         // 筛选配置 TODO
	FilterPolicies []*HTTPFilterPolicy `yaml:"filterPolicies" json:"filterPolicies"` // 筛选策略

	HostRedirects []*HTTPHostRedirectConfig `yaml:"hostRedirects" json:"hostRedirects"` // 主机跳转
	Auth          *HTTPAuthConfig           `yaml:"auth" json:"auth"`                   // 认证配置
	Referers      *ReferersConfig           `yaml:"referers" json:"referers"`           // 防盗链设置

	RemoteAddr   *HTTPRemoteAddrConfig `yaml:"remoteAddr" json:"remoteAddr"`     // 客户端IP获取方式
	MergeSlashes bool                  `yaml:"mergeSlashes" json:"mergeSlashes"` // 是否合并路径中的斜杠（/）

	RequestLimit   *HTTPRequestLimitConfig   `yaml:"requestLimit" json:"requestLimit"`     // 并发请求限制
	RequestScripts *HTTPRequestScriptsConfig `yaml:"requestScripts" json:"requestScripts"` // HTTP请求相关脚本

	// UAM, CC ...
	UAM *UAMConfig    `yaml:"uam" json:"uam"`
	CC  *HTTPCCConfig `yaml:"cc" json:"cc"`
}

func (this *HTTPWebConfig) Init(ctx context.Context) error {
	// root
	if this.Root != nil {
		err := this.Root.Init()
		if err != nil {
			return err
		}
	}

	// 路径规则
	if len(this.Locations) > 0 {
		for _, location := range this.Locations {
			err := location.Init(ctx)
			if err != nil {
				return err
			}
		}
	}

	// compression
	if this.Compression != nil {
		err := this.Compression.Init()
		if err != nil {
			return err
		}
	}

	// optimization
	if this.Optimization != nil {
		err := this.Optimization.Init()
		if err != nil {
			return err
		}
	}

	// charset
	if this.Charset != nil {
		err := this.Charset.Init()
		if err != nil {
			return err
		}
	}

	// shutdown
	if this.Shutdown != nil {
		err := this.Shutdown.Init()
		if err != nil {
			return err
		}
	}

	// pages
	if len(this.Pages) > 0 {
		for _, page := range this.Pages {
			err := page.Init()
			if err != nil {
				return err
			}
		}
	}

	// redirectToHTTPS
	if this.RedirectToHttps != nil {
		err := this.RedirectToHttps.Init()
		if err != nil {
			return err
		}
	}

	// accessLog
	if this.AccessLogRef != nil {
		err := this.AccessLogRef.Init()
		if err != nil {
			return err
		}
	}

	// stat
	if this.StatRef != nil {
		err := this.StatRef.Init()
		if err != nil {
			return err
		}
	}

	// cache
	if this.Cache != nil {
		err := this.Cache.Init()
		if err != nil {
			return err
		}
	}

	// firewall
	if this.FirewallRef != nil {
		err := this.FirewallRef.Init()
		if err != nil {
			return err
		}
	}
	if this.FirewallPolicy != nil {
		err := this.FirewallPolicy.Init()
		if err != nil {
			return err
		}
	}

	// websocket
	if this.WebsocketRef != nil {
		err := this.WebsocketRef.Init()
		if err != nil {
			return err
		}
	}
	if this.Websocket != nil {
		err := this.Websocket.Init()
		if err != nil {
			return err
		}
	}

	// request header
	if this.RequestHeaderPolicyRef != nil {
		err := this.RequestHeaderPolicyRef.Init()
		if err != nil {
			return err
		}
	}
	if this.RequestHeaderPolicy != nil {
		err := this.RequestHeaderPolicy.Init()
		if err != nil {
			return err
		}
	}

	// response header
	if this.ResponseHeaderPolicyRef != nil {
		err := this.ResponseHeaderPolicyRef.Init()
		if err != nil {
			return err
		}
	}
	if this.ResponseHeaderPolicy != nil {
		err := this.ResponseHeaderPolicy.Init()
		if err != nil {
			return err
		}
	}

	// filters
	if this.FilterRefs != nil {
		for _, ref := range this.FilterRefs {
			err := ref.Init()
			if err != nil {
				return err
			}
		}
	}
	if this.FilterPolicies != nil {
		for _, policy := range this.FilterPolicies {
			err := policy.Init()
			if err != nil {
				return err
			}
		}
	}

	// rewrite rules
	for _, rewriteRule := range this.RewriteRules {
		err := rewriteRule.Init()
		if err != nil {
			return err
		}
	}

	// 主机跳转
	for _, redirect := range this.HostRedirects {
		err := redirect.Init()
		if err != nil {
			return err
		}
	}

	// fastcgi
	for _, fastcgi := range this.FastcgiList {
		err := fastcgi.Init()
		if err != nil {
			return err
		}
	}

	// auth
	if this.Auth != nil {
		err := this.Auth.Init()
		if err != nil {
			return err
		}
	}

	// webp
	if this.WebP != nil {
		err := this.WebP.Init()
		if err != nil {
			return err
		}
	}

	// remoteAddr
	if this.RemoteAddr != nil {
		err := this.RemoteAddr.Init()
		if err != nil {
			return err
		}
	}

	// request limit
	if this.RequestLimit != nil {
		err := this.RequestLimit.Init()
		if err != nil {
			return err
		}
	}

	// request script
	if this.RequestScripts != nil {
		err := this.RequestScripts.Init()
		if err != nil {
			return err
		}
	}

	// uam
	if this.UAM != nil {
		err := this.UAM.Init()
		if err != nil {
			return err
		}
	}

	// cc
	if this.CC != nil {
		err := this.CC.Init()
		if err != nil {
			return err
		}
	}

	// user agent
	if this.UserAgent != nil {
		err := this.UserAgent.Init()
		if err != nil {
			return err
		}
	}

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
