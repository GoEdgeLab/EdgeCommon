package firewallconfigs

import "github.com/iwind/TeaGo/maps"

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
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "XSS"
		group.Code = "xss"
		group.Description = "防跨站脚本攻击（Cross Site Scripting）"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "XSS攻击检测"
			set.Code = "1010"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code:    HTTPFirewallActionPage,
					Options: maps.Map{"status": 403, "body": ""},
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorContainsXSS,
				Value:             "",
				IsCaseInsensitive: false,
			})
			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// upload
	{
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = false
		group.Name = "文件上传"
		group.Code = "upload"
		group.Description = "防止上传可执行脚本文件到服务器"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
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
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = false
		group.Name = "Web Shell"
		group.Code = "webShell"
		group.Description = "防止远程执行服务器命令"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
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
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = false
		group.Name = "命令注入"
		group.Code = "commandInjection"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
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
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "路径穿越"
		group.Code = "pathTraversal"
		group.Description = "防止读取网站目录之外的其他系统文件"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
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
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "特殊目录"
		group.Code = "denyDirs"
		group.Description = "防止通过Web访问到一些特殊目录"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "特殊目录"
			set.Code = "6001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code:    HTTPFirewallActionPage,
					Options: maps.Map{"status": 403, "body": ""},
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestPath}",
				Operator:          HTTPFirewallRuleOperatorContainsAnyWord,
				Value:             "/.git\n/.svn\n/.htaccess\n/.idea\n/.env\n/.vscode",
				IsCaseInsensitive: true,
			})
			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// sql injection
	{
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = false
		group.Name = "SQL注入"
		group.Code = "sqlInjection"
		group.Description = "防止SQL注入漏洞"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "SQL注入检测"
			set.Code = "7010"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code:    HTTPFirewallActionPage,
					Options: maps.Map{"status": 403, "body": ""},
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${requestAll}",
				Operator:          HTTPFirewallRuleOperatorContainsSQLInjection,
				Value:             "",
				IsCaseInsensitive: false,
			})

			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// bot
	{
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "网络爬虫"
		group.Code = "bot"
		group.Description = "禁止一些网络爬虫"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = false
			set.Name = "搜索引擎"
			set.Code = "20001"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code:    HTTPFirewallActionPage,
					Options: maps.Map{"status": 403, "body": ""},
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${userAgent}",
				Operator:          HTTPFirewallRuleOperatorContainsAnyWord,
				Value:             "360spider\nadldxbot\nadsbot-google\napplebot\nadmantx\nalexa\nbaidu\nbingbot\nbingpreview\nfacebookexternalhit\ngooglebot\nproximic\nslurp\nsogou\ntwitterbot\nyandex\nspider",
				IsCaseInsensitive: true,
			})

			group.AddRuleSet(set)
		}

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "爬虫工具"
			set.Code = "20003"
			set.Connector = HTTPFirewallRuleConnectorAnd
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code:    HTTPFirewallActionPage,
					Options: maps.Map{"status": 403, "body": ""},
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${userAgent}",
				Operator:          HTTPFirewallRuleOperatorContainsAnyWord,
				Value:             "python\npycurl\nhttp-client\nhttpclient\napachebench\nnethttp\nhttp_request\njava\nperl\nruby\nscrapy\nphp\nrust",
				IsCaseInsensitive: true,
			})
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${userAgent}",
				Operator:          HTTPFirewallRuleOperatorNotContainsAnyWord,
				Value:             "goedge",
				IsCaseInsensitive: true,
				Description:       "User-Agent白名单",
			})

			group.AddRuleSet(set)
		}

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "下载工具"
			set.Code = "20004"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionTag,
					Options: maps.Map{
						"tags": []string{"download"},
					},
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${userAgent}",
				Operator:          HTTPFirewallRuleOperatorContainsAnyWord,
				Value:             "wget\ncurl",
				IsCaseInsensitive: true,
			})

			group.AddRuleSet(set)
		}

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "空Agent"
			set.Code = "20002"
			set.Connector = HTTPFirewallRuleConnectorOr
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code:    HTTPFirewallActionPage,
					Options: maps.Map{"status": 403, "body": ""},
				},
			}

			// 空Agent
			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${userAgent}",
				Operator:          HTTPFirewallRuleOperatorEqString,
				Value:             "",
				IsCaseInsensitive: false,
			})

			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// cc2
	{
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "CC攻击"
		group.Description = "Challenge Collapsar，防止短时间大量请求涌入，请谨慎开启和设置"
		group.Code = "cc2"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "CC单URL请求数"
			set.Description = "限制单IP在一定时间内对单URL的请求数"
			set.Code = "8001"
			set.Connector = HTTPFirewallRuleConnectorAnd
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
					Options: maps.Map{
						"timeout": 1800,
					},
				},
			}
			set.IgnoreLocal = true
			set.AddRule(&HTTPFirewallRule{
				IsOn:     true,
				Param:    "${cc2}",
				Operator: HTTPFirewallRuleOperatorGt,
				Value:    "120",
				CheckpointOptions: map[string]interface{}{
					"keys":              []string{"${remoteAddr}", "${requestPath}"},
					"period":            "60",
					"threshold":         120,
					"enableFingerprint": true,
				},
				IsCaseInsensitive: false,
			})

			group.AddRuleSet(set)
		}
		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "CC请求数"
			set.Description = "限制单IP在一定时间内的总体请求数"
			set.Code = "8002"
			set.Connector = HTTPFirewallRuleConnectorAnd
			set.IgnoreLocal = true
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
					Options: maps.Map{
						"timeout": 1800,
					},
				},
			}
			set.AddRule(&HTTPFirewallRule{
				IsOn:     true,
				Param:    "${cc2}",
				Operator: HTTPFirewallRuleOperatorGt,
				Value:    "1200",
				CheckpointOptions: map[string]interface{}{
					"keys":              []string{"${remoteAddr}"},
					"period":            "60",
					"threshold":         1200,
					"enableFingerprint": true,
				},
				IsCaseInsensitive: false,
			})

			group.AddRuleSet(set)
		}

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "随机URL攻击"
			set.Description = "限制用户使用随机URL访问网站"
			set.Code = "8003"
			set.Connector = HTTPFirewallRuleConnectorAnd
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
					Options: maps.Map{
						"timeout": 1800,
					},
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:              true,
				Param:             "${args}",
				Operator:          HTTPFirewallRuleOperatorMatch,
				Value:             `^[0-9a-zA-Z_\-.]{12,}$`,
				IsCaseInsensitive: false,
			})

			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// custom
	{
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "防盗链"
		group.Description = "防止第三方网站引用本站资源。"
		group.Code = "referer"
		group.IsTemplate = true

		{
			var set = &HTTPFirewallRuleSet{}
			set.IsOn = true
			set.Name = "防盗链"
			set.Description = "防止第三方网站引用本站资源"
			set.Code = "9001"
			set.Connector = HTTPFirewallRuleConnectorAnd
			set.Actions = []*HTTPFirewallActionConfig{
				{
					Code: HTTPFirewallActionBlock,
					Options: maps.Map{
						"timeout": 60,
					},
				},
			}

			set.AddRule(&HTTPFirewallRule{
				IsOn:     true,
				Param:    "${refererBlock}",
				Operator: HTTPFirewallRuleOperatorEq,
				Value:    "0",
				CheckpointOptions: map[string]interface{}{
					"allowEmpty":      true,
					"allowSameDomain": true,
					"allowDomains":    []string{"*"},
				},
				IsCaseInsensitive: false,
			})

			group.AddRuleSet(set)
		}

		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	// custom
	{
		var group = &HTTPFirewallRuleGroup{}
		group.IsOn = true
		group.Name = "自定义规则分组"
		group.Description = "我的自定义规则分组，可以将自定义的规则放在这个分组下"
		group.Code = "custom"
		group.IsTemplate = true
		policy.Inbound.Groups = append(policy.Inbound.Groups, group)
	}

	return policy
}
