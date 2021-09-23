package firewallconfigs

type HTTPFirewallRuleConnector = string

const (
	HTTPFirewallRuleConnectorAnd = "and"
	HTTPFirewallRuleConnectorOr  = "or"
)

func HTTPFirewallTemplate() *HTTPFirewallPolicy {
	policy := &HTTPFirewallPolicy{}
	policy.IsOn = true
	policy.Inbound = &HTTPFirewallInboundConfig{}
	policy.Outbound = &HTTPFirewallOutboundConfig{}

	// xss
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "XSS"
		group.Code = "xss"
		group.Description = "防跨站脚本攻击（Cross Site Scripting）"

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "Javascript事件"
			set.Code = "1001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestURI}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `(onmouseover|onmousemove|onmousedown|onmouseup|onerror|onload|onclick|ondblclick|onkeydown|onkeyup|onkeypress)\s*=`, // TODO more keywords here
				IsCaseInsensitive: true,
			})
			group.AddRuleSet(set)
		}

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "Javascript函数"
			set.Code = "1002"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestURI}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `(alert|eval|prompt|confirm)\s*\(`, // TODO more keywords here
				IsCaseInsensitive: true,
			})
			group.AddRuleSet(set)
		}

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "HTML标签"
			set.Code = "1003"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestURI}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `<(script|iframe|link)`, // TODO more keywords here
				IsCaseInsensitive: true,
			})
			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// upload
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "文件上传"
		group.Code = "upload"
		group.Description = "防止上传可执行脚本文件到服务器"

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "上传文件扩展名"
			set.Code = "2001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestUpload.ext}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `\.(php|jsp|aspx|asp|exe|asa|rb|py)\b`, // TODO more keywords here
				IsCaseInsensitive: true,
			})
			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// web shell
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "Web Shell"
		group.Code = "webShell"
		group.Description = "防止远程执行服务器命令"

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "Web Shell"
			set.Code = "3001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `\b(eval|system|exec|execute|passthru|shell_exec|phpinfo)\s*\(`, // TODO more keywords here
				IsCaseInsensitive: true,
			})
			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// command injection
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "命令注入"
		group.Code = "commandInjection"

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "命令注入"
			set.Code = "4001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestURI}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `\b(pwd|ls|ll|whoami|id|net\s+user)\s*$`, // TODO more keywords here
				IsCaseInsensitive: false,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestBody}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `\b(pwd|ls|ll|whoami|id|net\s+user)\s*$`, // TODO more keywords here
				IsCaseInsensitive: false,
			})
			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// path traversal
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "路径穿越"
		group.Code = "pathTraversal"
		group.Description = "防止读取网站目录之外的其他系统文件"

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "路径穿越"
			set.Code = "5001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestURI}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `((\.+)(/+)){2,}`, // TODO more keywords here
				IsCaseInsensitive: false,
			})
			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// special dirs
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "特殊目录"
		group.Code = "denyDirs"
		group.Description = "防止通过Web访问到一些特殊目录"

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "特殊目录"
			set.Code = "6001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestPath}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `/\.(git|svn|htaccess|idea)\b`, // TODO more keywords here
				IsCaseInsensitive: true,
			})
			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// sql injection
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "SQL注入"
		group.Code = "sqlInjection"
		group.Description = "防止SQL注入漏洞"

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "Union SQL Injection"
			set.Code = "7001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `union[\s/\*]+select`,
				IsCaseInsensitive: true,
			})

			group.AddRuleSet(set)
		}

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "SQL注释"
			set.Code = "7002"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `/\*(!|\x00)`,
				IsCaseInsensitive: true,
			})

			group.AddRuleSet(set)
		}

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "SQL条件"
			set.Code = "7003"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `\s(and|or|rlike)\s+(if|updatexml)\s*\(`,
				IsCaseInsensitive: true,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `\s+(and|or|rlike)\s+(select|case)\s+`,
				IsCaseInsensitive: true,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `\s+(and|or|procedure)\s+[\w\p{L}]+\s*=\s*[\w\p{L}]+(\s|$|--|#)`,
				IsCaseInsensitive: true,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `\(\s*case\s+when\s+[\w\p{L}]+\s*=\s*[\w\p{L}]+\s+then\s+`,
				IsCaseInsensitive: true,
			})

			group.AddRuleSet(set)
		}

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "SQL函数"
			set.Code = "7004"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `(updatexml|extractvalue|ascii|ord|char|chr|count|concat|rand|floor|substr|length|len|user|database|benchmark|analyse)\s*\(`,
				IsCaseInsensitive: true,
			})

			group.AddRuleSet(set)
		}

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "SQL附加语句"
			set.Code = "7005"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `;\s*(declare|use|drop|create|exec|delete|update|insert)\s`,
				IsCaseInsensitive: true,
			})

			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// bot
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "网络爬虫"
		group.Code = "bot"
		group.Description = "禁止一些网络爬虫"

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "常见网络爬虫"
			set.Code = "20001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${userAgent}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `Googlebot|AdsBot|bingbot|BingPreview|facebookexternalhit|Slurp|Sogou|proximic|Baiduspider|yandex|twitterbot|spider|python`,
				IsCaseInsensitive: true,
			})

			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// cc2
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "CC攻击"
		group.Description = "Challenge Collapsar，防止短时间大量请求涌入，请谨慎开启和设置"
		group.Code = "cc2"

		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "CC单URL请求数"
			set.Description = "限制单IP在一定时间内对单URL的请求数"
			set.Code = "8001"
			set.Connector = HTTPFirewallRuleConnectorAnd
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:     true,
				Param:    "${cc2}",
				Operator: HTTPFirewallRuleOperatorGt,
				Value:    "120",
				CheckpointOptions: map[string]interface{}{
					"keys":      []string{"${remoteAddr}", "${requestPath}"},
					"period":    "60",
					"threshold": 120,
				},
				IsCaseInsensitive: false,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${remoteAddr}",
				Operator:          HTTPFirewallRuleOperatorNotIPRange,
				Value:             `127.0.0.1/8`,
				IsCaseInsensitive: false,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${remoteAddr}",
				Operator:          HTTPFirewallRuleOperatorNotIPRange,
				Value:             `192.168.0.1/16`,
				IsCaseInsensitive: false,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${remoteAddr}",
				Operator:          HTTPFirewallRuleOperatorNotIPRange,
				Value:             `10.0.0.1/8`,
				IsCaseInsensitive: false,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${remoteAddr}",
				Operator:          HTTPFirewallRuleOperatorNotIPRange,
				Value:             `172.16.0.1/12`,
				IsCaseInsensitive: false,
			})

			group.AddRuleSet(set)
		}
		{
			set := &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "CC请求数"
			set.Description = "限制单IP在一定时间内的总体请求数"
			set.Code = "8001"
			set.Connector = HTTPFirewallRuleConnectorAnd
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:     true,
				Param:    "${cc2}",
				Operator: HTTPFirewallRuleOperatorGt,
				Value:    "1200",
				CheckpointOptions: map[string]interface{}{
					"keys":      []string{"${remoteAddr}"},
					"period":    "60",
					"threshold": 1200,
				},
				IsCaseInsensitive: false,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${remoteAddr}",
				Operator:          HTTPFirewallRuleOperatorNotIPRange,
				Value:             `127.0.0.1/8`,
				IsCaseInsensitive: false,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${remoteAddr}",
				Operator:          HTTPFirewallRuleOperatorNotIPRange,
				Value:             `192.168.0.1/16`,
				IsCaseInsensitive: false,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${remoteAddr}",
				Operator:          HTTPFirewallRuleOperatorNotIPRange,
				Value:             `10.0.0.1/8`,
				IsCaseInsensitive: false,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${remoteAddr}",
				Operator:          HTTPFirewallRuleOperatorNotIPRange,
				Value:             `172.16.0.1/12`,
				IsCaseInsensitive: false,
			})

			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// custom
	{
		group := &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "自定义规则分组"
		group.Description = "我的自定义规则分组，可以将自定义的规则放在这个分组下"
		group.Code = "custom"
		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	return policy
}
