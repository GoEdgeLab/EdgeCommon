package dao

import "github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"

var sharedRPCClient RPCClient

func SetRPC(client RPCClient) {
	sharedRPCClient = client
}

type RPCClient interface {
	SysSettingRPC() pb.SysSettingServiceClient
	NodeClusterRPC() pb.NodeClusterServiceClient
	NodeRegionRPC() pb.NodeRegionServiceClient
	NodePriceItemRPC() pb.NodePriceItemServiceClient
	ServerRPC() pb.ServerServiceClient
	ServerGroupRPC() pb.ServerGroupServiceClient
	OriginRPC() pb.OriginServiceClient
	HTTPWebRPC() pb.HTTPWebServiceClient
	ReverseProxyRPC() pb.ReverseProxyServiceClient
	HTTPGzipRPC() pb.HTTPGzipServiceClient
	HTTPHeaderRPC() pb.HTTPHeaderServiceClient
	HTTPHeaderPolicyRPC() pb.HTTPHeaderPolicyServiceClient
	HTTPPageRPC() pb.HTTPPageServiceClient
	HTTPAccessLogPolicyRPC() pb.HTTPAccessLogPolicyServiceClient
	HTTPCachePolicyRPC() pb.HTTPCachePolicyServiceClient
	HTTPFirewallPolicyRPC() pb.HTTPFirewallPolicyServiceClient
	HTTPFirewallRuleGroupRPC() pb.HTTPFirewallRuleGroupServiceClient
	HTTPFirewallRuleSetRPC() pb.HTTPFirewallRuleSetServiceClient
	HTTPLocationRPC() pb.HTTPLocationServiceClient
	HTTPWebsocketRPC() pb.HTTPWebsocketServiceClient
	HTTPRewriteRuleRPC() pb.HTTPRewriteRuleServiceClient
	HTTPAccessLogRPC() pb.HTTPAccessLogServiceClient
	SSLCertRPC() pb.SSLCertServiceClient
	SSLPolicyRPC() pb.SSLPolicyServiceClient
	MessageRPC() pb.MessageServiceClient
	IPListRPC() pb.IPListServiceClient
	IPItemRPC() pb.IPItemServiceClient
	FileRPC() pb.FileServiceClient
	FileChunkRPC() pb.FileChunkServiceClient
	RegionCountryRPC() pb.RegionCountryServiceClient
	RegionProvinceRPC() pb.RegionProvinceServiceClient
	LogRPC() pb.LogServiceClient
	DNSDomainRPC() pb.DNSDomainServiceClient
	DNSRPC() pb.DNSServiceClient
	ACMEUserRPC() pb.ACMEUserServiceClient
	ACMETaskRPC() pb.ACMETaskServiceClient
	UserRPC() pb.UserServiceClient
}
