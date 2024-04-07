// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package firewallconfigs

import "net/http"

// HTTPFirewallPageAction default page action
type HTTPFirewallPageAction struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	Status int    `yaml:"status" json:"status"`
	Body   string `yaml:"body" json:"body"`
}


func NewHTTPFirewallPageAction() *HTTPFirewallPageAction {
	return &HTTPFirewallPageAction{
		Status:  http.StatusForbidden,
		Body:    `<!DOCTYPE html>
<html lang="en">
<head>
	<title>403 Forbidden</title>
	<style>
		address { line-height: 1.8; }
	</style>
</head>
<body>
<h1>403 Forbidden By WAF</h1>
<address>Connection: ${remoteAddr} (Client) -&gt; ${serverAddr} (Server)</address>
<address>Request ID: ${requestId}</address>
</body>
</html>`,
	}
}