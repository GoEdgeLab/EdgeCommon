// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type RecordType = string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeNS    RecordType = "NS"
	RecordTypeMX    RecordType = "MX"
	RecordTypeSRV   RecordType = "SRV"
	RecordTypeTXT   RecordType = "TXT"
	RecordTypeCAA   RecordType = "CAA"
	RecordTypeSOA   RecordType = "SOA"
)

type RecordTypeDefinition struct {
	Type        RecordType `json:"type"`
	Description string     `json:"description"`
	CanDefine   bool       `json:"canDefine"` // 用户是否可以自定义
}

func FindAllRecordTypeDefinitions() []*RecordTypeDefinition {
	return []*RecordTypeDefinition{
		{
			Type:        RecordTypeA,
			Description: "将域名指向一个IPV4地址",
			CanDefine:   true,
		},
		{
			Type:        RecordTypeCNAME,
			Description: "将域名指向另外一个域名",
			CanDefine:   true,
		},
		{
			Type:        RecordTypeAAAA,
			Description: "将域名指向一个IPV6地址",
			CanDefine:   true,
		},
		{
			Type:        RecordTypeNS,
			Description: "将子域名指定其他DNS服务器解析",
			CanDefine:   false,
		},
		{
			Type:        RecordTypeSOA,
			Description: "起始授权机构记录",
			CanDefine:   false,
		},
		{
			Type:        RecordTypeMX,
			Description: "将域名指向邮件服务器地址",
			CanDefine:   true,
		},
		{
			Type:        RecordTypeSRV,
			Description: "记录提供特定的服务的服务器",
			CanDefine:   true,
		},
		{
			Type:        RecordTypeTXT,
			Description: "文本长度限制512，通常做SPF记录或者校验域名所有者",
			CanDefine:   true,
		},
		{
			Type:        RecordTypeCAA,
			Description: "CA证书颁发机构授权校验",
			CanDefine:   true,
		},
	}
}

func FindAllUserRecordTypeDefinitions() []*RecordTypeDefinition {
	var result = []*RecordTypeDefinition{}
	for _, r := range FindAllRecordTypeDefinitions() {
		if r.CanDefine {
			result = append(result, r)
		}
	}
	return result
}
