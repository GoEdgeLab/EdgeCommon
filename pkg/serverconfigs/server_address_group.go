package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"strings"
	"sync"
)

type ServerAddressGroup struct {
	fullAddr string
	servers  []*ServerConfig

	// 域名和服务映射
	strictDomainMap map[string]map[string]*ServerConfig // domain[:2] => {domain => *ServerConfig}
	fuzzyDomainMap  map[string]*ServerConfig            // special domain => *ServerConfig

	cacheLocker       sync.RWMutex
	cacheDomainMap    map[string]map[string]*ServerConfig // domain[:2] => {domain => *ServerConfig}
	countCacheDomains int

	// 支持CNAME的服务
	cnameDomainMap map[string]map[string]*ServerConfig // domain[:2] => {domain => *ServerConfig}

	// 第一个TLS Server
	firstTLSServer *ServerConfig
}

func NewServerAddressGroup(fullAddr string) *ServerAddressGroup {
	return &ServerAddressGroup{
		fullAddr:        fullAddr,
		strictDomainMap: map[string]map[string]*ServerConfig{},
		fuzzyDomainMap:  map[string]*ServerConfig{},
		cacheDomainMap:  map[string]map[string]*ServerConfig{},
		cnameDomainMap:  map[string]map[string]*ServerConfig{},
	}
}

// Add 添加服务
func (this *ServerAddressGroup) Add(server *ServerConfig) {
	for _, serverName := range server.AllStrictNames() {
		var prefix = this.domainPrefix(serverName)
		domainsMap, ok := this.strictDomainMap[prefix]
		if ok {
			domainsMap[serverName] = server
		} else {
			this.strictDomainMap[prefix] = map[string]*ServerConfig{serverName: server}
		}

		// CNAME
		if server.SupportCNAME {
			cnameDomainsMap, ok := this.cnameDomainMap[prefix]
			if ok {
				cnameDomainsMap[serverName] = server
			} else {
				this.cnameDomainMap[prefix] = map[string]*ServerConfig{serverName: server}
			}
		}
	}
	for _, serverName := range server.AllFuzzyNames() {
		this.fuzzyDomainMap[serverName] = server
	}

	this.servers = append(this.servers, server)

	// 第一个TLS Server
	if this.firstTLSServer == nil && server.SSLPolicy() != nil && server.SSLPolicy().IsOn {
		this.firstTLSServer = server
	}
}

// FullAddr 获取完整的地址
func (this *ServerAddressGroup) FullAddr() string {
	return this.fullAddr
}

// Protocol 获取当前分组的协议
func (this *ServerAddressGroup) Protocol() Protocol {
	for _, p := range AllProtocols() {
		if strings.HasPrefix(this.fullAddr, p.String()+":") {
			return p
		}
	}
	return ProtocolHTTP
}

// Addr 获取当前分组的地址
func (this *ServerAddressGroup) Addr() string {
	protocol := this.Protocol()
	if protocol == ProtocolUnix {
		return strings.TrimPrefix(this.fullAddr, protocol.String()+":")
	}
	return strings.TrimPrefix(this.fullAddr, protocol.String()+"://")
}

// Servers 读取所有服务
func (this *ServerAddressGroup) Servers() []*ServerConfig {
	return this.servers
}

// IsHTTP 判断当前分组是否为HTTP
func (this *ServerAddressGroup) IsHTTP() bool {
	p := this.Protocol()
	return p == ProtocolHTTP || p == ProtocolHTTP4 || p == ProtocolHTTP6
}

// IsHTTPS 判断当前分组是否为HTTPS
func (this *ServerAddressGroup) IsHTTPS() bool {
	p := this.Protocol()
	return p == ProtocolHTTPS || p == ProtocolHTTPS4 || p == ProtocolHTTPS6
}

// IsTCP 判断当前分组是否为TCP
func (this *ServerAddressGroup) IsTCP() bool {
	p := this.Protocol()
	return p == ProtocolTCP || p == ProtocolTCP4 || p == ProtocolTCP6
}

// IsTLS 判断当前分组是否为TLS
func (this *ServerAddressGroup) IsTLS() bool {
	p := this.Protocol()
	return p == ProtocolTLS || p == ProtocolTLS4 || p == ProtocolTLS6
}

// IsUnix 判断当前分组是否为Unix
func (this *ServerAddressGroup) IsUnix() bool {
	p := this.Protocol()
	return p == ProtocolUnix
}

// IsUDP 判断当前分组是否为UDP
func (this *ServerAddressGroup) IsUDP() bool {
	p := this.Protocol()
	return p == ProtocolUDP
}

// FirstServer 获取第一个Server
func (this *ServerAddressGroup) FirstServer() *ServerConfig {
	if len(this.servers) > 0 {
		return this.servers[0]
	}
	return nil
}

// FirstTLSServer 获取第一个TLS Server
func (this *ServerAddressGroup) FirstTLSServer() *ServerConfig {
	return this.firstTLSServer
}

// MatchServerName 使用域名查找服务
func (this *ServerAddressGroup) MatchServerName(serverName string) *ServerConfig {
	var prefix = this.domainPrefix(serverName)

	// 试图从缓存中读取
	this.cacheLocker.RLock()
	if len(this.cacheDomainMap) > 0 {
		domainMap, ok := this.cacheDomainMap[prefix]
		if ok {
			server, ok := domainMap[serverName]
			if ok {
				return server
			}
		}
	}
	this.cacheLocker.RUnlock()

	domainMap, ok := this.strictDomainMap[prefix]
	if ok {
		server, ok := domainMap[serverName]
		if ok {
			return server
		}
	}
	for pattern, server := range this.fuzzyDomainMap {
		if configutils.MatchDomain(pattern, serverName) {
			// 加入到缓存
			this.cacheLocker.Lock()

			// 限制缓存的最大尺寸，防止内存耗尽
			if this.countCacheDomains < 1_000_000 {
				domainMap, ok := this.cacheDomainMap[prefix]
				if ok {
					domainMap[serverName] = server
				} else {
					this.cacheDomainMap[prefix] = map[string]*ServerConfig{serverName: server}
				}
				this.countCacheDomains++
			}
			this.cacheLocker.Unlock()

			return server
		}
	}
	return nil
}

// MatchServerCNAME 使用CNAME查找服务
func (this *ServerAddressGroup) MatchServerCNAME(serverName string) *ServerConfig {
	var prefix = this.domainPrefix(serverName)

	domainMap, ok := this.cnameDomainMap[prefix]
	if ok {
		server, ok := domainMap[serverName]
		if ok {
			return server
		}
	}

	return nil
}

func (this *ServerAddressGroup) domainPrefix(domain string) string {
	if len(domain) < 2 {
		return domain
	}
	return domain[:2]
}
